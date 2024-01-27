# Dhaka Comics Downloader in Go

A tool written in go to download purchased digital comics from Dhaka Comics.

## Overview

This application allows you to download only purchased or free digital comics from Dhaka Comics.

| Descriptions                            | Notes                                                                                    |
| --------------------------------------- | ---------------------------------------------------------------------------------------- |
| **Internet Connection Required**        | Requires an active internet connection to download digital comics.                         |
| **Legal Restrictions**                  | Abides by legal restrictions and does not support downloading comics illegally.           |
| **Doesn't Bypass Premium Content**      | This tool does not bypass premium content restrictions; it adheres to access permissions.  |


## Features

- Download comics from Dhaka Comics using the provided URL.
- Organize downloads in a "downloads" directory.
- Display progress information during the download process.

## Prerequisites

- Go (Golang) installed on your system.


## Install

To use this application

1. Clone the repo: `git clone https://github.com/joybiswas007/dhakacomics-dl`
2. Install required packages: `go mod download`
3. Fill the .env file:
```
COOKIE="paste your cookie here"
VIEW_COMIC=https://www.dhakacomics.com/view-comic
X_CSRF_UMP_TOKEN="UMP TOKEN here"
USER_AGENT="Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

```
To get the cookie after loggin  open network tab and select `www.dhakacomics.com` and
from `Request Headers` get the Cookie value and paste it inside the .env file. 

To get the UMP token open any digital comic and open the network tab and filter results 
using `Fetch/XHR` and click on the `view-comic` and check  `Request Headers` from there
you'll find the `X-Csrf-Ump-Token` and copy the value and paste in .env.

Use any user agent you like.

4. Run it: `go run ./main.go` 

Valid url type: `https://www.dhakacomics.com/digital-comics/comic-name`

## Disclaimer
* This tool was written for educational purposes. I will not be responsible if you use this program in bad faith.
* `dhakacomics-dl` is not affiliated with Dhaka Comics.
