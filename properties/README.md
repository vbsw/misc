# properties

[![GoDoc](https://godoc.org/github.com/vbsw/misc/properties?status.svg)](https://godoc.org/github.com/vbsw/misc/properties)

## About
This package provides functions to read and write properties. It is published on <https://github.com/vbsw/misc/properties>.

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
- https://docs.oracle.com/cd/E23095_01/Platform.93/ATGProgGuide/html/s0204propertiesfileformat01.html
