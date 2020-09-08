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
	bytes := []byte("aaa; bbb ;ccc;  ddd")
	sep := []byte(";")

	i := scanner.ScanLine(bytes, sep, 0)
	if i != 19 {
		t.Error("offset", i, 19)
	}
	if len(scanner.Begin) != 4 {
		t.Error("field number", len(scanner.Begin), 4)
	} else if scanner.FieldValue(bytes, 0) != "aaa" {
		t.Error("field 0", scanner.FieldValue(bytes, 0), "aaa")
	} else if scanner.FieldValue(bytes, 1) != "bbb" {
		t.Error("field 1", scanner.FieldValue(bytes, 1), "bbb")
	} else if scanner.FieldValue(bytes, 2) != "ccc" {
		t.Error("field 2", scanner.FieldValue(bytes, 2), "ccc")
	} else if scanner.FieldValue(bytes, 3) != "ddd" {
		t.Error("field 3", scanner.FieldValue(bytes, 3), "ddd")
	}
}
