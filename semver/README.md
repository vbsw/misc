# semver

[![GoDoc](https://godoc.org/github.com/vbsw/misc/semver?status.svg)](https://godoc.org/github.com/vbsw/misc/semver)

## About
semver is a Go package that provides functions to create version numbers as defined on <http://semver.org>. semver is published on <https://github.com/vbsw/misc/semver>.

## Example

Source:

	package main

	import (
		"fmt"
		"github.com/vbsw/misc/semver"
	)

	func main() {
		version = semver.New(1, 0, 2, "", "")
		fmt.Println("version:", version)
	}

Output:

	version: 1.0.2

## References
- <http://semver.org>
