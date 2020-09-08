/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package linescanner provides functions to parse property files of Java Properties File Format.
package linescanner

import ()

// LineScanner holds indices for fields of a line.
type LineScanner struct {
	Begin []int
	End   []int
	Empty bool
}

// ScanLine processes one line searching for begin and end index of fields.
// Result is written to LineScanner fields. Returns offset of next line.
func (scanner *LineScanner) ScanLine(bytes, separator []byte, offset int) int {
	scanner.Begin = scanner.Begin[:0]
	scanner.End = scanner.End[:0]
	scanner.Empty = true
	lineEnd, nextLineBegin := seekLineEnd(bytes, offset, len(bytes))
	fieldBegin := seekContent(bytes, offset, lineEnd)
	for fieldBegin < lineEnd {
		separatorBegin := seekBytes(bytes, separator, fieldBegin, lineEnd)
		fieldEnd := seekContentRight(bytes, fieldBegin, separatorBegin)
		if fieldBegin < fieldEnd {
			scanner.Begin = append(scanner.Begin, fieldBegin)
			scanner.End = append(scanner.End, fieldEnd)
			scanner.Empty = false
		} else {
			scanner.Begin = append(scanner.Begin, fieldBegin)
			scanner.End = append(scanner.End, fieldBegin)
		}
		fieldBegin = seekContent(bytes, separatorBegin+len(separator), lineEnd)
	}
	return nextLineBegin
}

// FieldValue returns field value as string.
func (scanner *LineScanner) FieldValue(bytes []byte, index int) string {
	return string(bytes[scanner.Begin[index]:scanner.End[index]])
}
