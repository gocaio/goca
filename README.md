![](img/Goca_banner.png)

[![](https://img.shields.io/badge/Chat-Telegram-blue.svg?style=plastic&logo=telegram)](https://t.me/GocaIO) 
[![Build Status](https://travis-ci.org/gocaio/goca.svg?branch=dev)](https://travis-ci.org/gocaio/goca)
[![GitHub Issues](https://img.shields.io/github/issues/gocaio/goca.svg)](https://github.com/gocaio/goca/issues)
[![GitHub tag](https://img.shields.io/github/tag/gocaio/goca.svg)](https://github.com/gocaio/goca/tags)
[![Go Version](https://img.shields.io/badge/go-1.12-blue.svg?logo=go)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Contribute Yes](https://img.shields.io/badge/contribute-yes-brightgreen.svg)](https://github.com/gocaio/goca/blob/master/CONTRIBUTING.md)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgocaio%2Fgoca.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgocaio%2Fgoca?ref=badge_shield)

**Goca** is a [FOCA](https://github.com/ElevenPaths/FOCA) fork written in Go, which is a tool used mainly to find metadata and hidden information in the documents its scans. These documents may be on web pages, and can be downloaded and analyzed with **Goca**.

It is capable of analyzing a wide variety of documents, with the most common being Microsoft Office, Open Office, or PDF files, although it also analyzes Adobe InDesign or SVG files, for instance.

These documents are searched for using search engines such as:

+ [x] Google
+ [ ] Bing
+ [ ] DuckDuckGo
+ [ ] Yahoo
+ [ ] Ask

Then downloads the documents and extracts the EXIF information from graphic files, and a complete analysis of the information discovered through the URL is conducted even before downloading the file.

***

## USAGE

Download built packages from [**Releases**](https://github.com/gocaio/goca/releases)

To build from source, you will need Go installed.

```bash
$ export GO111MODULE=on 
$ go get ./...
$ go run goca/goca.go -h
```

To run Goca from Docker:

```bash
$ docker build -t gocaio/goca /path/to/goca
$ docker run gocaio/goca -h
```

***

## Contributing Guide

Please reade the Contributing guide:

+ [Contributing](CONTRIBUTING.md)

***

## Documentation

Refer to the [Official Doc](https://github.com/gocaio/Doc).


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgocaio%2Fgoca.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgocaio%2Fgoca?ref=badge_large)