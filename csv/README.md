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
		header := []string{"id", "name"}
		csvData := csv.New(header, ";")
		err := csvData.ReadFile("some/path/to/file")

		if err == nil && csvData.Size() > 0 {
			for i := 0; i < csvData.Size(); i++ {
				id := csvData.Value(i, 0)
				name := csvData.Value(i, 1)
				fmt.Println(id, name)
			}
		}
	}

