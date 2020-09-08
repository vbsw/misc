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

func TestStatsBytes(t *testing.T) {
	str := "a;b;c\n1;2;3\n \n4;5\n"
	header := []string{"a", "b", "c"}
	strBytes := []byte(str)
	stats := StatsFromBytes(strBytes, header, ";")
	if len(stats.FieldsNum) != 5 {
		t.Error(len(stats.FieldsNum), 5)
	} else if stats.HeaderAvailable == false {
		t.Error(stats.HeaderAvailable, true)
	} else if stats.FieldsNum[0] != 3 {
		t.Error(stats.FieldsNum[0], 3)
	} else if stats.FieldsNum[1] != 3 {
		t.Error(stats.FieldsNum[1], 3)
	} else if stats.FieldsNum[2] != 0 {
		t.Error(stats.FieldsNum[2], 0)
	} else if stats.FieldsNum[3] != 2 {
		t.Error(stats.FieldsNum[3], 2)
	}
}

func TestReadBytesA(t *testing.T) {
	str := "a;b;c\n1;2;3\n \n;;\n"
	header := []string{"a", "b", "c"}
	strBytes := []byte(str)
	csvData := ReadBytes(strBytes, header, ";")
	if len(csvData) != 3 {
		t.Error(len(csvData), 3)
	} else if len(csvData[0]) != 2 {
		t.Error(len(csvData[0]), 2)
	} else if len(csvData[1]) != 2 {
		t.Error(len(csvData[1]), 2)
	} else if len(csvData[2]) != 2 {
		t.Error(len(csvData[2]), 2)
	} else if csvData[0][1] != "1" {
		t.Error(csvData[0][1], "1")
	} else if csvData[1][1] != "2" {
		t.Error(csvData[1][1], "2")
	} else if csvData[2][1] != "3" {
		t.Error(csvData[2][1], "3")
	}
}

func TestReadBytesB(t *testing.T) {
	str := "c;a;b\n3;1;2\n \n;;\n"
	header := []string{"a", "b", "c"}
	strBytes := []byte(str)
	csvData := ReadBytes(strBytes, header, ";")
	if len(csvData) != 3 {
		t.Error(len(csvData), 3)
	} else if len(csvData[0]) != 2 {
		t.Error(len(csvData[0]), 2)
	} else if len(csvData[1]) != 2 {
		t.Error(len(csvData[1]), 2)
	} else if len(csvData[2]) != 2 {
		t.Error(len(csvData[2]), 2)
	} else if csvData[0][1] != "1" {
		t.Error(csvData[0][1], "1")
	} else if csvData[1][1] != "2" {
		t.Error(csvData[1][1], "2")
	} else if csvData[2][1] != "3" {
		t.Error(csvData[2][1], "3")
	}
}

func TestReadBytesC(t *testing.T) {
	str := "c;a;b;d;e\n3;1;2\n \n;;\n"
	header := []string{"a", "b", "c"}
	strBytes := []byte(str)
	csvData := ReadBytes(strBytes, header, ";")
	if len(csvData) != 3 {
		t.Error(len(csvData), 3)
	} else if len(csvData[0]) != 2 {
		t.Error(len(csvData[0]), 2)
	} else if len(csvData[1]) != 2 {
		t.Error(len(csvData[1]), 2)
	} else if len(csvData[2]) != 2 {
		t.Error(len(csvData[2]), 2)
	} else if csvData[0][1] != "1" {
		t.Error(csvData[0][1], "1")
	} else if csvData[1][1] != "2" {
		t.Error(csvData[1][1], "2")
	} else if csvData[2][1] != "3" {
		t.Error(csvData[2][1], "3")
	}
}

func TestReadBytesD(t *testing.T) {
	str := "c;a\n3;1\n \n;\n"
	header := []string{"a", "b", "c"}
	strBytes := []byte(str)
	csvData := ReadBytes(strBytes, header, ";")
	if len(csvData) != 3 {
		t.Error(len(csvData), 3)
	} else if len(csvData[0]) != 2 {
		t.Error(len(csvData[0]), 2)
	} else if len(csvData[1]) != 2 {
		t.Error(len(csvData[1]), 2)
	} else if len(csvData[2]) != 2 {
		t.Error(len(csvData[2]), 2)
	} else if csvData[0][1] != "1" {
		t.Error(csvData[0][1], "1")
	} else if csvData[1][1] != "" {
		t.Error(csvData[1][1])
	} else if csvData[2][1] != "3" {
		t.Error(csvData[2][1], "3")
	}
}

func TestReadBytesE(t *testing.T) {
	str := "1;2;3;4;5\n6;7;8;9\n \n;;\n"
	header := []string{"a", "b", "c"}
	strBytes := []byte(str)
	csvData := ReadBytes(strBytes, header, ";")
	if len(csvData) != 3 {
		t.Error(len(csvData), 3)
	} else if len(csvData[0]) != 3 {
		t.Error(len(csvData[0]), 3)
	} else if csvData[0][1] != "1" {
		t.Error(csvData[0][1], "1")
	} else if csvData[1][1] != "2" {
		t.Error(csvData[1][1], "2")
	} else if csvData[2][1] != "3" {
		t.Error(csvData[2][1], "3")
	} else if csvData[0][2] != "6" {
		t.Error(csvData[0][2], "6")
	} else if csvData[1][2] != "7" {
		t.Error(csvData[1][2], "7")
	} else if csvData[2][2] != "8" {
		t.Error(csvData[2][2], "8")
	}
}

func newLine(operatingSystem string) (string, string) {
	if operatingSystem == "windows" {
		return "\r\n", "\\r\\n"
	}
	return "\n", "\\n"
}
