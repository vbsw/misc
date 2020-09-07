/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package csv

import (
	"runtime"
	"strings"
	"testing"
)

func TestToBytes(t *testing.T) {
	nl, nlStr := newLine(runtime.GOOS)
	csvData := [][]string{{"1001", "2002"}, {"alice", "bob"}}
	totalLen := 16 + 2 + len(nl)*2
	lnA := csvData[0][0] + ";" + csvData[1][0]
	lnB := csvData[0][1] + ";" + csvData[1][1]
	bytes := ToBytes(csvData, ";")
	str := strings.Split(string(bytes), nl)
	if len(bytes) != totalLen {
		t.Error(len(bytes))
	} else if len(str) != 3 {
		t.Error(len(str), strings.ReplaceAll(string(bytes), nl, nlStr))
	} else if str[0] != lnA {
		t.Error(str[0], lnA)
	} else if str[1] != lnB {
		t.Error(str[1], lnB)
	}
}

func newLine(operatingSystem string) (string, string) {
	if operatingSystem == "windows" {
		return "\r\n", "\\r\\n"
	}
	return "\n", "\\n"
}
