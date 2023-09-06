package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"

	"github.com/Blooym/proto/cmd"
	"github.com/spf13/cobra/doc"
)

var build_dir = "./.build_data/"

// Tasks to run as a pre-build hook
func main() {
	cleanup()
	createBuildDir(build_dir)

	generateMANPages()
}

// Create the build directory if it doesn't exist
func createBuildDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
}

func generateMANPages() {
	// Generate the man pages
	header := &doc.GenManHeader{
		Title:   "PROTO",
		Section: "1",
		Source:  "Proto",
		Manual:  "Proto Manual",
	}
	manDir := build_dir + "man"
	os.MkdirAll(manDir, os.ModePerm)
	cmd.RootCmd.DisableAutoGenTag = true
	doc.GenManTree(cmd.RootCmd, header, manDir)

	// GZip them.
	manFiles, err := os.ReadDir(manDir)
	if err != nil {
		panic(err)
	}

	for _, file := range manFiles {
		var b bytes.Buffer
		w := gzip.NewWriter(&b)
		f, err := os.Open(manDir + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		// write the file contents into the gzip file
		contents, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		_, err = w.Write([]byte(contents))
		if err != nil {
			panic(err)
		}
		w.Close()
		f.Close()

		// write the gzip file to disk
		gzFile, err := os.Create(manDir + "/" + file.Name() + ".gz")
		if err != nil {
			panic(err)
		}

		_, err = gzFile.Write(b.Bytes())
		if err != nil {
			panic(err)
		}

		gzFile.Close()

		// remove the original file
		os.Remove(manDir + "/" + file.Name())
	}
}

// Cleanup the build directory
func cleanup() {
	os.RemoveAll(build_dir)
}
