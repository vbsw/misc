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
	"unsafe"
)

func TestReadBytes(t *testing.T) {
	props := map[string]string{"alice": "a", "bob": "", "clair": "c", "david": "d"}
	bytes := propsToBytes(props)
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

func propsToBytes(props map[string]string) []byte {
	length := propsToBytesLength(props)
	bytes := make([]byte, length)
	if runtime.GOOS == "windows" {
		copyPropsToBytesWindows(bytes, props)
	} else {
		copyPropsToBytes(bytes, props)
	}
	return bytes
}

func propsToBytesLength(props map[string]string) int {
	var length int
	for k, v := range props {
		if len(k) > 0 {
			length += len(k) + len(v)
		}
	}
	// plus assignment operator and line endings
	if runtime.GOOS == "windows" {
		return length + 3*len(props)
	}
	return length + 2*len(props)
}

func copyPropsToBytesWindows(bytes []byte, props map[string]string) {
	for k, v := range props {
		if len(k) > 0 {
			kBytes := *(*[]byte)(unsafe.Pointer(&k))
			vBytes := *(*[]byte)(unsafe.Pointer(&v))
			lineLength := len(kBytes) + len(vBytes) + 1
			copy(bytes, kBytes)
			bytes[len(kBytes)] = '='
			copy(bytes[len(kBytes)+1:], vBytes)
			bytes[lineLength] = '\r'
			bytes[lineLength+1] = '\n'
			bytes = bytes[lineLength+2:]
		}
	}
}

func copyPropsToBytes(bytes []byte, props map[string]string) {
	for k, v := range props {
		if len(k) > 0 {
			kBytes := *(*[]byte)(unsafe.Pointer(&k))
			vBytes := *(*[]byte)(unsafe.Pointer(&v))
			lineLength := len(kBytes) + len(vBytes) + 1
			copy(bytes, kBytes)
			bytes[len(kBytes)] = '='
			copy(bytes[len(kBytes)+1:], vBytes)
			bytes[lineLength] = '\n'
			bytes = bytes[lineLength+1:]
		}
	}
}
