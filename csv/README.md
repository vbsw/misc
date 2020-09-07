# csv

[![GoDoc](https://godoc.org/github.com/vbsw/misc/csv?status.svg)](https://godoc.org/github.com/vbsw/misc/csv)

## About
This package provides functions to read and write CSV. It is published on <https://github.com/vbsw/misc/csv>.

## Example

	package main

	import (
		"fmt"
		"github.com/vbsw/misc/csv"
	)

	func main() {
		columns := []string{"id", "name"}
		csvData, err := csv.ReadFile("some/path/to/file", columns)

		if err == nil && len(csvData) > 1 {
			for i := 0; i < len(csvData[0]); i++ {
				id := csvData[0][i]
				name := csvData[1][i]
				fmt.Println(id, name)
			}
		}
	}

