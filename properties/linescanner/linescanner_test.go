/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package linescanner

import (
	"testing"
)

func TestScanLineA(t *testing.T) {
	var scanner LineScanner
	bytes := []byte("  \t \na=1\r\n   b = 2 \n  c 4 \\\n  5\\\n  \n d\\\n")

	// empty line
	i := scanner.ScanLine(bytes, 0)
	if i != 5 {
		t.Error("offset", i, 5)
	}
	if scanner.LineType != LEmpty {
		t.Error("line type", scanner.LineType, LEmpty)
	}

	// property a
	i = scanner.ScanLine(bytes, i)
	if i != 10 {
		t.Error("offset", i, 10)
	}
	if scanner.LineType != LProperty {
		t.Error("line type", scanner.LineType, LProperty)
	}
	if scanner.PropBegin != 5 || scanner.PropEnd != 6 {
		t.Error("missmatch property name index", scanner.PropBegin, scanner.PropEnd, 5, 6)
	}
	if scanner.ValBegin != 7 || scanner.ValEnd != 8 {
		t.Error("missmatch property value index", scanner.ValBegin, scanner.ValEnd, 7, 8)
	}
}

func TestScanLineB(t *testing.T) {
	var scanner LineScanner
	bytes := []byte("  \t \na=1\r\n   b = 2 \n  c 4 \\\n  5\\\n  \n d\\\n")

	i := scanner.ScanLine(bytes, 0)
	i = scanner.ScanLine(bytes, i)

	// property b
	i = scanner.ScanLine(bytes, i)
	if i != 20 {
		t.Error("offset", i, 20)
	}
	if scanner.LineType != LProperty {
		t.Error("line type", scanner.LineType, LProperty)
	}
	if scanner.PropBegin != 13 || scanner.PropEnd != 14 {
		t.Error("missmatch property name index", scanner.PropBegin, scanner.PropEnd, 13, 14)
	}
	if scanner.ValBegin != 17 || scanner.ValEnd != 18 {
		t.Error("missmatch property value index", scanner.ValBegin, scanner.ValEnd, 17, 18)
	}
}

func TestScanLineC(t *testing.T) {
	var scanner LineScanner
	bytes := []byte("  \t \na=1\r\n   b = 2 \n  c 4 \\\n  5\\\n  \n d\\\n")

	i := scanner.ScanLine(bytes, 0)
	i = scanner.ScanLine(bytes, i)
	i = scanner.ScanLine(bytes, i)

	// property c
	i = scanner.ScanLine(bytes, i)
	if i != 28 {
		t.Error("offset", i, 28)
	}
	if scanner.LineType != LPropertyNext {
		t.Error("line type", scanner.LineType, LPropertyNext)
	}
	if scanner.PropBegin != 22 || scanner.PropEnd != 23 {
		t.Error("missmatch property name index", scanner.PropBegin, scanner.PropEnd, 22, 23)
	}
	if scanner.ValBegin != 24 || scanner.ValEnd != 26 {
		t.Error("missmatch property value index", scanner.ValBegin, scanner.ValEnd, 24, 26)
	}

	// property c continuation
	i = scanner.ScanLine(bytes, i)
	if i != 33 {
		t.Error("offset", i, 33)
	}
	if scanner.LineType != LPropertyContNext {
		t.Error("line type", scanner.LineType, LPropertyContNext)
	}
	if scanner.ValBegin != 30 || scanner.ValEnd != 31 {
		t.Error("missmatch property value index", scanner.ValBegin, scanner.ValEnd, 30, 31)
	}

	// property c continuation
	i = scanner.ScanLine(bytes, i)
	if i != 36 {
		t.Error("offset", i, 36)
	}
	if scanner.LineType != LPropertyCont {
		t.Error("line type", scanner.LineType, LPropertyCont)
	}
	if scanner.ValBegin != scanner.ValEnd {
		t.Error("missmatch property value index", scanner.ValBegin, scanner.ValEnd)
	}
}

func TestScanLineD(t *testing.T) {
	var scanner LineScanner
	bytes := []byte("  \t \na=1\r\n   b = 2 \n  c 4 \\\n  5\\\n  \n d\\\n")

	i := scanner.ScanLine(bytes, 0)
	i = scanner.ScanLine(bytes, i)
	i = scanner.ScanLine(bytes, i)
	i = scanner.ScanLine(bytes, i)
	i = scanner.ScanLine(bytes, i)
	i = scanner.ScanLine(bytes, i)

	// property d
	i = scanner.ScanLine(bytes, i)
	if i != 40 {
		t.Error("offset", i, 40)
	}
	if scanner.LineType != LUnknownFormat {
		t.Error("line type", scanner.LineType, LUnknownFormat)
	}

	// empty line
	i = scanner.ScanLine(bytes, i)
	if i != 40 {
		t.Error("offset", i, 40)
	}
	if scanner.LineType != LEmpty {
		t.Error("line type", scanner.LineType, LEmpty)
	}
}
