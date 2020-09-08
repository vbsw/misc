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
	"github.com/vbsw/misc/ref"
	"io/ioutil"
	"runtime"
)

// Stats holds statistics of CSV data.
type Stats struct {
	HeaderAvailable bool
	FieldsNum       []int
	DataAvailable   []bool
}

// ReadBytes reads CSV data from byte array. Content of byte array must be in UTF-8.
// Returns CSV data as columns. Index 0 is alwas header, even, if
// data has no header, it is copied from the parameter.
func ReadBytes(bytes []byte, header []string, separator string) [][]string {
	var scanner linescanner.LineScanner
	var offset int
	var colMap []int
	csvData := allocCSVData(bytes, header)
	sepr := []byte(separator)
	// seek first line
	offset = scanner.ScanLine(bytes, sepr, 0)
	for offset < len(bytes) && scanner.Empty {
		offset = scanner.ScanLine(bytes, sepr, offset)
	}
	// init header
	colMap, headerAvailable := columnMapping(bytes, scanner.Begin, scanner.End, header)
	if !headerAvailable {
		for i, columnName := range header {
			csvData[i] = append(csvData[i], columnName)
		}
	}
	for i, mapping := range colMap {
		if mapping >= 0 {
			csvData[i] = append(csvData[i], scanner.FieldValue(bytes, mapping))
		} else {
			csvData[i] = append(csvData[i], "")
		}
	}
	// process data lines
	for offset < len(bytes) {
		offset = scanner.ScanLine(bytes, sepr, offset)
		if !scanner.Empty {
			for i, mapping := range colMap {
				if mapping >= 0 {
					csvData[i] = append(csvData[i], scanner.FieldValue(bytes, mapping))
				} else {
					csvData[i] = append(csvData[i], "")
				}
			}
		}
	}
	return csvData
}

// ReadFile reads CSV data from file. File must be in UTF-8.
// Returns CSV data as columns. Index 0 is alwas header, even, if
// data has no header, it is copied from the parameter.
func ReadFile(path string, header []string, separator string) ([][]string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		return ReadBytes(bytes, header, separator), err
	}
	return nil, err
}

// StatsFromBytes scans byte array and returns its CSV data statistics.
func StatsFromBytes(bytes []byte, header []string, separator string) *Stats {
	var scanner linescanner.LineScanner
	var offset int
	apprRowsNum := approximatlyRowsNum(bytes, header)
	stats := &Stats{false, make([]int, 0, apprRowsNum), make([]bool, 0, apprRowsNum)}
	sepr := []byte(separator)
	// seek first line
	offset = scanner.ScanLine(bytes, sepr, 0)
	stats.FieldsNum = append(stats.FieldsNum, len(scanner.Begin))
	stats.DataAvailable = append(stats.DataAvailable, scanner.Empty)
	for offset < len(bytes) && scanner.Empty {
		offset = scanner.ScanLine(bytes, sepr, offset)
		stats.FieldsNum = append(stats.FieldsNum, len(scanner.Begin))
		stats.DataAvailable = append(stats.DataAvailable, scanner.Empty)
	}
	// init header
	_, headerAvailable := columnMapping(bytes, scanner.Begin, scanner.End, header)
	stats.HeaderAvailable = headerAvailable
	// process data lines
	for offset < len(bytes) {
		offset = scanner.ScanLine(bytes, sepr, offset)
		stats.FieldsNum = append(stats.FieldsNum, len(scanner.Begin))
		stats.DataAvailable = append(stats.DataAvailable, scanner.Empty)
	}
	if len(bytes) > 0 && (bytes[len(bytes)-1] == '\n' || bytes[len(bytes)-1] == '\r') {
		stats.FieldsNum = append(stats.FieldsNum, 0)
		stats.DataAvailable = append(stats.DataAvailable, true)
	}
	return stats
}

// StatsFromFile reads CSV data from file, scans it and returns its CSV data
// statistics.
func StatsFromFile(path string, header []string, separator string) (*Stats, error) {
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		return StatsFromBytes(bytes, header, separator), err
	}
	return nil, err
}

// WriteFile writes CSV data to file.
func WriteFile(path string, csvData [][]string, separator string) error {
	if len(csvData) > 0 {
		bytes := ToBytes(csvData, separator)
		err := files.Write(path, bytes)
		return err
	}
	return nil
}

// ToBytes converts CSV data to byte array.
func ToBytes(csvData [][]string, separator string) []byte {
	bytes := allocCSVByteBuffer(csvData, separator)
	if len(bytes) > 0 {
		if runtime.GOOS == "windows" {
			bytes = toBytesWindows(bytes, csvData, separator)
		} else {
			bytes = toBytesOther(bytes, csvData, separator)
		}
	}
	return bytes
}

func allocCSVData(bytes []byte, header []string) [][]string {
	csvData := make([][]string, len(header))
	apprRowsNum := approximatlyRowsNum(bytes, header)
	for i := range csvData {
		csvData[i] = make([]string, 0, apprRowsNum)
	}
	return csvData
}

func approximatlyRowsNum(bytes []byte, header []string) int {
	approximatlyRowsNum := len(bytes) / len(header) / 100
	if approximatlyRowsNum < 8 {
		approximatlyRowsNum = 8
	}
	return approximatlyRowsNum
}

func columnMapping(bytes []byte, begin, end []int, header []string) ([]int, bool) {
	colMap := make([]int, 0, len(header))
	// column mapping
	if len(begin) >= len(header) {
		colMap = columnMappingSupA(bytes, begin, end, colMap, header)
	} else {
		colMap = columnMappingSupB(bytes, begin, end, colMap, header)
	}
	// default mapping
	if len(colMap) == 0 {
		for i := 0; i < len(header); i++ {
			if i < len(begin) {
				colMap = append(colMap, i)
			} else {
				colMap = append(colMap, -1)
			}
		}
		return colMap, false
	}
	return colMap, true
}

func columnMappingSupA(bytes []byte, begin, end, colMap []int, header []string) []int {
	for i, columnName := range header {
		for j := range begin {
			if isEqual(bytes, begin[j], end[j], columnName) {
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

func columnMappingSupB(bytes []byte, begin, end, colMap []int, header []string) []int {
	var count int
	for i, columnName := range header {
		for j := range begin {
			if isEqual(bytes, begin[j], end[j], columnName) {
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

func isEqual(bytes []byte, from, to int, str string) bool {
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

func toBytesWindows(bytes []byte, csvData [][]string, separator string) []byte {
	bytesW := bytes
	sep := ref.Bytes(separator)
	for row := 0; row < len(csvData[0]); row++ {
		for col := 0; col < len(csvData); col++ {
			data := ref.Bytes(csvData[col][row])
			if row > 0 {
				copy(bytesW, sep)
				bytesW = bytesW[len(sep):]
			}
			if len(data) > 0 {
				copy(bytesW, data)
				bytesW = bytesW[len(data):]
			}
		}
		bytesW[0] = '\r'
		bytesW[1] = '\n'
		bytesW = bytesW[2:]
	}
	return bytes
}

func toBytesOther(bytes []byte, csvData [][]string, separator string) []byte {
	bytesW := bytes
	sep := ref.Bytes(separator)
	for row := 0; row < len(csvData[0]); row++ {
		for col := 0; col < len(csvData); col++ {
			data := ref.Bytes(csvData[col][row])
			if col > 0 {
				copy(bytesW, sep)
				bytesW = bytesW[len(sep):]
			}
			if len(data) > 0 {
				copy(bytesW, data)
				bytesW = bytesW[len(data):]
			}
		}
		bytesW[0] = '\n'
		bytesW = bytesW[1:]
	}
	return bytes
}

func allocCSVByteBuffer(csvData [][]string, separator string) []byte {
	var buffer []byte
	if len(csvData) > 0 {
		var bufferLen int
		sepLnLen := (len(csvData) - 1) * len(separator)
		nlLen := newLineLength(runtime.GOOS)
		for row := 0; row < len(csvData[0]); row++ {
			bufferLen += sepLnLen + nlLen
		}
		for col := 0; col < len(csvData); col++ {
			for row := 0; row < len(csvData[col]); row++ {
				bufferLen += len(csvData[col][row])
			}
		}
		buffer = make([]byte, bufferLen)
	}
	return buffer
}

func newLineLength(operatingSystem string) int {
	if operatingSystem == "windows" {
		return 2
	}
	return 1
}
