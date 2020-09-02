/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package ensurecap provides the "ensure capacity" function for slices of basic types.
package ensurecap

import (
	"unsafe"
)

// NewCapFunc returns new capacity for a slice.
// New capacity must be greater or equal to reqCap.
type NewCapFunc func(oldLen, oldCap, reqCap int) int

// Bool ensures capacity is greater or equal to reqCap.
func Bool(list []bool, reqCap int, newCap NewCapFunc) []bool {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// BoolD2 ensures capacity is greater or equal to reqCap.
func BoolD2(list [][]bool, reqCap int, newCap NewCapFunc) [][]bool {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]bool, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// BoolD3 ensures capacity is greater or equal to reqCap.
func BoolD3(list [][][]bool, reqCap int, newCap NewCapFunc) [][][]bool {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]bool, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// BoolD4 ensures capacity is greater or equal to reqCap.
func BoolD4(list [][][][]bool, reqCap int, newCap NewCapFunc) [][][][]bool {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]bool, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// BoolD5 ensures capacity is greater or equal to reqCap.
func BoolD5(list [][][][][]bool, reqCap int, newCap NewCapFunc) [][][][][]bool {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]bool, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Byte ensures capacity is greater or equal to reqCap.
func Byte(list []byte, reqCap int, newCap NewCapFunc) []byte {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]byte, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// ByteD2 ensures capacity is greater or equal to reqCap.
func ByteD2(list [][]byte, reqCap int, newCap NewCapFunc) [][]byte {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]byte, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// ByteD3 ensures capacity is greater or equal to reqCap.
func ByteD3(list [][][]byte, reqCap int, newCap NewCapFunc) [][][]byte {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]byte, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// ByteD4 ensures capacity is greater or equal to reqCap.
func ByteD4(list [][][][]byte, reqCap int, newCap NewCapFunc) [][][][]byte {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]byte, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// ByteD5 ensures capacity is greater or equal to reqCap.
func ByteD5(list [][][][][]byte, reqCap int, newCap NewCapFunc) [][][][][]byte {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]byte, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex64 ensures capacity is greater or equal to reqCap.
func Complex64(list []complex64, reqCap int, newCap NewCapFunc) []complex64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]complex64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex64D2 ensures capacity is greater or equal to reqCap.
func Complex64D2(list [][]complex64, reqCap int, newCap NewCapFunc) [][]complex64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]complex64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex64D3 ensures capacity is greater or equal to reqCap.
func Complex64D3(list [][][]complex64, reqCap int, newCap NewCapFunc) [][][]complex64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]complex64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex64D4 ensures capacity is greater or equal to reqCap.
func Complex64D4(list [][][][]complex64, reqCap int, newCap NewCapFunc) [][][][]complex64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]complex64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex64D5 ensures capacity is greater or equal to reqCap.
func Complex64D5(list [][][][][]complex64, reqCap int, newCap NewCapFunc) [][][][][]complex64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]complex64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex128 ensures capacity is greater or equal to reqCap.
func Complex128(list []complex128, reqCap int, newCap NewCapFunc) []complex128 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]complex128, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex128D2 ensures capacity is greater or equal to reqCap.
func Complex128D2(list [][]complex128, reqCap int, newCap NewCapFunc) [][]complex128 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]complex128, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex128D3 ensures capacity is greater or equal to reqCap.
func Complex128D3(list [][][]complex128, reqCap int, newCap NewCapFunc) [][][]complex128 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]complex128, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex128D4 ensures capacity is greater or equal to reqCap.
func Complex128D4(list [][][][]complex128, reqCap int, newCap NewCapFunc) [][][][]complex128 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]complex128, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Complex128D5 ensures capacity is greater or equal to reqCap.
func Complex128D5(list [][][][][]complex128, reqCap int, newCap NewCapFunc) [][][][][]complex128 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]complex128, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Error ensures capacity is greater or equal to reqCap.
func Error(list []error, reqCap int, newCap NewCapFunc) []error {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]error, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// ErrorD2 ensures capacity is greater or equal to reqCap.
func ErrorD2(list [][]error, reqCap int, newCap NewCapFunc) [][]error {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]error, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// ErrorD3 ensures capacity is greater or equal to reqCap.
func ErrorD3(list [][][]error, reqCap int, newCap NewCapFunc) [][][]error {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]error, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// ErrorD4 ensures capacity is greater or equal to reqCap.
func ErrorD4(list [][][][]error, reqCap int, newCap NewCapFunc) [][][][]error {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]error, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// ErrorD5 ensures capacity is greater or equal to reqCap.
func ErrorD5(list [][][][][]error, reqCap int, newCap NewCapFunc) [][][][][]error {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]error, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float32 ensures capacity is greater or equal to reqCap.
func Float32(list []float32, reqCap int, newCap NewCapFunc) []float32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]float32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float32D2 ensures capacity is greater or equal to reqCap.
func Float32D2(list [][]float32, reqCap int, newCap NewCapFunc) [][]float32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]float32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float32D3 ensures capacity is greater or equal to reqCap.
func Float32D3(list [][][]float32, reqCap int, newCap NewCapFunc) [][][]float32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]float32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float32D4 ensures capacity is greater or equal to reqCap.
func Float32D4(list [][][][]float32, reqCap int, newCap NewCapFunc) [][][][]float32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]float32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float32D5 ensures capacity is greater or equal to reqCap.
func Float32D5(list [][][][][]float32, reqCap int, newCap NewCapFunc) [][][][][]float32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]float32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float64 ensures capacity is greater or equal to reqCap.
func Float64(list []float64, reqCap int, newCap NewCapFunc) []float64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]float64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float64D2 ensures capacity is greater or equal to reqCap.
func Float64D2(list [][]float64, reqCap int, newCap NewCapFunc) [][]float64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]float64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float64D3 ensures capacity is greater or equal to reqCap.
func Float64D3(list [][][]float64, reqCap int, newCap NewCapFunc) [][][]float64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]float64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float64D4 ensures capacity is greater or equal to reqCap.
func Float64D4(list [][][][]float64, reqCap int, newCap NewCapFunc) [][][][]float64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]float64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Float64D5 ensures capacity is greater or equal to reqCap.
func Float64D5(list [][][][][]float64, reqCap int, newCap NewCapFunc) [][][][][]float64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]float64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int ensures capacity is greater or equal to reqCap.
func Int(list []int, reqCap int, newCap NewCapFunc) []int {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]int, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// IntD2 ensures capacity is greater or equal to reqCap.
func IntD2(list [][]int, reqCap int, newCap NewCapFunc) [][]int {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]int, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// IntD3 ensures capacity is greater or equal to reqCap.
func IntD3(list [][][]int, reqCap int, newCap NewCapFunc) [][][]int {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]int, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// IntD4 ensures capacity is greater or equal to reqCap.
func IntD4(list [][][][]int, reqCap int, newCap NewCapFunc) [][][][]int {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]int, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// IntD5 ensures capacity is greater or equal to reqCap.
func IntD5(list [][][][][]int, reqCap int, newCap NewCapFunc) [][][][][]int {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]int, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int8 ensures capacity is greater or equal to reqCap.
func Int8(list []int8, reqCap int, newCap NewCapFunc) []int8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]int8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int8D2 ensures capacity is greater or equal to reqCap.
func Int8D2(list [][]int8, reqCap int, newCap NewCapFunc) [][]int8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]int8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int8D3 ensures capacity is greater or equal to reqCap.
func Int8D3(list [][][]int8, reqCap int, newCap NewCapFunc) [][][]int8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]int8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int8D4 ensures capacity is greater or equal to reqCap.
func Int8D4(list [][][][]int8, reqCap int, newCap NewCapFunc) [][][][]int8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]int8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int8D5 ensures capacity is greater or equal to reqCap.
func Int8D5(list [][][][][]int8, reqCap int, newCap NewCapFunc) [][][][][]int8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]int8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int16 ensures capacity is greater or equal to reqCap.
func Int16(list []int16, reqCap int, newCap NewCapFunc) []int16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]int16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int16D2 ensures capacity is greater or equal to reqCap.
func Int16D2(list [][]int16, reqCap int, newCap NewCapFunc) [][]int16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]int16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int16D3 ensures capacity is greater or equal to reqCap.
func Int16D3(list [][][]int16, reqCap int, newCap NewCapFunc) [][][]int16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]int16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int16D4 ensures capacity is greater or equal to reqCap.
func Int16D4(list [][][][]int16, reqCap int, newCap NewCapFunc) [][][][]int16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]int16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int16D5 ensures capacity is greater or equal to reqCap.
func Int16D5(list [][][][][]int16, reqCap int, newCap NewCapFunc) [][][][][]int16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]int16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int32 ensures capacity is greater or equal to reqCap.
func Int32(list []int32, reqCap int, newCap NewCapFunc) []int32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]int32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int32D2 ensures capacity is greater or equal to reqCap.
func Int32D2(list [][]int32, reqCap int, newCap NewCapFunc) [][]int32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]int32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int32D3 ensures capacity is greater or equal to reqCap.
func Int32D3(list [][][]int32, reqCap int, newCap NewCapFunc) [][][]int32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]int32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int32D4 ensures capacity is greater or equal to reqCap.
func Int32D4(list [][][][]int32, reqCap int, newCap NewCapFunc) [][][][]int32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]int32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int32D5 ensures capacity is greater or equal to reqCap.
func Int32D5(list [][][][][]int32, reqCap int, newCap NewCapFunc) [][][][][]int32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]int32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int64 ensures capacity is greater or equal to reqCap.
func Int64(list []int64, reqCap int, newCap NewCapFunc) []int64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]int64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int64D2 ensures capacity is greater or equal to reqCap.
func Int64D2(list [][]int64, reqCap int, newCap NewCapFunc) [][]int64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]int64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int64D3 ensures capacity is greater or equal to reqCap.
func Int64D3(list [][][]int64, reqCap int, newCap NewCapFunc) [][][]int64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]int64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int64D4 ensures capacity is greater or equal to reqCap.
func Int64D4(list [][][][]int64, reqCap int, newCap NewCapFunc) [][][][]int64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]int64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Int64D5 ensures capacity is greater or equal to reqCap.
func Int64D5(list [][][][][]int64, reqCap int, newCap NewCapFunc) [][][][][]int64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]int64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Interface ensures capacity is greater or equal to reqCap.
func Interface(list []interface{}, reqCap int, newCap NewCapFunc) []interface{} {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]interface{}, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// InterfaceD2 ensures capacity is greater or equal to reqCap.
func InterfaceD2(list [][]interface{}, reqCap int, newCap NewCapFunc) [][]interface{} {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]interface{}, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// InterfaceD3 ensures capacity is greater or equal to reqCap.
func InterfaceD3(list [][][]interface{}, reqCap int, newCap NewCapFunc) [][][]interface{} {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]interface{}, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// InterfaceD4 ensures capacity is greater or equal to reqCap.
func InterfaceD4(list [][][][]interface{}, reqCap int, newCap NewCapFunc) [][][][]interface{} {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]interface{}, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// InterfaceD5 ensures capacity is greater or equal to reqCap.
func InterfaceD5(list [][][][][]interface{}, reqCap int, newCap NewCapFunc) [][][][][]interface{} {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]interface{}, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Pointer ensures capacity is greater or equal to reqCap.
func Pointer(list []unsafe.Pointer, reqCap int, newCap NewCapFunc) []unsafe.Pointer {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]unsafe.Pointer, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// PointerD2 ensures capacity is greater or equal to reqCap.
func PointerD2(list [][]unsafe.Pointer, reqCap int, newCap NewCapFunc) [][]unsafe.Pointer {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]unsafe.Pointer, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// PointerD3 ensures capacity is greater or equal to reqCap.
func PointerD3(list [][][]unsafe.Pointer, reqCap int, newCap NewCapFunc) [][][]unsafe.Pointer {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]unsafe.Pointer, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// PointerD4 ensures capacity is greater or equal to reqCap.
func PointerD4(list [][][][]unsafe.Pointer, reqCap int, newCap NewCapFunc) [][][][]unsafe.Pointer {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]unsafe.Pointer, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// PointerD5 ensures capacity is greater or equal to reqCap.
func PointerD5(list [][][][][]unsafe.Pointer, reqCap int, newCap NewCapFunc) [][][][][]unsafe.Pointer {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]unsafe.Pointer, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// Rune ensures capacity is greater or equal to reqCap.
func Rune(list []rune, reqCap int, newCap NewCapFunc) []rune {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]rune, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// RuneD2 ensures capacity is greater or equal to reqCap.
func RuneD2(list [][]rune, reqCap int, newCap NewCapFunc) [][]rune {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]rune, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// RuneD3 ensures capacity is greater or equal to reqCap.
func RuneD3(list [][][]rune, reqCap int, newCap NewCapFunc) [][][]rune {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]rune, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// RuneD4 ensures capacity is greater or equal to reqCap.
func RuneD4(list [][][][]rune, reqCap int, newCap NewCapFunc) [][][][]rune {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]rune, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// RuneD5 ensures capacity is greater or equal to reqCap.
func RuneD5(list [][][][][]rune, reqCap int, newCap NewCapFunc) [][][][][]rune {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]rune, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// String ensures capacity is greater or equal to reqCap.
func String(list []string, reqCap int, newCap NewCapFunc) []string {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]string, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// StringD2 ensures capacity is greater or equal to reqCap.
func StringD2(list [][]string, reqCap int, newCap NewCapFunc) [][]string {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]string, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// StringD3 ensures capacity is greater or equal to reqCap.
func StringD3(list [][][]string, reqCap int, newCap NewCapFunc) [][][]string {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]string, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// StringD4 ensures capacity is greater or equal to reqCap.
func StringD4(list [][][][]string, reqCap int, newCap NewCapFunc) [][][][]string {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]string, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// StringD5 ensures capacity is greater or equal to reqCap.
func StringD5(list [][][][][]string, reqCap int, newCap NewCapFunc) [][][][][]string {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]string, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt ensures capacity is greater or equal to reqCap.
func UInt(list []uint, reqCap int, newCap NewCapFunc) []uint {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]uint, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UIntD2 ensures capacity is greater or equal to reqCap.
func UIntD2(list [][]uint, reqCap int, newCap NewCapFunc) [][]uint {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]uint, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UIntD3 ensures capacity is greater or equal to reqCap.
func UIntD3(list [][][]uint, reqCap int, newCap NewCapFunc) [][][]uint {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]uint, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UIntD4 ensures capacity is greater or equal to reqCap.
func UIntD4(list [][][][]uint, reqCap int, newCap NewCapFunc) [][][][]uint {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]uint, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UIntD5 ensures capacity is greater or equal to reqCap.
func UIntD5(list [][][][][]uint, reqCap int, newCap NewCapFunc) [][][][][]uint {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]uint, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt8 ensures capacity is greater or equal to reqCap.
func UInt8(list []uint8, reqCap int, newCap NewCapFunc) []uint8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]uint8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt8D2 ensures capacity is greater or equal to reqCap.
func UInt8D2(list [][]uint8, reqCap int, newCap NewCapFunc) [][]uint8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]uint8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt8D3 ensures capacity is greater or equal to reqCap.
func UInt8D3(list [][][]uint8, reqCap int, newCap NewCapFunc) [][][]uint8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]uint8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt8D4 ensures capacity is greater or equal to reqCap.
func UInt8D4(list [][][][]uint8, reqCap int, newCap NewCapFunc) [][][][]uint8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]uint8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt8D5 ensures capacity is greater or equal to reqCap.
func UInt8D5(list [][][][][]uint8, reqCap int, newCap NewCapFunc) [][][][][]uint8 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]uint8, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt16 ensures capacity is greater or equal to reqCap.
func UInt16(list []uint16, reqCap int, newCap NewCapFunc) []uint16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]uint16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt16D2 ensures capacity is greater or equal to reqCap.
func UInt16D2(list [][]uint16, reqCap int, newCap NewCapFunc) [][]uint16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]uint16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt16D3 ensures capacity is greater or equal to reqCap.
func UInt16D3(list [][][]uint16, reqCap int, newCap NewCapFunc) [][][]uint16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]uint16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt16D4 ensures capacity is greater or equal to reqCap.
func UInt16D4(list [][][][]uint16, reqCap int, newCap NewCapFunc) [][][][]uint16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]uint16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt16D5 ensures capacity is greater or equal to reqCap.
func UInt16D5(list [][][][][]uint16, reqCap int, newCap NewCapFunc) [][][][][]uint16 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]uint16, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt32 ensures capacity is greater or equal to reqCap.
func UInt32(list []uint32, reqCap int, newCap NewCapFunc) []uint32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]uint32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt32D2 ensures capacity is greater or equal to reqCap.
func UInt32D2(list [][]uint32, reqCap int, newCap NewCapFunc) [][]uint32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]uint32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt32D3 ensures capacity is greater or equal to reqCap.
func UInt32D3(list [][][]uint32, reqCap int, newCap NewCapFunc) [][][]uint32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]uint32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt32D4 ensures capacity is greater or equal to reqCap.
func UInt32D4(list [][][][]uint32, reqCap int, newCap NewCapFunc) [][][][]uint32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]uint32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt32D5 ensures capacity is greater or equal to reqCap.
func UInt32D5(list [][][][][]uint32, reqCap int, newCap NewCapFunc) [][][][][]uint32 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]uint32, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt64 ensures capacity is greater or equal to reqCap.
func UInt64(list []uint64, reqCap int, newCap NewCapFunc) []uint64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([]uint64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt64D2 ensures capacity is greater or equal to reqCap.
func UInt64D2(list [][]uint64, reqCap int, newCap NewCapFunc) [][]uint64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][]uint64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt64D3 ensures capacity is greater or equal to reqCap.
func UInt64D3(list [][][]uint64, reqCap int, newCap NewCapFunc) [][][]uint64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][]uint64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt64D4 ensures capacity is greater or equal to reqCap.
func UInt64D4(list [][][][]uint64, reqCap int, newCap NewCapFunc) [][][][]uint64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][]uint64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}

// UInt64D5 ensures capacity is greater or equal to reqCap.
func UInt64D5(list [][][][][]uint64, reqCap int, newCap NewCapFunc) [][][][][]uint64 {
	if reqCap <= cap(list) {
		return list
	}
	newList := make([][][][][]uint64, len(list), newCap(len(list), cap(list), reqCap))
	if len(list) > 0 {
		copy(newList, list)
	}
	return newList
}
