/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package properties

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestToBytes(t *testing.T) {
	propNames := []string{"al ice"}
	propValues := []string{"a"}
	nlLength := newLineLength(runtime.GOOS)

	bytes := ToBytes(propNames, propValues)
	if len(bytes) != len("al\\ ice=a")+nlLength {
		t.Error(len(bytes), len("al\\ ice=a")+nlLength)
	} else if string(bytes[:len(bytes)-nlLength]) != "al\\ ice=a" {
		t.Error(string(bytes[:len(bytes)-nlLength]), "al\\ ice=a")
	}

	bytes = ToBytes(propNames, propValues, Spaces, OpCollon)
	if len(bytes) != len("al\\ ice : a")+nlLength {
		t.Error(len(bytes), len("al\\ ice : a")+nlLength)
	} else if string(bytes[:len(bytes)-nlLength]) != "al\\ ice : a" {
		t.Error(string(bytes[:len(bytes)-nlLength]), "al\\ ice : a")
	}

	bytes = ToBytes(propNames, propValues, OpSpace)
	if len(bytes) != len("al\\ ice a")+nlLength {
		t.Error(len(bytes), len("al\\ ice a")+nlLength)
	} else if string(bytes[:len(bytes)-nlLength]) != "al\\ ice a" {
		t.Error(string(bytes[:len(bytes)-nlLength]), "al\\ ice a")
	}
}

func TestReadBytes(t *testing.T) {
	propNames := []string{"alice", "bob", "clair", "david"}
	propValues := []string{"a", "", "c", "d"}
	bytes := ToBytes(propNames, propValues)
	propsResult := ReadBytes(bytes)

	if len(propNames) == len(propsResult) {
		for i, propName := range propNames {
			propValue := propValues[i]
			checkProps(propsResult, propName, propValue, t)
		}
	} else {
		t.Error(len(propsResult))
	}
	bytes = []byte("alice=a\nbob\nclair")
	propsResult = ReadBytes(bytes)

	if len(propsResult) == 3 {
		checkProps(propsResult, "alice", "a", t)
		checkProps(propsResult, "bob", "", t)
		checkProps(propsResult, "clair", "", t)
	} else {
		t.Error(len(propsResult))
	}
}

func TestWriteRead(t *testing.T) {
	path := filepath.Join(os.TempDir(), "test.properties")
	propNames := []string{"alice", "bob", "clair", "david"}
	propValues := []string{"a", "b", "c", "d"}
	WriteFile(path, propNames, propValues)
	propsResult, _ := ReadFile(path)

	if len(propNames) == len(propsResult) {
		for i, propName := range propNames {
			propValue := propValues[i]
			checkProps(propsResult, propName, propValue, t)
		}
	} else {
		t.Error(len(propsResult))
	}
}

func checkProps(props map[string]string, propName, propValue string, t *testing.T) {
	propsVal, exists := props[propName]
	if exists {
		if propsVal != propValue {
			t.Error(propsVal)
		}
	} else {
		t.Error(propName, "does not exist")
	}
}
