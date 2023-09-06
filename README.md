<!-- Repository Header Begin -->
<div align="center">

<img src="./.assets/Banners/banner.png" alt="Proto Logo">
  
Install and manage custom runners with ease on supported systems.

**[View Issues](https://github.com/Blooym/proto/issues) · [Install](#installation) · [Contributing](https://github.com/Blooym/proto/blob/main/CONTRIBUTING.md)**
  
<a href="#"> 
  <img src="https://img.shields.io/github/downloads/Blooym/proto/total?style=flat" alt="Download Count Badge">
  <img src="https://img.shields.io/github/v/tag/Blooym/proto?color=blue&label=Version&sort=semver&style=flat" alt="Release Badge">
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

> [!IMPORTANT]  
> The previously available repository at `packages.bitsofabyte.dev` is deprecated and will not be used for futher updates. Please remove it from your system's repositories, thanks!

### Dependencies
Proto currently requires the following packages in order to function: [tar](https://www.gnu.org/software/tar/)

### Methods

---

#### Manual Installation

Manually download a release file from the GitHub releases.
<details>  
<summary>Show Steps</summary>
  
1. Download the [newest release](https://github.com/Blooym/proto/releases/latest) for your system/architecture
2. Extract the tar archive and place it somewhere inside of your `$PATH`

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
