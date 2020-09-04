# properties

[![GoDoc](https://godoc.org/github.com/vbsw/misc/properties?status.svg)](https://godoc.org/github.com/vbsw/misc/properties) [![Go Report Card](https://goreportcard.com/badge/github.com/vbsw/misc/properties)](https://goreportcard.com/report/github.com/vbsw/misc/properties) [![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

## About
This package provides functions to read and write properties. It is published on <https://github.com/vbsw/misc/properties>.

## Copyright
Copyright 2020, Vitali Baumtrok (vbsw@mailbox.org).

properties is distributed under the Boost Software License, version 1.0. (See accompanying file LICENSE or copy at http://www.boost.org/LICENSE_1_0.txt)

properties is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Example

	package main

	import (
		"fmt"
		"github.com/vbsw/misc/properties"
	)

	func main() {
		props, err := properties.ReadFile("some/path/to/file")

		if err == nil {
			for propName, propValue := range props {
				fmt.Println(propName, propValue)
			}
		}
	}

## References
- https://golang.org/doc/install
- https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
- https://docs.oracle.com/cd/E23095_01/Platform.93/ATGProgGuide/html/s0204propertiesfileformat01.html
