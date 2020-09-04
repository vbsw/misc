/*
 *       Copyright 2018, 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package indexl provides "indexl" function for slices of basic types. "indexl" searches for a value in array from left to right and returns its index.
package indexl

import (
	"unsafe"
)

// Bool searches for value in list from left to right and returns its index. If not found, returns -1.
func Bool(list []bool, value bool) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// BoolD2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func BoolD2(list [][]bool, value bool) (int, int) {
	for i, listValue := range list {
		j := Bool(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// BoolD3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func BoolD3(list [][][]bool, value bool) (int, int, int) {
	for i, listValue := range list {
		j, k := BoolD2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// BoolD4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func BoolD4(list [][][][]bool, value bool) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := BoolD3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// BoolD5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func BoolD5(list [][][][][]bool, value bool) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := BoolD4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Byte searches for value in list from left to right and returns its index. If not found, returns -1.
func Byte(list []byte, value byte) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// ByteD2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func ByteD2(list [][]byte, value byte) (int, int) {
	for i, listValue := range list {
		j := Byte(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// ByteD3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func ByteD3(list [][][]byte, value byte) (int, int, int) {
	for i, listValue := range list {
		j, k := ByteD2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// ByteD4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func ByteD4(list [][][][]byte, value byte) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := ByteD3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// ByteD5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func ByteD5(list [][][][][]byte, value byte) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := ByteD4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Complex64 searches for value in list from left to right and returns its index. If not found, returns -1.
func Complex64(list []complex64, value complex64) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// Complex64D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func Complex64D2(list [][]complex64, value complex64) (int, int) {
	for i, listValue := range list {
		j := Complex64(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// Complex64D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func Complex64D3(list [][][]complex64, value complex64) (int, int, int) {
	for i, listValue := range list {
		j, k := Complex64D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// Complex64D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func Complex64D4(list [][][][]complex64, value complex64) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := Complex64D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// Complex64D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func Complex64D5(list [][][][][]complex64, value complex64) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := Complex64D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Complex128 searches for value in list from left to right and returns its index. If not found, returns -1.
func Complex128(list []complex128, value complex128) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// Complex128D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func Complex128D2(list [][]complex128, value complex128) (int, int) {
	for i, listValue := range list {
		j := Complex128(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// Complex128D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func Complex128D3(list [][][]complex128, value complex128) (int, int, int) {
	for i, listValue := range list {
		j, k := Complex128D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// Complex128D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func Complex128D4(list [][][][]complex128, value complex128) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := Complex128D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// Complex128D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func Complex128D5(list [][][][][]complex128, value complex128) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := Complex128D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Error searches for value in list from left to right and returns its index. If not found, returns -1.
func Error(list []error, value error) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// ErrorD2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func ErrorD2(list [][]error, value error) (int, int) {
	for i, listValue := range list {
		j := Error(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// ErrorD3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func ErrorD3(list [][][]error, value error) (int, int, int) {
	for i, listValue := range list {
		j, k := ErrorD2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// ErrorD4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func ErrorD4(list [][][][]error, value error) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := ErrorD3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// ErrorD5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func ErrorD5(list [][][][][]error, value error) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := ErrorD4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Float32 searches for value in list from left to right and returns its index. If not found, returns -1.
func Float32(list []float32, value float32) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// Float32D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func Float32D2(list [][]float32, value float32) (int, int) {
	for i, listValue := range list {
		j := Float32(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// Float32D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func Float32D3(list [][][]float32, value float32) (int, int, int) {
	for i, listValue := range list {
		j, k := Float32D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// Float32D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func Float32D4(list [][][][]float32, value float32) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := Float32D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// Float32D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func Float32D5(list [][][][][]float32, value float32) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := Float32D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Float64 searches for value in list from left to right and returns its index. If not found, returns -1.
func Float64(list []float64, value float64) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// Float64D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func Float64D2(list [][]float64, value float64) (int, int) {
	for i, listValue := range list {
		j := Float64(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// Float64D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func Float64D3(list [][][]float64, value float64) (int, int, int) {
	for i, listValue := range list {
		j, k := Float64D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// Float64D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func Float64D4(list [][][][]float64, value float64) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := Float64D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// Float64D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func Float64D5(list [][][][][]float64, value float64) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := Float64D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Int searches for value in list from left to right and returns its index. If not found, returns -1.
func Int(list []int, value int) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// IntD2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func IntD2(list [][]int, value int) (int, int) {
	for i, listValue := range list {
		j := Int(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// IntD3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func IntD3(list [][][]int, value int) (int, int, int) {
	for i, listValue := range list {
		j, k := IntD2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// IntD4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func IntD4(list [][][][]int, value int) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := IntD3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// IntD5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func IntD5(list [][][][][]int, value int) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := IntD4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Int8 searches for value in list from left to right and returns its index. If not found, returns -1.
func Int8(list []int8, value int8) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// Int8D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func Int8D2(list [][]int8, value int8) (int, int) {
	for i, listValue := range list {
		j := Int8(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// Int8D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func Int8D3(list [][][]int8, value int8) (int, int, int) {
	for i, listValue := range list {
		j, k := Int8D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// Int8D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func Int8D4(list [][][][]int8, value int8) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := Int8D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// Int8D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func Int8D5(list [][][][][]int8, value int8) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := Int8D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Int16 searches for value in list from left to right and returns its index. If not found, returns -1.
func Int16(list []int16, value int16) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// Int16D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func Int16D2(list [][]int16, value int16) (int, int) {
	for i, listValue := range list {
		j := Int16(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// Int16D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func Int16D3(list [][][]int16, value int16) (int, int, int) {
	for i, listValue := range list {
		j, k := Int16D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// Int16D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func Int16D4(list [][][][]int16, value int16) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := Int16D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// Int16D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func Int16D5(list [][][][][]int16, value int16) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := Int16D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Int32 searches for value in list from left to right and returns its index. If not found, returns -1.
func Int32(list []int32, value int32) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// Int32D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func Int32D2(list [][]int32, value int32) (int, int) {
	for i, listValue := range list {
		j := Int32(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// Int32D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func Int32D3(list [][][]int32, value int32) (int, int, int) {
	for i, listValue := range list {
		j, k := Int32D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// Int32D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func Int32D4(list [][][][]int32, value int32) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := Int32D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// Int32D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func Int32D5(list [][][][][]int32, value int32) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := Int32D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Int64 searches for value in list from left to right and returns its index. If not found, returns -1.
func Int64(list []int64, value int64) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// Int64D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func Int64D2(list [][]int64, value int64) (int, int) {
	for i, listValue := range list {
		j := Int64(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// Int64D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func Int64D3(list [][][]int64, value int64) (int, int, int) {
	for i, listValue := range list {
		j, k := Int64D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// Int64D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func Int64D4(list [][][][]int64, value int64) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := Int64D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// Int64D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func Int64D5(list [][][][][]int64, value int64) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := Int64D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Interface searches for value in list from left to right and returns its index. If not found, returns -1.
func Interface(list []interface{}, value interface{}) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// InterfaceD2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func InterfaceD2(list [][]interface{}, value interface{}) (int, int) {
	for i, listValue := range list {
		j := Interface(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// InterfaceD3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func InterfaceD3(list [][][]interface{}, value interface{}) (int, int, int) {
	for i, listValue := range list {
		j, k := InterfaceD2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// InterfaceD4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func InterfaceD4(list [][][][]interface{}, value interface{}) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := InterfaceD3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// InterfaceD5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func InterfaceD5(list [][][][][]interface{}, value interface{}) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := InterfaceD4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Pointer searches for value in list from left to right and returns its index. If not found, returns -1.
func Pointer(list []unsafe.Pointer, value unsafe.Pointer) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// PointerD2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func PointerD2(list [][]unsafe.Pointer, value unsafe.Pointer) (int, int) {
	for i, listValue := range list {
		j := Pointer(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// PointerD3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func PointerD3(list [][][]unsafe.Pointer, value unsafe.Pointer) (int, int, int) {
	for i, listValue := range list {
		j, k := PointerD2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// PointerD4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func PointerD4(list [][][][]unsafe.Pointer, value unsafe.Pointer) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := PointerD3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// PointerD5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func PointerD5(list [][][][][]unsafe.Pointer, value unsafe.Pointer) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := PointerD4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// Rune searches for value in list from left to right and returns its index. If not found, returns -1.
func Rune(list []rune, value rune) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// RuneD2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func RuneD2(list [][]rune, value rune) (int, int) {
	for i, listValue := range list {
		j := Rune(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// RuneD3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func RuneD3(list [][][]rune, value rune) (int, int, int) {
	for i, listValue := range list {
		j, k := RuneD2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// RuneD4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func RuneD4(list [][][][]rune, value rune) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := RuneD3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// RuneD5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func RuneD5(list [][][][][]rune, value rune) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := RuneD4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// String searches for value in list from left to right and returns its index. If not found, returns -1.
func String(list []string, value string) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// StringD2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func StringD2(list [][]string, value string) (int, int) {
	for i, listValue := range list {
		j := String(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// StringD3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func StringD3(list [][][]string, value string) (int, int, int) {
	for i, listValue := range list {
		j, k := StringD2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// StringD4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func StringD4(list [][][][]string, value string) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := StringD3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// StringD5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func StringD5(list [][][][][]string, value string) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := StringD4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// UInt searches for value in list from left to right and returns its index. If not found, returns -1.
func UInt(list []uint, value uint) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// UIntD2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func UIntD2(list [][]uint, value uint) (int, int) {
	for i, listValue := range list {
		j := UInt(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// UIntD3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func UIntD3(list [][][]uint, value uint) (int, int, int) {
	for i, listValue := range list {
		j, k := UIntD2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// UIntD4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func UIntD4(list [][][][]uint, value uint) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := UIntD3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// UIntD5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func UIntD5(list [][][][][]uint, value uint) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := UIntD4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// UInt8 searches for value in list from left to right and returns its index. If not found, returns -1.
func UInt8(list []uint8, value uint8) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// UInt8D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func UInt8D2(list [][]uint8, value uint8) (int, int) {
	for i, listValue := range list {
		j := UInt8(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// UInt8D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func UInt8D3(list [][][]uint8, value uint8) (int, int, int) {
	for i, listValue := range list {
		j, k := UInt8D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// UInt8D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func UInt8D4(list [][][][]uint8, value uint8) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := UInt8D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// UInt8D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func UInt8D5(list [][][][][]uint8, value uint8) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := UInt8D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// UInt16 searches for value in list from left to right and returns its index. If not found, returns -1.
func UInt16(list []uint16, value uint16) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// UInt16D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func UInt16D2(list [][]uint16, value uint16) (int, int) {
	for i, listValue := range list {
		j := UInt16(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// UInt16D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func UInt16D3(list [][][]uint16, value uint16) (int, int, int) {
	for i, listValue := range list {
		j, k := UInt16D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// UInt16D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func UInt16D4(list [][][][]uint16, value uint16) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := UInt16D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// UInt16D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func UInt16D5(list [][][][][]uint16, value uint16) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := UInt16D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// UInt32 searches for value in list from left to right and returns its index. If not found, returns -1.
func UInt32(list []uint32, value uint32) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// UInt32D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func UInt32D2(list [][]uint32, value uint32) (int, int) {
	for i, listValue := range list {
		j := UInt32(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// UInt32D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func UInt32D3(list [][][]uint32, value uint32) (int, int, int) {
	for i, listValue := range list {
		j, k := UInt32D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// UInt32D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func UInt32D4(list [][][][]uint32, value uint32) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := UInt32D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// UInt32D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func UInt32D5(list [][][][][]uint32, value uint32) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := UInt32D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}

// UInt64 searches for value in list from left to right and returns its index. If not found, returns -1.
func UInt64(list []uint64, value uint64) int {
	for i, listValue := range list {
		if listValue == value {
			return i
		}
	}
	return -1
}

// UInt64D2 searches for value in list from left to right and returns its index. If not found, returns -1, -1.
func UInt64D2(list [][]uint64, value uint64) (int, int) {
	for i, listValue := range list {
		j := UInt64(listValue, value)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}

// UInt64D3 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1.
func UInt64D3(list [][][]uint64, value uint64) (int, int, int) {
	for i, listValue := range list {
		j, k := UInt64D2(listValue, value)
		if j >= 0 && k >= 0 {
			return i, j, k
		}
	}
	return -1, -1, -1
}

// UInt64D4 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1.
func UInt64D4(list [][][][]uint64, value uint64) (int, int, int, int) {
	for i, listValue := range list {
		j, k, l := UInt64D3(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 {
			return i, j, k, l
		}
	}
	return -1, -1, -1, -1
}

// UInt64D5 searches for value in list from left to right and returns its index. If not found, returns -1, -1, -1, -1, -1.
func UInt64D5(list [][][][][]uint64, value uint64) (int, int, int, int, int) {
	for i, listValue := range list {
		j, k, l, m := UInt64D4(listValue, value)
		if j >= 0 && k >= 0 && l >= 0 && m >= 0 {
			return i, j, k, l, m
		}
	}
	return -1, -1, -1, -1, -1
}
