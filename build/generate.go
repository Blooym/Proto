/*
Copyright © 2022 BitsOfAByte

*/
package main

import (
	"BitsOfAByte/proto/shared"
	"fmt"
	"io"
	"os"
)

var build_dir = "./.build_data/"

// Tasks to run as a pre-build hook
func main() {
	cleanup()
	createBuildDir(build_dir)

	generateDesktop()
	generateMetainfo()
	generateIcon()
}

// Create the build directory if it doesn't exist
func createBuildDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
}

// Create a file in the build directory
func createBuildFile(fileName string, data string) {
	file, err := os.Create(build_dir + fileName)
	shared.Check(err)

	defer file.Close()

	_, err = file.WriteString(data)
	shared.Check(err)

	file.Sync()
}

// Generate the .desktop file
func generateDesktop() {

	version := os.Args[1]

	if version == "" {
		panic("No version specified")
	}

	fileData := fmt.Sprintf(`[Desktop Entry]
Version=%s
Type=Application
Name=Proto
GenericName=Proto
Comment=Proto compatability tool manager
Icon=/usr/share/icons/proto/icon.png
Exec=proto gui
Terminal=true
Actions=NewShortcut;
Categories=ConsoleOnly;Utility;X-GNOME-Utilities;FileTools;
Keywords=proton;steamplay;
StartupNotify=true
NoDisplay=true

[Desktop Action NewShortcut]
Name=Install Latest Proton
Exec=proto install`, version)

	createBuildFile("dev.bitsofabyte.proto.desktop", fileData)
}

// Generate the .metainfo.xml file
func generateMetainfo() {
	fileData := `<?xml version="1.0" encoding="UTF-8"?>
<!-- Copyright 2020 BitsOfAByte -->
<component type="desktop-application">
  <id>dev.bitsofabyte.proto</id>
  <name>Proto</name>
  <developer_name>BitsOfAByte</developer_name>
  <content_rating type="oars-1.1" />
  <icon type="local" width="128" height="128">/usr/share/icons/proto/icon.png</icon>
  <launchable type="desktop-id">dev.bitsofabyte.proto.desktop</launchable>
  <metadata_license>MIT</metadata_license>
  <project_license>GPL-3.0-only</project_license>
  <summary>Manage custom Proton installations</summary>
  <description>
    <p>
      Install and manage custom Proton installations easily and quickly using a clean user interface (coming soon), or even the command line.
    </p>
  </description>

  <categories>
	<category>utility</category>
  </categories>

  <provides>
    <binary>proto</binary>
  </provides>

  <screenshots>
    <screenshot type="default">
      <caption>The Main CLI Page</caption>
      <image type="source">https://raw.githubusercontent.com/BitsOfAByte/proto/main/.assets/Screenshots/main_app_screenshot.png</image>
    </screenshot>
  </screenshots>

  <recommends>
	<display_length compare="ge">medium</display_length>
	<control>keyboard</control>
	<control>pointing</control>
	<control>console</control>
  </recommends>

  <url type="homepage">http://github.com/BitsOfAByte/proto</url>
  <url type="bugtracker">https://github.com/BitsOfAByte/proto/issues</url>
  <url type="faq">https://github.com/BitsOfAByte/proto#readme</url>
  <url type="help">https://github.com/BitsOfAByte/proto#readme</url>
  <url type="contact">https://github.com/BitsOfAByte/proto/discussions</url>
  <url type="vcs-browser">https://github.com/BitsOfAByte/proto</url>
  <url type="contribute">https://github.com/BitsOfAByte/proto/blob/main/CONTRIBUTING.md</url>
</component>`

	createBuildFile("dev.bitsofabyte.proto.metainfo.xml", fileData)
}

// Fetch the icon from the assets and put it in the build directory
func generateIcon() {
	srcFile, err := os.Open("./.assets/Logos/icon.png")
	shared.Check(err)
	defer srcFile.Close()

	destFile, err := os.Create(build_dir + "icon.png")
	shared.Check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	shared.Check(err)

	err = destFile.Sync()
	shared.Check(err)
}

// Cleanup the build directory
func cleanup() {
	os.RemoveAll(build_dir)
}
