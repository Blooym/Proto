<!-- Repository Header Begin -->
<div align="center">

<img src="./.assets/Banners/banner.png" alt="Proto Logo">
  
Install and manage custom Proton installations with ease on supported systems.

**[View Issues](https://github.com/BitsOfAByte/proto/issues) · [Install](#installation) · [Contributing](https://github.com/BitsOfAByte/proto/blob/main/CONTRIBUTING.md)**
  
<a href="#"> 
  <img src="https://img.shields.io/github/downloads/BitsOfAByte/proto/total?style=flat" alt="Download Count Badge">
  <img src="https://img.shields.io/github/v/tag/BitsOfAByte/proto?color=blue&label=Version&sort=semver&style=flat" alt="Release Badge">
</a>
  
</div>

---

<!-- Repository Header End -->

## About
Proto is a tool designed to make downloading and managing custom WINE/Proton installations as convinent and easy as possible. It provides support for multiple sources, custom installation directories, in-tool release data and more. 

## How to Use
First off, you'll need to download Proto to your system using one of the supported [installation methods](#installation). Once you have it installed, the entire app is documented in the command line by running the `--help` flag after any command, which will provide details on how to use it. 

Configuration is also provided by running the `proto config` command, which will allow you to tweak a variety of settings straight from the command line.

Proto can only install from archive formats supported by GNU tar.

## Installation

### Dependancies
Proto currently requires the following packages in order to function: [tar](https://www.gnu.org/software/tar/)

If you are using a package manager to install, these should be automatically installed alongside Proto if they are missing from your system, however if you are building from source or installing from an archive, make sure these are also present. It is planned to be dependency free by time a full release is made. 

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
Manually install a Binary from the release Archives.
<details>  
<summary>Show Steps</summary>
  
1. Download the [newest release](https://github.com/BitsOfAByte/proto/releases/latest) for your system/architecture
2. Extract the binary into your system path or add the binary to your path.

If you aren't sure on what architecture you need to download, you should try `amd64` first as it is the most common for everyday hardware.

</details>

---

#### From Source
Build Proto directly from the GitHub source for any supported platform.
<details>  
<summary>Show Steps</summary>
  
Building Proto from source is not recommended for beginners, but if you know what you're doing then follow these steps: 
1. Install [Go](https://go.dev/) on your system
2. Download the [GoReleaser](https://goreleaser.com/) package
3. Clone the repository to your system with `git clone https://github.com/BitsOfAByte/proto`
4. Inside the repository directory, run `goreleaser build --single-target --rm-dist --snapshot` to build.

You will find the compiled binary for your OS & Arch inside of the `/dist` folder.

</details>  
