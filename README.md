# Dhaka Comics Downloader

Download purchased digital comics from Dhaka Comics.

## Overview

This CLI app allows you to download purchased or free digital comics from Dhaka Comics.

# NOTE

-   **To download a comic, you need to have premium account. Please do not attempt to use this tool unless you've premium account.**
- **This tool does not bypass premium content restrictions.**
-   **I am NOT held responsible for your account getting suspended as a result from the use of this program!**
-   This program is WIP, the code is provided as-is and I am not held resposible for any legal issues resulting from the use of this program.

# Description

Linux is the primary development OS (Windows, Mac untested).

_**Note**:_ _You will need to download and install these manually!_

-   [Go](https://go.dev/dl/)

# Build & Installation

Clone the repository: `git clone https://github.com/joybiswas007/dc-dl` <br/>
Tidy up the Go modules: `go mod tidy` <br/>
Build the project. You can use one of the following commands based on your operating system: <br/>
For all platforms: `make build` <br/>
For Linux: `make build-linux`

# Usage

Visit Dhakacomics.com and export cookie file as netscape format and put it inside `cookies` diretory
as `cookies.txt` file use any firefox or chrome extension.

```dc-dl --help
A CLI app which allows you to download free and purchased comics from dhaka comics

Usage:
  dc-dl [flags]

Flags:
  -h, --help         help for dc-dl
      --url string   Pass comic url you want to download

dc-dl --url https://www.dhakacomics.com/digital-comics/durjoy-1
``

## Disclaimer
* This tool was written for educational purposes. I will not be responsible if you use this program in bad faith.
* `dc-dl` is not affiliated with Dhaka Comics.
