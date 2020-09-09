/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package csv provides functions to read/write CSV data.
package csv

import (
	"github.com/vbsw/misc/csv/linescanner"
	"github.com/vbsw/misc/files"
	"github.com/vbsw/misc/insert"
	"github.com/vbsw/misc/ref"
	"github.com/vbsw/misc/remove"
	"io/ioutil"
	"runtime"
)

// CSV holds properties to read/write files in CSV format.
type CSV struct {
	Header      []string
	Separator   string
	Columns     [][]string
	LineNumbers []int
}

// New returns a new instance of CSV.
func New(header []string, separator string) *CSV {
	csvData := new(CSV)
	csvData.Header = header
	csvData.Separator = separator
	csvData.Columns = make([][]string, len(header))
	csvData.LineNumbers = make([]int, 0, 16)
	for i := range csvData.Columns {
		csvData.Columns[i] = make([]string, 0, 16)
	}
	return csvData
}

// Append appends values as a new row.
func (csvData *CSV) Append(values ...string) {
	for i := range csvData.Columns {
		if i < len(values) {
			csvData.Columns[i] = append(csvData.Columns[i], values[i])
		} else {
			csvData.Columns[i] = append(csvData.Columns[i], "")
		}
	}
	csvData.LineNumbers = append(csvData.LineNumbers, 0)
}

// Bytes returns CSV data as byte array.
func (csvData *CSV) Bytes(includeHeader bool) []byte {
	size := csvData.neededDataSize()
	if includeHeader {
		size += csvData.neededHeaderSize()
	}
	bytes := make([]byte, size)
	sepBytes := ref.Bytes(csvData.Separator)
	nlBytes := csvData.newLineBytes()
	bytesW := bytes
	if includeHeader {
		bytesW = csvData.writeHeader(bytesW, sepBytes, nlBytes)
	}
	if len(csvData.Columns) > 0 {
		for row := range csvData.Columns[0] {
			bytesW = csvData.writeData(bytesW, sepBytes, nlBytes, row)
		}
	}
	return bytes
}

// Clear deletes all rows.
func (csvData *CSV) Clear() {
	csvData.LineNumbers = csvData.LineNumbers[:0]
	for i := range csvData.Columns {
		csvData.Columns[i] = csvData.Columns[i][:0]
	}
}

// Insert insert values as a row.
func (csvData *CSV) Insert(row int, values ...string) {
	csvData.LineNumbers = insert.Int(csvData.LineNumbers, row, 0)
	for i := range csvData.Columns {
		if i < len(values) {
			csvData.Columns[i] = insert.String(csvData.Columns[i], row, values[i])
		} else {
			csvData.Columns[i] = insert.String(csvData.Columns[i], row, "")
		}
	}
}

// ReadBytes reads CSV data from byte array. Content of byte array must be in UTF-8.
func (csvData *CSV) ReadBytes(bytes []byte) {
	var scanner linescanner.LineScanner
	seprBytes := ref.Bytes(csvData.Separator)
	offset := scanner.ScanLine(bytes, seprBytes, 0)
	line := 1
	for offset < len(bytes) && scanner.Empty {
		offset = scanner.ScanLine(bytes, seprBytes, offset)
		line++
	}
	mapping, isData := csvData.headerMapping(bytes, scanner.Begin, scanner.End)
	if isData {
		for col := range csvData.Header {
			field := scanner.FieldValue(bytes, mapping[col])
			csvData.Columns[col] = append(csvData.Columns[col], field)
		}
		csvData.LineNumbers = append(csvData.LineNumbers, line)
	}
	for offset < len(bytes) {
		offset = scanner.ScanLine(bytes, seprBytes, offset)
		line++
		if !scanner.Empty {
			for col := range csvData.Header {
				field := scanner.FieldValue(bytes, mapping[col])
				csvData.Columns[col] = append(csvData.Columns[col], field)
			}
			csvData.LineNumbers = append(csvData.LineNumbers, line)
		}
	}
}

// ReadFile reads CSV data from file. File must be in UTF-8.
func (csvData *CSV) ReadFile(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		csvData.ReadBytes(bytes)
	}
	return err
}

// Remove removes a row.
func (csvData *CSV) Remove(row int) {
	csvData.LineNumbers = remove.Int(csvData.LineNumbers, row)
	for i := range csvData.Columns {
		csvData.Columns[i] = remove.String(csvData.Columns[i], row)
	}
}

// Set overwrites a row with values.
func (csvData *CSV) Set(row int, values ...string) {
	for i := range csvData.Columns {
		if i < len(values) {
			csvData.Columns[i][row] = values[i]
		} else {
			csvData.Columns[i][row] = ""
		}
	}
}

// Size returns number of rows.
func (csvData *CSV) Size() int {
	if len(csvData.Columns) > 0 {
		return len(csvData.Columns[0])
	}
	return 0
}

// Value returns value at specified row and column.
func (csvData *CSV) Value(row, column int) string {
	return csvData.Columns[column][row]
}

// WriteFile writes CSV data to file.
func (csvData *CSV) WriteFile(path string, includeHeader bool) error {
	bytes := csvData.Bytes(includeHeader)
	err := files.Write(path, bytes)
	return err
}

func (csvData *CSV) headerMapping(bytes []byte, begin, end []int) ([]int, bool) {
	mapping := make([]int, 0, len(csvData.Header))
	// column mapping
	if len(begin) >= len(csvData.Header) {
		mapping = csvData.headerMappingSupA(bytes, begin, end, mapping)
	} else {
		mapping = csvData.headerMappingSupB(bytes, begin, end, mapping)
	}
	// default mapping
	if len(mapping) == 0 {
		for i := 0; i < len(csvData.Header); i++ {
			if i < len(begin) {
				mapping = append(mapping, i)
			} else {
				mapping = append(mapping, -1)
			}
		}
		return mapping, true
	}
	return mapping, false
}

func (csvData *CSV) headerMappingSupA(bytes []byte, begin, end, colMap []int) []int {
	for i, columnName := range csvData.Header {
		for j := range begin {
			if csvData.isEqual(bytes, begin[j], end[j], columnName) {
				colMap = append(colMap, j)
				break
			}
		}
		// corresponding column wasn't found
		if len(colMap) != i+1 {
			colMap = colMap[:0]
			break
		}
	}
	return colMap
}

func (csvData *CSV) headerMappingSupB(bytes []byte, begin, end, colMap []int) []int {
	var count int
	for i, columnName := range csvData.Header {
		for j := range begin {
			if csvData.isEqual(bytes, begin[j], end[j], columnName) {
				colMap = append(colMap, j)
				count++
				break
			}
		}
		// corresponding column wasn't found
		if len(colMap) != i+1 {
			colMap = append(colMap, -1)
		}
	}
	if count != len(begin) {
		return colMap[:0]
	}
	return colMap
}

func (csvData *CSV) isEqual(bytes []byte, from, to int, str string) bool {
	if len(str) == to-from {
		strBytes := ref.Bytes(str)
		for i, b := range strBytes {
			if bytes[from+i] != b {
				return false
			}
		}
		return true
	}
	return false
}

func (csvData *CSV) neededDataSize() int {
	var size int
	if len(csvData.Columns) > 0 {
		nonDataSize := csvData.nonDataLineSize()
		for range csvData.Columns[0] {
			size += nonDataSize
		}
		for _, columnData := range csvData.Columns {
			for _, field := range columnData {
				size += len(field)
			}
		}
	}
	return size
}

func (csvData *CSV) neededHeaderSize() int {
	var size int
	if len(csvData.Header) > 0 {
		size = csvData.nonDataLineSize()
		for _, field := range csvData.Header {
			size += len(field)
		}
	}
	return size
}

func (csvData *CSV) newLineBytes() []byte {
	if runtime.GOOS == "windows" {
		return []byte{'\r', '\n'}
	}
	return []byte{'\n'}
}

func (csvData *CSV) nonDataLineSize() int {
	var newLineSize int
	separatorLineSize := (len(csvData.Header) - 1) * len(csvData.Separator)
	if runtime.GOOS == "windows" {
		newLineSize = 2
	} else {
		newLineSize = 1
	}
	return separatorLineSize + newLineSize
}

func (csvData *CSV) writeData(bytes, sepBytes, nlBytes []byte, row int) []byte {
	for col, column := range csvData.Columns {
		if col > 0 {
			copy(bytes, sepBytes)
			bytes = bytes[len(sepBytes):]
		}
		copy(bytes, ref.Bytes(column[row]))
		bytes = bytes[len(column[row]):]
	}
	copy(bytes, nlBytes)
	return bytes[len(nlBytes):]
}

func (csvData *CSV) writeHeader(bytes, sepBytes, nlBytes []byte) []byte {
	for col, columnName := range csvData.Header {
		if col > 0 {
			copy(bytes, sepBytes)
			bytes = bytes[len(sepBytes):]
		}
		copy(bytes, ref.Bytes(columnName))
		bytes = bytes[len(columnName):]
	}
	copy(bytes, nlBytes)
	return bytes[len(nlBytes):]
}
