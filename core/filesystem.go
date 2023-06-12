/*
Copyright © 2022 Blooym

GPLv3 License, see the LICENSE file for more information.
*/
package core

import (
	"crypto"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/viper"
)

/*
	GetCustomLocation returns the custom location of the passed arg is any of the pre-saved locations, otherwise it just returns the arg.
	Arguments:
		arg<string>: The argument to check.
	Example:
		dir := GetCustomLocation("steam")
		fmt.Println(dir) // $HOME/.steam/root/compatabilitytools.d/
	Returns:
		string: A path.
*/
func GetCustomLocation(arg string) string {
	customLocations := viper.GetStringMapString("app.customlocations")
	for key, value := range customLocations {
		if key == arg {
			return value
		}
	}
	return arg
}

/*
	UsePath returns the path with sane changes to it.
	Arguments:
		path<string>: The path to adjust.
		trailSlash<bool>: Whether or not to have a trailing slash.
	Example:
		dir := UsePath("$HOME/Dowloads/, false)
		fmt.Println(dir) // $HOME/Downloads
	Returns:
		string: A path.
*/
func UsePath(path string, trailSlash bool) string {

	Debug("UsePath: Attempting to format path: " + path)

	// If trail slash is true and there is no trailing slash add one
	if trailSlash && !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	// If trail slash is false and there is a trailing slash remove it
	if !trailSlash && strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	// If short notation for the home directory is used, expand it to $HOME (~/ -> $HOME/)
	if strings.HasPrefix(path, "~/") {
		homeDir, _ := os.UserHomeDir()
		path = strings.Replace(path, "~/", homeDir+"/", 1)
	}

	Debug("UsePath: Finished formatting path, result was: " + path)

	return path
}

/*
	DeleteUserTemp clears the user's temp directory
	Returns:
		error: An error if one occurs.
*/
func DeleteUserTemp() error {
	tempDir := os.TempDir() + "/proto/" + fmt.Sprint(os.Getuid())
	err := os.RemoveAll(UsePath(tempDir, true))
	if err != nil {
		return err
	}
	Debug("DeleteDir: Deleted directory at: " + tempDir)
	return nil
}

/*
	GetUserTemp creates a temporary directory in the proto temp directory
	Example:
		dir, err := GetUserTemp()
		fmt.Println(dir) // /tmp/proto/1000/
	Returns:
		string: The path to the temp directory.
		error: An error if one occurs.
*/
func GetUserTemp() (string, error) {
	userTempDir := viper.GetString("storage.tmp") + "proto/"
	err := os.MkdirAll(userTempDir, os.ModePerm)
	os.Chmod(userTempDir, 0777)
	if err != nil {
		return "", err
	}

	tempDir := userTempDir + fmt.Sprint(os.Getuid())
	err = os.Mkdir(tempDir, 0700)
	if err != nil {
		return "", err
	}
	Debug("CreateTemp: Created temp directory at: " + tempDir)
	return UsePath(tempDir, true), nil
}

/*
	DownloadFile downloads the file from the given URL, following redirects if needed. The final file will be put at the given path
	Arguments:
		path<string>: The path to download the file to.
		url<string>: The URL to download the file from.
	Example:
		file, err := DownloadFile("$HOME/Downloads/file.tar.gz", "https://example.com/file.tar.gz")
	Returns:
		os.FileInfo: The file that was downloaded.
		error: An error if one occurs.
*/
func DownloadFile(path, url string) (os.FileInfo, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return nil, err
		}
		Debug("DownloadFile: Created directory: " + filepath.Dir(path))
	}

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	defer out.Close()

	// Fetch the file from the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	Debug("DownloadFile: Downloading file from: " + url)

	// Set up a progress bar
	tmpl := `{{ cycle . "⠃" "⠆" "⠤" "⠰" "⠘" "⠉" }} Installing {{string . "src"}} [{{percent .}} | {{speed . "%s/s"}} | {{ rtime .}}]`
	bar := pb.ProgressBarTemplate(tmpl).Start64(resp.ContentLength).Set("src", strings.Split(url, "/")[len(strings.Split(url, "/"))-1])
	reader := bar.NewProxyReader(resp.Body)

	defer resp.Body.Close()

	// Write the data to the file
	_, err = io.Copy(out, reader)
	if err != nil {
		return nil, err
	}

	// Check if the file is valid
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	Debug("DownloadFile: Downloaded file to: " + path)

	bar.Finish()

	// Get the downloaded file and return it
	return os.Stat(path)
}

/*
	ExtractTar extracts the given tar file to the given path using gnu-tar.
	Arguments:
		tarPath<string>: The path to the tar file.
		extractPath<string>: The path to extract the tar file to.
	Example:
		err := ExtractTar("$HOME/Downloads/file.tar.gz", "$HOME/Downloads/")
	Returns:
		error: An error if one occurs.
*/
func ExtractTar(tarPath, extractPath string) error {

	// If path doesnt exist create it
	if _, err := os.Stat(extractPath); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(extractPath), os.ModePerm)
		if err != nil {
			return err
		}

		Debug("DownloadFile: Created directory: " + filepath.Dir(extractPath))
	}

	cmd := exec.Command("tar", "-xf", tarPath, "-C", extractPath)
	err := cmd.Start()

	if err != nil {
		return err
	}

	err = cmd.Wait()

	if err != nil {
		return err
	}

	return nil
}

/*
	Tries to match a given file's sha512sum against the given sum file
	Arguments:
		filePath<string>: The path to the file to check.
		sumPath<string>: The path to the sum file.
	Example:
		match, err := CheckSum("$HOME/Downloads/file.tar.gz", "$HOME/Downloads/file.tar.gz.sha512sum")
		fmt.Println(match) // true
	Returns:
		bool: Whether or not the file matches the sum.
		error: An error if one occurs.
*/
func MatchChecksum(filePath, sumPath string) (bool, error) {
	// Get the sum of the file with crypto inbuilt
	h := crypto.SHA512.New()
	f, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	// Get the sum of the file in the sum file
	sum, err := ioutil.ReadFile(sumPath)
	if err != nil {
		return false, err
	}

	// Check all lines for the files sum
	for _, line := range strings.Split(string(sum), "\n") {
		Debug("MatchChecksum: Attempting to match checksum for files: " + filePath + " and " + sumPath)
		if strings.HasPrefix(line, fmt.Sprintf("%x", h.Sum(nil))) {
			return true, nil
		}
	}

	return false, nil
}

/*
	GetDirSize gets the size of the given directory in bytes.
	Arguments:
		path<string>: The path to the directory.
	Example:
		size, err := GetDirSize("$HOME/Downloads/")
		fmt.Println(size) // 4194304
	Returns:
		int64: The size of the directory in bytes.
		error: An error if one occurs.
*/
func GetDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

/*
	HumanReadableBytes converts the given bytes to a human readable amount of bytes and a unit.
	Arguments:
		bytes<int64>: The bytes to convert.
	Example:
		humanReadableBytes, unit := HumanReadableBytes(4194304)
		fmt.Println(humanReadableBytes) // 4
		fmt.Println(unit) // MB
	Returns:
		int64: The human readable amount of bytes.
		string: The unit of the bytes.
*/
func HumanReadableBytes(bytes int64) (int64, string) {
	switch {
	case bytes < 1024:
		return bytes, "B"
	case bytes < 1024*1024:
		return bytes / 1024, "KB"
	case bytes < 1024*1024*1024:
		return bytes / (1024 * 1024), "MB"
	default:
		return bytes / (1024 * 1024 * 1024), "GB"
	}
}
