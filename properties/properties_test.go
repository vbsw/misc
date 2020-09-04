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
	props := map[string]string{"al ice": "a"}
	nlLength := newLineLength(runtime.GOOS)

	bytes := ToBytes(props)
	if len(bytes) != len("al\\ ice=a")+nlLength {
		t.Error(len(bytes), len("al\\ ice=a")+nlLength)
	} else if string(bytes[:len(bytes)-nlLength]) != "al\\ ice=a" {
		t.Error(string(bytes[:len(bytes)-nlLength]), "al\\ ice=a")
	}

	bytes = ToBytes(props, Spaces, OpCollon)
	if len(bytes) != len("al\\ ice : a")+nlLength {
		t.Error(len(bytes), len("al\\ ice : a")+nlLength)
	} else if string(bytes[:len(bytes)-nlLength]) != "al\\ ice : a" {
		t.Error(string(bytes[:len(bytes)-nlLength]), "al\\ ice : a")
	}

	bytes = ToBytes(props, OpSpace)
	if len(bytes) != len("al\\ ice a")+nlLength {
		t.Error(len(bytes), len("al\\ ice a")+nlLength)
	} else if string(bytes[:len(bytes)-nlLength]) != "al\\ ice a" {
		t.Error(string(bytes[:len(bytes)-nlLength]), "al\\ ice a")
	}
}

func TestReadBytes(t *testing.T) {
	props := map[string]string{"alice": "a", "bob": "", "clair": "c", "david": "d"}
	bytes := ToBytes(props)
	propsResult := ReadBytes(bytes)

	if len(props) == len(propsResult) {
		for k, v := range props {
			checkProps(propsResult, k, v, t)
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
	props := map[string]string{"alice": "a", "bob": "b", "clair": "c", "david": "d"}
	WriteFile(path, props)
	propsResult, _ := ReadFile(path)

	if len(props) == len(propsResult) {
		for k, v := range props {
			checkProps(propsResult, k, v, t)
		}
	} else {
		t.Error(len(propsResult))
	}
}

func checkProps(props map[string]string, k, v string, t *testing.T) {
	propsVal, exists := props[k]
	if exists {
		if propsVal != v {
			t.Error(propsVal)
		}
	} else {
		t.Error(k, "does not exist")
	}
}
