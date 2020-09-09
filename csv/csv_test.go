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

func TestAppend(t *testing.T) {
	header := []string{"a", "b", "c"}
	separator := ";"
	csvData := New(header, separator)

	if csvData.Size() != 0 {
		t.Error(csvData.Size(), 0)
	}

	csvData.Append("1", "2")
	if csvData.Size() != 1 {
		t.Error(csvData.Size(), 1)
	} else if csvData.Value(0, 0) != "1" {
		t.Error(csvData.Value(0, 0), "1")
	} else if csvData.Value(0, 2) != "" {
		t.Error(csvData.Value(0, 2))
	}

	csvData.Append("3", "", "4")
	if csvData.Size() != 2 {
		t.Error(csvData.Size(), 2)
	} else if csvData.Value(1, 0) != "3" {
		t.Error(csvData.Value(1, 0), "3")
	} else if csvData.Value(1, 2) != "4" {
		t.Error(csvData.Value(1, 2), "4")
	}
}

func TestInsert(t *testing.T) {
	header := []string{"a", "b", "c"}
	separator := ";"
	csvData := New(header, separator)

	if csvData.Size() != 0 {
		t.Error(csvData.Size(), 0)
	}

	csvData.Insert(0, "1", "2")
	if csvData.Size() != 1 {
		t.Error(csvData.Size(), 1)
	} else if csvData.Value(0, 0) != "1" {
		t.Error(csvData.Value(0, 0), "1")
	} else if csvData.Value(0, 2) != "" {
		t.Error(csvData.Value(0, 2))
	}

	csvData.Insert(0, "3", "", "4")
	if csvData.Size() != 2 {
		t.Error(csvData.Size(), 2)
	} else if csvData.Value(0, 0) != "3" {
		t.Error(csvData.Value(0, 0), "3")
	} else if csvData.Value(0, 2) != "4" {
		t.Error(csvData.Value(0, 2), "4")
	} else if csvData.Value(1, 0) != "1" {
		t.Error(csvData.Value(1, 0), "1")
	} else if csvData.Value(1, 2) != "" {
		t.Error(csvData.Value(1, 2))
	}
}

func TestBytes(t *testing.T) {
	header := []string{"alice", "bob"}
	separator := ";"
	nl, nlStr := newLine()
	csvData := New(header, separator)
	csvData.Append("1001", "1002")
	csvData.Append("2001", "2002")

	strOrig := "alice;bob" + nl + "1001;1002" + nl + "2001;2002" + nl
	str := string(csvData.Bytes(true))

	if str != strOrig {
		t.Error(strings.ReplaceAll(string(str), nl, nlStr))
	}
}

func TestHeaderMappingA(t *testing.T) {
	header := []string{"alice", "bob"}
	separator := ";"
	nl, _ := newLine()
	strOrig := "alice;bob" + nl
	csvData := New(header, separator)

	mapping, isData := csvData.headerMapping([]byte(strOrig), []int{0, 6}, []int{5, 9})
	if isData != false {
		t.Error(isData, false)
	} else if len(mapping) != 2 {
		t.Error(len(mapping), 2)
	} else if mapping[0] != 0 {
		t.Error(mapping[0], 0)
	} else if mapping[1] != 1 {
		t.Error(mapping[1], 1)
	}
}

func TestHeaderMappingB(t *testing.T) {
	header := []string{"alice", "bob", "claire"}
	separator := ";"
	nl, _ := newLine()
	strOrig := "alice;bob" + nl
	csvData := New(header, separator)

	mapping, isData := csvData.headerMapping([]byte(strOrig), []int{0, 6}, []int{5, 9})
	if isData != false {
		t.Error(isData, false)
	} else if len(mapping) != 3 {
		t.Error(len(mapping), 3)
	} else if mapping[0] != 0 {
		t.Error(mapping[0], 0)
	} else if mapping[1] != 1 {
		t.Error(mapping[1], 1)
	} else if mapping[2] >= 0 && mapping[2] < 2 {
		t.Error(mapping[2])
	}
}

func TestHeaderMappingC(t *testing.T) {
	header := []string{"aliceX", "bobX", "claireX"}
	separator := ";"
	nl, _ := newLine()
	strOrig := "alice;bob" + nl
	csvData := New(header, separator)

	mapping, isData := csvData.headerMapping([]byte(strOrig), []int{0, 6}, []int{5, 9})
	if isData != true {
		t.Error(isData, true)
	} else if len(mapping) != 3 {
		t.Error(len(mapping), 3)
	} else if mapping[0] != 0 {
		t.Error(mapping[0], 0)
	} else if mapping[1] != 1 {
		t.Error(mapping[1], 1)
	} else if mapping[2] >= 0 && mapping[2] < 2 {
		t.Error(mapping[2])
	}
}

func TestReadBytesA(t *testing.T) {
	header := []string{"alice", "bob"}
	separator := ";"
	nl, _ := newLine()
	strOrig := "alice;bob" + nl + "1001;1002" + nl + "2001;2002" + nl
	csvData := New(header, separator)

	csvData.ReadBytes([]byte(strOrig))
	if csvData.Size() != 2 || len(csvData.Columns) != 2 {
		t.Error(csvData.Size(), len(csvData.Columns), 2, 2)
	} else if csvData.Value(0, 0) != "1001" {
		t.Error(csvData.Value(0, 0), "1001")
	} else if csvData.Value(0, 1) != "1002" {
		t.Error(csvData.Value(0, 1), "1002")
	} else if csvData.Value(1, 0) != "2001" {
		t.Error(csvData.Value(1, 0), "2001")
	} else if csvData.Value(1, 1) != "2002" {
		t.Error(csvData.Value(1, 1), "2002")
	} else if csvData.LineNumbers[0] != 2 {
		t.Error(csvData.LineNumbers[0], 2)
	} else if csvData.LineNumbers[1] != 3 {
		t.Error(csvData.LineNumbers[1], 3)
	}
}

func TestReadBytesB(t *testing.T) {
	header := []string{"alice", "bob"}
	separator := ";"
	nl, _ := newLine()
	strOrig := "bob;alice" + nl + "1002;1001" + nl + "2002;2001" + nl
	csvData := New(header, separator)

	csvData.ReadBytes([]byte(strOrig))
	if csvData.Size() != 2 || len(csvData.Columns) != 2 {
		t.Error(csvData.Size(), len(csvData.Columns), 2, 2)
	} else if csvData.Value(0, 0) != "1001" {
		t.Error(csvData.Value(0, 0), "1001")
	} else if csvData.Value(0, 1) != "1002" {
		t.Error(csvData.Value(0, 1), "1002")
	} else if csvData.Value(1, 0) != "2001" {
		t.Error(csvData.Value(1, 0), "2001")
	} else if csvData.Value(1, 1) != "2002" {
		t.Error(csvData.Value(1, 1), "2002")
	} else if csvData.LineNumbers[0] != 2 {
		t.Error(csvData.LineNumbers[0], 2)
	} else if csvData.LineNumbers[1] != 3 {
		t.Error(csvData.LineNumbers[1], 3)
	}
}

func TestReadBytesC(t *testing.T) {
	header := []string{"alice", "bob"}
	separator := ";"
	nl, _ := newLine()
	strOrig := nl + nl + "alice;bob" + nl + ";" + nl + "1001;1002" + nl + "2001;2002" + nl
	csvData := New(header, separator)

	csvData.ReadBytes([]byte(strOrig))
	if csvData.Size() != 2 || len(csvData.Columns) != 2 {
		t.Error(csvData.Size(), len(csvData.Columns), 2, 2)
	} else if csvData.Value(0, 0) != "1001" {
		t.Error(csvData.Value(0, 0), "1001")
	} else if csvData.Value(0, 1) != "1002" {
		t.Error(csvData.Value(0, 1), "1002")
	} else if csvData.Value(1, 0) != "2001" {
		t.Error(csvData.Value(1, 0), "2001")
	} else if csvData.Value(1, 1) != "2002" {
		t.Error(csvData.Value(1, 1), "2002")
	} else if csvData.LineNumbers[0] != 5 {
		t.Error(csvData.LineNumbers[0], 5)
	} else if csvData.LineNumbers[1] != 6 {
		t.Error(csvData.LineNumbers[1], 6)
	}
}

func TestReadBytesD(t *testing.T) {
	header := []string{"alice", "bob", "clair"}
	separator := ";"
	nl, _ := newLine()
	strOrig := "bob;alice" + nl + "1002;1001" + nl + "2002;2001" + nl
	csvData := New(header, separator)

	csvData.ReadBytes([]byte(strOrig))
	if csvData.Size() != 2 || len(csvData.Columns) != 3 {
		t.Error(csvData.Size(), len(csvData.Columns), 2, 3)
	} else if csvData.Value(0, 0) != "1001" {
		t.Error(csvData.Value(0, 0), "1001")
	} else if csvData.Value(0, 1) != "1002" {
		t.Error(csvData.Value(0, 1), "1002")
	} else if csvData.Value(0, 2) != "" {
		t.Error(csvData.Value(0, 2), "")
	} else if csvData.Value(1, 0) != "2001" {
		t.Error(csvData.Value(1, 0), "2001")
	} else if csvData.Value(1, 1) != "2002" {
		t.Error(csvData.Value(1, 1), "2002")
	} else if csvData.Value(1, 2) != "" {
		t.Error(csvData.Value(1, 2), "")
	}
}

func newLine() (string, string) {
	if runtime.GOOS == "windows" {
		return "\r\n", "\\r\\n"
	}
	return "\n", "\\n"
}
