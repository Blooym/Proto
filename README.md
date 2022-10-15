<!-- Repository Header Begin -->
<div align="center">

<img src="./.assets/Banners/banner.png" alt="Proto Logo">
  
Install and manage custom runners with ease on supported systems.

**[View Issues](https://github.com/BitsOfAByte/proto/issues) · [Install](#installation) · [Contributing](https://github.com/BitsOfAByte/proto/blob/main/CONTRIBUTING.md)**
  
<a href="#"> 
  <img src="https://img.shields.io/github/downloads/BitsOfAByte/proto/total?style=flat" alt="Download Count Badge">
  <img src="https://img.shields.io/github/v/tag/BitsOfAByte/proto?color=blue&label=Version&sort=semver&style=flat" alt="Release Badge">
</a>
  
</div>

---

<!-- Repository Header End -->

## About

Proto is a command line tool designed to make it easier to manage custom runner installations (eg. Proton-GE) without the need to manually navigate the filesystem or extract tar files when a new build is released. 

### Key Features

  - The ability to add multiple runner release sources (GitHub only, must ship the runner in a `.tar` format)
  - The ability to bind directories to keywords so you don't have to remember or type them every time (ig. `steam` -> `~/.steam/root/compatabilitytools.d`)
  - The ability to pull information about any release directly from GitHub
  - Powerful but minimal configuration (which is stored in a very portable format)
  - Fully documented through the command line using the `-h` flag after any command
  - Shell completion for bash, fish, powershell and zsh (via the `completion` command)
  - A built in app-updater for manual binary installs
  - Responsive & easy to use
  - Checksum validation (sha512sum only)

## Usage

Proto is fully documentated from the command line by using the `-h` flag after any command, which includes usage examples and a full list of flags/arguments. A basic installation from *Source #1* to the *Steam* directory would look like:
```
proto install --dir steam --source 1
```

You can tweak the configuration for Proto by running the following:
```
proto config
```


## Installation

### Dependancies
Proto currently requires the following packages in order to function: [tar](https://www.gnu.org/software/tar/)

If you are using a package manager to install, these should be automatically installed alongside Proto if they are missing from your system, however if you are building from source or installing from an archive, make sure these are also present. 

### Methods

#### APT Package Manager

If you are using an Ubuntu-derivative system then use this installation method.

<details>
<summary>Show Steps</summary>

<br>
  
1. Add the repository hosting Proto to your apt sources directory (Only run this once)
```
echo "deb [trusted=yes] https://packages.bitsofabyte.dev/apt/ /" | sudo tee -a /etc/apt/sources.list.d/bitsofabyte.list && sudo apt update
``` 

2. Install Proto to your system
```
sudo apt install proto
```

</details>  

---

#### Yum/DNF Package Manager

If you are using Fedora, OpenSUSE, or any other system that supports the yum/dnf package manager then use this installation method.

<details>
<summary>Show Steps</summary>
<br>
  
1. Add the repository hosting Proto to your yum/dnf repo directory (Only run this once)
```
echo "[BitsOfAByte]            
name=BitsOfAByte Packages         
baseurl=https://packages.bitsofabyte.dev/yum/
enabled=1
gpgcheck=0" | sudo tee -a /etc/yum.repos.d/bitsofabyte.repo && sudo yum update
``` 

2. Install Proto to your system
```
sudo yum install proto
```

</details>  

---

#### Homebrew Package Manager

If your distributions package manager is not listed here or you wish to use [Homebrew](https://brew.sh).

<details>
<summary>Show Steps</summary>
<br>
  
1. Install homebrew if you haven't already got it
```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

2. Add the tap for Proto to homebrew
```
brew tap BitsOfAByte/proto https://github.com/BitsOfAByte/proto.git
```

3. Install proto to your system
```
brew install proto
```
  
</details>

---

#### Manual Installation

Manually download a release file from the GitHub releases.
<details>  
<summary>Show Steps</summary>
  
1. Download the [newest release](https://github.com/BitsOfAByte/proto/releases/latest) for your system/architecture
2. Extract the tar archive or install a `.rpm`/`.deb` package (these will also provide the repository to handle automatic updates)

If you aren't sure on what architecture you need to download, you should try `amd64` first as it is the most common.

</details>

---

#### From Source

Build Proto directly from the source for your system. Only recommended for advanced users that prefer to build from source or contributors.
<details>  
<summary>Show Steps</summary>

1. Make sure you have [Go](https://go.dev/) installed on your system and setup properly, alternatively use the [Devcontainer](./.devcontainer) setup.
2. Install [GoReleaser](https://goreleaser.com/) if you want to build using the supported buildsystem (Optional unless contributing)
3. Run `make build` to build the binary for your system, or `make build-all` to build for all supported systems. You can optionally use `./build/scripts/upx.sh <file>` to compress the binary with UPX (This is done automatically when using make and having GoReleaser installed with `SKIP_COMPRESS=false` set)
4. You will find all the binaries in the `./dist` directory alongside any other build artifacts.

</details>  
