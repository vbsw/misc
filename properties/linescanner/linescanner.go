/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package linescanner provides functions to scan lines of property files.
package linescanner

import (
	"unicode/utf8"
)

// Line types.
const (
	// LUndefined denotes initial state, i.e. nothing has been parsed.
	LUndefined = 0
	// LEmpty denotes empty line or a line with whitespace, only.
	LEmpty = 1
	// LProperty denotes line with property.
	LProperty = 2
	// LPropertyNext denotes line with property and continuation of property value on next line.
	LPropertyNext = 3
	// LPropertyCont denotes line with continuation of property value from previous line.
	LPropertyCont = 4
	// LPropertyContNext denotes line with continuation of property value from previous line and its continuation on next line.
	LPropertyContNext = 5
	// LComment denotes comment line.
	LComment = 6
	// LUnknownFormat denotes line with something that can not be interpreted.
	LUnknownFormat = 7
)

// LineScanner holds line type and indices for property name and property value.
type LineScanner struct {
	LineType  int
	PropBegin int
	PropEnd   int
	ValBegin  int
	ValEnd    int
}

// ScanLine processes one line searching for property name and property value. Result
// is written to LineScanner fields. Returns offset of next line.
func (scanner *LineScanner) ScanLine(bytes []byte, offset int) int {
	lineEnd, nextLineBegin := seekLineEnd(bytes, offset, len(bytes))
	contentBegin := seekContent(bytes, offset, lineEnd)
	if contentBegin < lineEnd {
		contentEnd := seekContentRight(bytes, contentBegin, lineEnd)
		if isComment(bytes[contentBegin]) {
			scanner.Set(LComment, offset, offset, contentBegin, contentEnd)
		} else if scanner.LineType != LPropertyNext && scanner.LineType != LPropertyContNext {
			scanner.parseProperty(bytes, contentBegin, contentEnd)
		} else {
			scanner.parseValueCont(bytes, contentBegin, contentEnd)
		}
	} else if scanner.LineType == LPropertyNext || scanner.LineType == LPropertyContNext {
		scanner.Set(LPropertyCont, offset, offset, offset, offset)
	} else {
		scanner.Set(LEmpty, offset, offset, offset, offset)
	}
	return nextLineBegin
}

// Set assignes values to LineScanner fields.
func (scanner *LineScanner) Set(lineType, propBegin, propEnd, valBegin, valEnd int) {
	scanner.LineType = lineType
	scanner.PropBegin = propBegin
	scanner.PropEnd = propEnd
	scanner.ValBegin = valBegin
	scanner.ValEnd = valEnd
}

// PropertyName returns property name as string.
func (scanner *LineScanner) PropertyName(bytes, buffer []byte) []byte {
	return convertEscapedBytesToBytes(bytes, buffer, scanner.PropBegin, scanner.PropEnd)
}

// PropertyValue returns property value as string.
func (scanner *LineScanner) PropertyValue(bytes, buffer []byte) []byte {
	return convertEscapedBytesToBytes(bytes, buffer, scanner.ValBegin, scanner.ValEnd)
}

func (scanner *LineScanner) parseProperty(bytes []byte, from, to int) {
	propBegin := from
	asgOpBegin := seekAssignmentOp(bytes, propBegin, to)
	propEnd := seekContentRight(bytes, propBegin, asgOpBegin)
	if asgOpBegin < to {
		valBegin := seekContent(bytes, asgOpBegin+1, to)
		valEnd := seekContentRight(bytes, valBegin, to)
		if isPropValueNext(bytes, valBegin, valEnd) {
			scanner.Set(LPropertyNext, propBegin, propEnd, valBegin, valEnd-1)
		} else {
			scanner.Set(LProperty, propBegin, propEnd, valBegin, valEnd)
		}
	} else if isPropValueNext(bytes, propBegin, propEnd) {
		scanner.Set(LUnknownFormat, propBegin, propEnd, propEnd, propEnd)
	} else {
		scanner.Set(LProperty, propBegin, propEnd, propEnd, propEnd)
	}
}

func (scanner *LineScanner) parseValueCont(bytes []byte, from, to int) {
	if isPropValueNext(bytes, from, to) {
		scanner.Set(LPropertyContNext, from, from, from, to-1)
	} else {
		scanner.Set(LPropertyCont, from, from, from, to)
	}
}

func isComment(b byte) bool {
	return b == '#' || b == '!'
}

func isAssignmentOp(b byte) bool {
	return b == '=' || b == ':'
}

func isPropValueNext(bytes []byte, from, to int) bool {
	return from < to && bytes[to-1] == '\\' && (from+1 == to || bytes[to-2] != '\\')
}

func convertEscapedBytesToBytes(bytes, buffer []byte, from, to int) []byte {
	i, chunkBegin := from, from
	for ; i < to; i++ {
		if bytes[i] == '\\' {
			if chunkBegin < i {
				buffer = append(buffer, bytes[chunkBegin:i]...)
			}
			i, buffer = appendEscapedByte(bytes, buffer, i+1, to)
			chunkBegin = i
			i--
		}
	}
	if chunkBegin < i {
		buffer = append(buffer, bytes[chunkBegin:i]...)
	}
	return buffer
}

func appendEscapedByte(bytes, buffer []byte, from, to int) (int, []byte) {
	if from < to {
		switch bytes[from] {
		case 't':
			buffer = append(buffer, '\t')
		case 'n':
			buffer = append(buffer, '\n')
		case 'r':
			buffer = append(buffer, '\r')
		case 'u':
			from, buffer = appendUnicodeChar(bytes, buffer, from+1, to)
		default:
			buffer = append(buffer, bytes[from])
		}
		return from, buffer
	}
	return to, buffer
}

func appendUnicodeChar(bytes, buffer []byte, from, to int) (int, []byte) {
	end := seekUnicodeCharEnd(bytes, from, to)
	if end > from {
		var r rune
		for i := from; i < to; i++ {
			number := convertCharToInt(bytes[i])
			for j := i + 1; j < to; j++ {
				number *= 16
			}
			r += rune(number)
		}
		buffer = ensureCap(buffer, len(buffer)+4)
		n := utf8.EncodeRune(buffer[len(buffer):len(buffer)+4], r)
		buffer = buffer[:len(buffer)+n]
		return end, buffer
	}
	return end, append(buffer, '?')
}

func convertCharToInt(char byte) int {
	if char >= '0' && char <= '9' {
		return int(char - 48)
	} else if char >= 'A' && char <= 'F' {
		return int(char - 65 + 10)
	}
	return int(char - 97 + 10)
}

func ensureCap(bytes []byte, minCap int) []byte {
	if minCap <= cap(bytes) {
		return bytes
	}
	newBytes := make([]byte, len(bytes), cap(bytes)*2)
	if len(bytes) > 0 {
		copy(newBytes, bytes)
	}
	return newBytes
}
