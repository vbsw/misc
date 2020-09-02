/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package ref provides unsafe functions regarding data referencing.
package ref

import (
	"reflect"
	"unsafe"
)

// Bytes retuns the string as byte slice, but without copying the bytes.
func Bytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// String retuns the byte slice as string, but without copying the bytes.
func String(bytes []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	strHeader := reflect.StringHeader{bytesHeader.Data, bytesHeader.Len}
	return *(*string)(unsafe.Pointer(&strHeader))
}
