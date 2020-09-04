/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package propertiesparser provides functions to parse property files of Java Properties File Format.
package propertiesparser

import (
	"github.com/vbsw/misc/slices/ensurecap"
	"unicode/utf8"
)

// Line types (LT).
const (
	// LTUndefined denotes initial state, i.e. nothing has been parsed.
	LTUndefined = 0
	// LTEmptyLine denotes empty line or a line with whitespace, only.
	LTEmptyLine = 1
	// LTProperty denotes line with property.
	LTProperty = 2
	// LTPropertyNext denotes line with property and continuation of property value on next line.
	LTPropertyNext = 3
	// LTPropertyCont denotes line with continuation of property value from previous line.
	LTPropertyCont = 4
	// LTPropertyContNext denotes line with continuation of property value from previous line and its continuation on next line.
	LTPropertyContNext = 5
	// LTComment denotes comment line.
	LTComment = 6
	// LTUnknownFormat denotes line with something that can not be interpreted.
	LTUnknownFormat = 7
)

// LineParser holds line type and indices for property name and property value.
type LineParser struct {
	LineType  int
	PropBegin int
	PropEnd   int
	ValBegin  int
	ValEnd    int
}

// ParseLine processes one line searching for property name and property value. Result
// is written to LineParser fields. Returns offset of next line.
func (parser *LineParser) ParseLine(bytes []byte, offset int) int {
	lineEnd, nextLineBegin := seekLineEnd(bytes, offset, len(bytes))
	contentBegin := seekContent(bytes, offset, lineEnd)
	if contentBegin < lineEnd {
		contentEnd := seekContentRight(bytes, contentBegin, lineEnd)
		if isComment(bytes[contentBegin]) {
			parser.Set(LTComment, offset, offset, contentBegin, contentEnd)
		} else if parser.LineType != LTPropertyNext && parser.LineType != LTPropertyContNext {
			parser.parseProperty(bytes, contentBegin, contentEnd)
		} else {
			parser.parseValueCont(bytes, contentBegin, contentEnd)
		}
	} else if parser.LineType == LTPropertyNext || parser.LineType == LTPropertyContNext {
		parser.Set(LTPropertyCont, offset, offset, offset, offset)
	} else {
		parser.Set(LTEmptyLine, offset, offset, offset, offset)
	}
	return nextLineBegin
}

// Set assignes values to LineParser fields.
func (parser *LineParser) Set(lineType, propBegin, propEnd, valBegin, valEnd int) {
	parser.LineType = lineType
	parser.PropBegin = propBegin
	parser.PropEnd = propEnd
	parser.ValBegin = valBegin
	parser.ValEnd = valEnd
}

// PropertyName returns property name as string.
func (parser *LineParser) PropertyName(bytes, buffer []byte) []byte {
	return convertEscapedBytesToBytes(bytes, buffer, parser.PropBegin, parser.PropEnd)
}

// PropertyValue returns property value as string.
func (parser *LineParser) PropertyValue(bytes, buffer []byte) []byte {
	return convertEscapedBytesToBytes(bytes, buffer, parser.ValBegin, parser.ValEnd)
}

func (parser *LineParser) parseProperty(bytes []byte, from, to int) {
	propBegin := from
	asgOpBegin := seekAssignmentOp(bytes, propBegin, to)
	propEnd := seekContentRight(bytes, propBegin, asgOpBegin)
	if asgOpBegin < to {
		valBegin := seekContent(bytes, asgOpBegin+1, to)
		valEnd := seekContentRight(bytes, valBegin, to)
		if isPropValueNext(bytes, valBegin, valEnd) {
			parser.Set(LTPropertyNext, propBegin, propEnd, valBegin, valEnd-1)
		} else {
			parser.Set(LTProperty, propBegin, propEnd, valBegin, valEnd)
		}
	} else if isPropValueNext(bytes, propBegin, propEnd) {
		parser.Set(LTUnknownFormat, propBegin, propEnd, propEnd, propEnd)
	} else {
		parser.Set(LTProperty, propBegin, propEnd, propEnd, propEnd)
	}
}

func (parser *LineParser) parseValueCont(bytes []byte, from, to int) {
	if isPropValueNext(bytes, from, to) {
		parser.Set(LTPropertyContNext, from, from, from, to-1)
	} else {
		parser.Set(LTPropertyCont, from, from, from, to)
	}
}

func seekLineEnd(bytes []byte, from, to int) (int, int) {
	for i := from; i < to; i++ {
		if bytes[i] == '\n' {
			return i, i + 1
		} else if bytes[i] == '\r' {
			if i+1 < to && bytes[i+1] == '\n' {
				return i, i + 2
			}
			return i, i + 1
		}
	}
	return to, to
}

func seekContent(bytes []byte, from, to int) int {
	for i := from; i < to; i++ {
		if bytes[i] < 0 || bytes[i] > 32 {
			return i
		}
	}
	return to
}

func seekContentRight(bytes []byte, from, to int) int {
	for i := to - 1; i >= from; i-- {
		if bytes[i] < 0 || bytes[i] > 32 {
			return i + 1
		}
	}
	return from
}

func seekAssignmentOp(bytes []byte, from, to int) int {
	for i := from; i < to; i++ {
		if bytes[i] == '\\' {
			i++
		} else if isAssignmentOp(bytes[i]) {
			return i
		} else if bytes[i] == ' ' {
			contentBegin := seekContent(bytes, i, to)
			if contentBegin < to && isAssignmentOp(bytes[contentBegin]) {
				return contentBegin
			}
			return i
		}
	}
	return to
}

func seekUnicodeCharEnd(bytes []byte, from, to int) int {
	i := from
	for ; i < to && i < from+4; i++ {
		b := bytes[i]
		if (b < '0' || b > '9') && (b < 'A' || b > 'F') && (b < 'a' || b > 'b') {
			break
		}
	}
	return i
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
		buffer = ensurecap.Byte(buffer, len(buffer)+4, newCap)
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

func newCap(oldLen, oldCap, reqCap int) int {
	return oldCap * 2
}
