/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package csv provides functions to read/write CSV data.
package csv

import (
	"github.com/vbsw/misc/files"
	"github.com/vbsw/misc/ref"
	"io/ioutil"
	"runtime"
)

// ReadFile reads CSV data from file. File must be in UTF-8.
// Returns CSV data as columns. Index 0 is alwas header.
func ReadFile(path string, header []string, separator string) ([][]string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		return ReadBytes(bytes, header, separator), err
	}
	return nil, err
}

// ReadBytes reads CSV data from byte array. Content of byte array must be in UTF-8.
// Returns CSV data as columns. Index 0 is alwas header.
func ReadBytes(bytes []byte, header []string, separator string) [][]string {
	data := make([][]string, 0)
	// TODO !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	return data
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

// ToBytes converts strings to byte array.
func ToBytes(csvData [][]string, separator string) []byte {
	bytes := newCSVByteBuffer(csvData, separator)
	if len(bytes) > 0 {
		if runtime.GOOS == "windows" {
			bytes = toBytesWindows(bytes, csvData, separator)
		} else {
			bytes = toBytesOther(bytes, csvData, separator)
		}
	}
	return bytes
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

func newCSVByteBuffer(csvData [][]string, separator string) []byte {
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
