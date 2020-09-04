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
	"github.com/vbsw/misc/parser/propertiesparser"
	"github.com/vbsw/misc/ref"
	"github.com/vbsw/misc/slices/contains"
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
	var parser propertiesparser.LineParser
	var name string
	props := make(map[string]string)
	buffer := make([]byte, 32)
	for i := 0; i < len(bytes); {
		i = parser.ParseLine(bytes, i)
		if parser.LineType == propertiesparser.LTProperty {
			buffer = parser.PropertyName(bytes, buffer[:0])
			name = string(buffer)
			buffer = parser.PropertyValue(bytes, buffer[:0])
			props[name] = string(buffer)
			name = ""
		} else if parser.LineType == propertiesparser.LTPropertyNext {
			buffer = parser.PropertyName(bytes, buffer[:0])
			name = string(buffer)
			buffer = parser.PropertyValue(bytes, buffer[:0])
		} else if parser.LineType == propertiesparser.LTPropertyContNext {
			buffer = parser.PropertyValue(bytes, buffer)
		} else if parser.LineType == propertiesparser.LTPropertyCont {
			buffer = parser.PropertyValue(bytes, buffer)
			props[name] = string(buffer)
			name = ""
		} else if parser.LineType == propertiesparser.LTUnknownFormat && len(name) > 0 {
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
	asgSpotBytes := assignmentSpotBytes(formatting)
	nlBytes := newLineBytes()
	spaces := contains.Int(formatting, Spaces)
	bytes := newPropertiesByteBuffer(propNames, propValues, spaces)
	bytesW := bytes
	for i, propName := range propNames {
		if len(propName) > 0 {
			propValue := propValues[i]
			propNameBytes := ref.Bytes(propName)
			propValueBytes := ref.Bytes(propValue)
			bytesW = writePropertyNameBytes(bytesW, propNameBytes)
			bytesW = writeBytes(bytesW, asgSpotBytes)
			bytesW = writeBytes(bytesW, propValueBytes)
			bytesW = writeBytes(bytesW, nlBytes)
		}
	}
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

func assignmentSpotBytes(formatting []int) []byte {
	var asgOp byte
	var asgSpotBytes []byte
	if contains.Int(formatting, OpCollon) {
		asgOp = ':'
	} else if contains.Int(formatting, OpSpace) {
		asgOp = ' '
	} else {
		asgOp = '='
	}
	if contains.Int(formatting, Spaces) {
		asgSpotBytes = make([]byte, 3)
		asgSpotBytes[0] = ' '
		asgSpotBytes[1] = asgOp
		asgSpotBytes[2] = ' '
	} else {
		asgSpotBytes = make([]byte, 1)
		asgSpotBytes[0] = asgOp
	}
	return asgSpotBytes
}

func newLineBytes() []byte {
	if runtime.GOOS == "windows" {
		return []byte{'\r', '\n'}
	}
	return []byte{'\n'}
}

func newPropertiesByteBuffer(propNames, propValues []string, spaces bool) []byte {
	propsBytesNum, linesNum := totalPropsBytesNumber(propNames, propValues)
	escCharsNum := totalEscapedCharsNumber(propNames)
	nlLength := newLineLength(runtime.GOOS)
	asgSpotLength := assignmentSpotLength(spaces)
	bytesLength := propsBytesNum + escCharsNum + (nlLength+asgSpotLength)*linesNum
	bytes := make([]byte, bytesLength)
	return bytes
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

func newLineLength(operatingSystem string) int {
	if operatingSystem == "windows" {
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
