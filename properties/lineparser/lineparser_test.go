/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package lineparser

import (
	"testing"
)

func TestParseLineA(t *testing.T) {
	var parser LineParser
	bytes := []byte("  \t \na=1\r\n   b = 2 \n  c 4 \\\n  5\\\n  \n d\\\n")

	// empty line
	i := parser.ParseLine(bytes, 0)
	if i != 5 {
		t.Error("offset", i, 5)
	}
	if parser.LineType != LTEmptyLine {
		t.Error("line type", parser.LineType, LTEmptyLine)
	}

	// property a
	i = parser.ParseLine(bytes, i)
	if i != 10 {
		t.Error("offset", i, 10)
	}
	if parser.LineType != LTProperty {
		t.Error("line type", parser.LineType, LTProperty)
	}
	if parser.PropBegin != 5 || parser.PropEnd != 6 {
		t.Error("missmatch property name index", parser.PropBegin, parser.PropEnd, 5, 6)
	}
	if parser.ValBegin != 7 || parser.ValEnd != 8 {
		t.Error("missmatch property value index", parser.ValBegin, parser.ValEnd, 7, 8)
	}
}

func TestParseLineB(t *testing.T) {
	var parser LineParser
	bytes := []byte("  \t \na=1\r\n   b = 2 \n  c 4 \\\n  5\\\n  \n d\\\n")

	i := parser.ParseLine(bytes, 0)
	i = parser.ParseLine(bytes, i)

	// property b
	i = parser.ParseLine(bytes, i)
	if i != 20 {
		t.Error("offset", i, 20)
	}
	if parser.LineType != LTProperty {
		t.Error("line type", parser.LineType, LTProperty)
	}
	if parser.PropBegin != 13 || parser.PropEnd != 14 {
		t.Error("missmatch property name index", parser.PropBegin, parser.PropEnd, 13, 14)
	}
	if parser.ValBegin != 17 || parser.ValEnd != 18 {
		t.Error("missmatch property value index", parser.ValBegin, parser.ValEnd, 17, 18)
	}
}

func TestParseLineC(t *testing.T) {
	var parser LineParser
	bytes := []byte("  \t \na=1\r\n   b = 2 \n  c 4 \\\n  5\\\n  \n d\\\n")

	i := parser.ParseLine(bytes, 0)
	i = parser.ParseLine(bytes, i)
	i = parser.ParseLine(bytes, i)

	// property c
	i = parser.ParseLine(bytes, i)
	if i != 28 {
		t.Error("offset", i, 28)
	}
	if parser.LineType != LTPropertyNext {
		t.Error("line type", parser.LineType, LTPropertyNext)
	}
	if parser.PropBegin != 22 || parser.PropEnd != 23 {
		t.Error("missmatch property name index", parser.PropBegin, parser.PropEnd, 22, 23)
	}
	if parser.ValBegin != 24 || parser.ValEnd != 26 {
		t.Error("missmatch property value index", parser.ValBegin, parser.ValEnd, 24, 26)
	}

	// property c continuation
	i = parser.ParseLine(bytes, i)
	if i != 33 {
		t.Error("offset", i, 33)
	}
	if parser.LineType != LTPropertyContNext {
		t.Error("line type", parser.LineType, LTPropertyContNext)
	}
	if parser.ValBegin != 30 || parser.ValEnd != 31 {
		t.Error("missmatch property value index", parser.ValBegin, parser.ValEnd, 30, 31)
	}

	// property c continuation
	i = parser.ParseLine(bytes, i)
	if i != 36 {
		t.Error("offset", i, 36)
	}
	if parser.LineType != LTPropertyCont {
		t.Error("line type", parser.LineType, LTPropertyCont)
	}
	if parser.ValBegin != parser.ValEnd {
		t.Error("missmatch property value index", parser.ValBegin, parser.ValEnd)
	}
}

func TestParseLineD(t *testing.T) {
	var parser LineParser
	bytes := []byte("  \t \na=1\r\n   b = 2 \n  c 4 \\\n  5\\\n  \n d\\\n")

	i := parser.ParseLine(bytes, 0)
	i = parser.ParseLine(bytes, i)
	i = parser.ParseLine(bytes, i)
	i = parser.ParseLine(bytes, i)
	i = parser.ParseLine(bytes, i)
	i = parser.ParseLine(bytes, i)

	// property d
	i = parser.ParseLine(bytes, i)
	if i != 40 {
		t.Error("offset", i, 40)
	}
	if parser.LineType != LTUnknownFormat {
		t.Error("line type", parser.LineType, LTUnknownFormat)
	}

	// empty line
	i = parser.ParseLine(bytes, i)
	if i != 40 {
		t.Error("offset", i, 40)
	}
	if parser.LineType != LTEmptyLine {
		t.Error("line type", parser.LineType, LTEmptyLine)
	}
}
