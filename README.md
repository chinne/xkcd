# XKCD scraper

This tool is scraping the latest image from [xkcd](https://xkcd.com/).

## Installation
To build the Go program in Windows, Linux, and macOS, you can use the go build command.
This command will compile the Go source code into an executable binary file that you can run on the target platform.

The Go source code is in the file named main.go, you can build it for Windows with the following command:

```shell
go build -o xkcd.exe main.go
```

This will produce an executable file named xkcd.exe in the current directory. You can then run this file on a Windows
machine to launch your application.

To build the same program for Linux and macOS[^1], you can use the following command:

```shell
go build -o myapp main.go
```

This will produce an executable file named myapp in the current directory. You can then run this file on a Linux or Mac
machine to launch your application.

[^1]: Note that Go does not support building for macOS on Windows or Linux systems. To build a macOS application, you
must use a macOS system.

### Note: robots.txt

The [robots.txt](https://xkcd.com/robots.txt) allows:
```text
User-agent: *
Disallow: /personal/
```
The robots.txt file is a text file that webmasters create to instruct web robots (typically search engine robots) on
which pages or files the robot should not access. In this case, the User-agent: * line applies the rule to all robots,
and the Disallow: /personal/ line tells them not to access any pages or files in the /personal/ directory. This means
that any robot that accesses the website will not be able to crawl or index any pages or files in the /personal/
directory.

