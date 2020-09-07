/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package properties provides functions to read/write properties.
package properties

import (
	"github.com/vbsw/misc/files"
	"github.com/vbsw/misc/properties/linescanner"
	"github.com/vbsw/misc/ref"
	"io/ioutil"
	"runtime"
)

// Formatting options.
const (
	// Spaces enables spaces around the assignment operator.
	Spaces = 0
	// OpCollon enables collon as the assignment operator.
	OpCollon = 1
	// OpEqual enables equal sign as the assignment operator.
	OpEqual = 2
	// OpSpace enables space as the assignment operator.
	OpSpace = 3
)

// ReadFile reads properties from file. File must be in UTF-8.
func ReadFile(path string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		return ReadBytes(bytes), err
	}
	return nil, err
}

// ReadBytes reads properties from byte array. Content of byte array must be in UTF-8.
func ReadBytes(bytes []byte) map[string]string {
	var scanner linescanner.LineScanner
	var name string
	props := make(map[string]string)
	buffer := make([]byte, 32)
	for i := 0; i < len(bytes); {
		i = scanner.ScanLine(bytes, i)
		if scanner.LineType == linescanner.LProperty {
			buffer = scanner.PropertyName(bytes, buffer[:0])
			name = string(buffer)
			buffer = scanner.PropertyValue(bytes, buffer[:0])
			props[name] = string(buffer)
			name = ""
		} else if scanner.LineType == linescanner.LPropertyNext {
			buffer = scanner.PropertyName(bytes, buffer[:0])
			name = string(buffer)
			buffer = scanner.PropertyValue(bytes, buffer[:0])
		} else if scanner.LineType == linescanner.LPropertyContNext {
			buffer = scanner.PropertyValue(bytes, buffer)
		} else if scanner.LineType == linescanner.LPropertyCont {
			buffer = scanner.PropertyValue(bytes, buffer)
			props[name] = string(buffer)
			name = ""
		} else if scanner.LineType == linescanner.LUnknownFormat && len(name) > 0 {
			props[name] = string(buffer)
			name = ""
		}
	}
	return props
}

// WriteFile writes properties to file.
func WriteFile(path string, propNames, propValues []string) error {
	if len(propNames) > 0 {
		bytes := ToBytes(propNames, propValues)
		err := files.Write(path, bytes)
		return err
	}
	return nil
}

// ToBytes converts properties to byte array.
func ToBytes(propNames, propValues []string, formatting ...int) []byte {
	asgBytes := assignmentBytes(formatting)
	nlBytes := newLineBytes()
	spaces := contains(formatting, Spaces)
	bytes := newPropertiesByteBuffer(propNames, propValues, spaces)
	bytesW := bytes
	for i, propName := range propNames {
		if len(propName) > 0 {
			propValue := propValues[i]
			propNameBytes := ref.Bytes(propName)
			propValueBytes := ref.Bytes(propValue)
			bytesW = writePropertyNameBytes(bytesW, propNameBytes)
			bytesW = writeBytes(bytesW, asgBytes)
			bytesW = writeBytes(bytesW, propValueBytes)
			bytesW = writeBytes(bytesW, nlBytes)
		}
	}
	return bytes
}

func assignmentBytes(formatting []int) []byte {
	var asgOp byte
	var asgBytes []byte
	if contains(formatting, OpCollon) {
		asgOp = ':'
	} else if contains(formatting, OpSpace) {
		asgOp = ' '
	} else {
		asgOp = '='
	}
	if contains(formatting, Spaces) {
		asgBytes = make([]byte, 3)
		asgBytes[0] = ' '
		asgBytes[1] = asgOp
		asgBytes[2] = ' '
	} else {
		asgBytes = make([]byte, 1)
		asgBytes[0] = asgOp
	}
	return asgBytes
}

func newLineBytes() []byte {
	if runtime.GOOS == "windows" {
		return []byte{'\r', '\n'}
	}
	return []byte{'\n'}
}

func contains(list []int, value int) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}
	return false
}

func newPropertiesByteBuffer(propNames, propValues []string, spaces bool) []byte {
	propsBytesNum, linesNum := totalPropsBytesNumber(propNames, propValues)
	escCharsNum := totalEscapedCharsNumber(propNames)
	nlLength := newLineLength()
	asgSpotLength := assignmentSpotLength(spaces)
	bytesLength := propsBytesNum + escCharsNum + (nlLength+asgSpotLength)*linesNum
	bytes := make([]byte, bytesLength)
	return bytes
}

func writePropertyNameBytes(destBytes, srcBytes []byte) []byte {
	j, chunkBegin := 0, 0
	for i, b := range srcBytes {
		if b == ' ' || b == '=' || b == ':' || b == '\\' {
			copy(destBytes[j-(i-chunkBegin):], srcBytes[chunkBegin:i])
			destBytes[j] = '\\'
			destBytes[j+1] = b
			chunkBegin = i + 1
			j++
		}
		j++
	}
	if chunkBegin < len(srcBytes) {
		copy(destBytes[j-(len(srcBytes)-chunkBegin):], srcBytes[chunkBegin:])
	}
	return destBytes[j:]
}

func writeBytes(destBytes, srcBytes []byte) []byte {
	copy(destBytes, srcBytes)
	return destBytes[len(srcBytes):]
}

func totalPropsBytesNumber(propNames, propValues []string) (int, int) {
	var bytesNum, linesNum int
	for i, propName := range propNames {
		if len(propName) > 0 {
			propValue := propValues[i]
			bytesNum += len(propName) + len(propValue)
			linesNum++
		}
	}
	return bytesNum, linesNum
}

func totalEscapedCharsNumber(propNames []string) int {
	var escCharsNum int
	for _, propName := range propNames {
		bytes := ref.Bytes(propName)
		for _, b := range bytes {
			if b == ' ' || b == ':' || b == '=' || b == '\\' {
				escCharsNum++
			}
		}
	}
	return escCharsNum
}

func newLineLength() int {
	if runtime.GOOS == "windows" {
		return 2
	}
	return 1
}

func assignmentSpotLength(spaces bool) int {
	if spaces {
		return 3
	}
	return 1
}
