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

// NewCapFunc defines a function that returns the new capacity for a slice.
// New capacity must be greater or equal to oldCap + n.
type NewCapFunc func(oldLen, oldCap, n int) int

// Bool ensures capacity of list is greater or equal to len(list) + n.
func Bool(list []bool, n int, newCap NewCapFunc) []bool {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// BoolD2 ensures capacity of list is greater or equal to len(list) + n.
func BoolD2(list [][]bool, index int) [][]bool {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// BoolD3 ensures capacity of list is greater or equal to len(list) + n.
func BoolD3(list [][][]bool, index int) [][][]bool {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// BoolD4 ensures capacity of list is greater or equal to len(list) + n.
func BoolD4(list [][][][]bool, index int) [][][][]bool {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// BoolD5 ensures capacity of list is greater or equal to len(list) + n.
func BoolD5(list [][][][][]bool, index int) [][][][][]bool {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][][]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Byte ensures capacity of list is greater or equal to len(list) + n.
func Byte(list []byte, index int) []byte {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]byte, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// ByteD2 ensures capacity of list is greater or equal to len(list) + n.
func ByteD2(list [][]byte, index int) [][]byte {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][]byte, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// ByteD3 ensures capacity of list is greater or equal to len(list) + n.
func ByteD3(list [][][]byte, index int) [][][]byte {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][]byte, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// ByteD4 ensures capacity of list is greater or equal to len(list) + n.
func ByteD4(list [][][][]byte, index int) [][][][]byte {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][]byte, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// ByteD5 ensures capacity of list is greater or equal to len(list) + n.
func ByteD5(list [][][][][]byte, index int) [][][][][]byte {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][][]byte, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex64 ensures capacity of list is greater or equal to len(list) + n.
func Complex64(list []complex64, index int) []complex64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]complex64, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex64D2 ensures capacity of list is greater or equal to len(list) + n.
func Complex64D2(list [][]complex64, index int) [][]complex64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][]complex64, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex64D3 ensures capacity of list is greater or equal to len(list) + n.
func Complex64D3(list [][][]complex64, index int) [][][]complex64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][]complex64, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex64D4 ensures capacity of list is greater or equal to len(list) + n.
func Complex64D4(list [][][][]complex64, index int) [][][][]complex64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][]complex64, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex64D5 ensures capacity of list is greater or equal to len(list) + n.
func Complex64D5(list [][][][][]complex64, index int) [][][][][]complex64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][][]complex64, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex128 ensures capacity of list is greater or equal to len(list) + n.
func Complex128(list []complex128, index int) []complex128 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]complex128, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex128D2 ensures capacity of list is greater or equal to len(list) + n.
func Complex128D2(list [][]complex128, index int) [][]complex128 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][]complex128, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex128D3 ensures capacity of list is greater or equal to len(list) + n.
func Complex128D3(list [][][]complex128, index int) [][][]complex128 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][]complex128, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex128D4 ensures capacity of list is greater or equal to len(list) + n.
func Complex128D4(list [][][][]complex128, index int) [][][][]complex128 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][]complex128, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Complex128D5 ensures capacity of list is greater or equal to len(list) + n.
func Complex128D5(list [][][][][]complex128, index int) [][][][][]complex128 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][][]complex128, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Error ensures capacity of list is greater or equal to len(list) + n.
func Error(list []error, index int) []error {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]error, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// ErrorD2 ensures capacity of list is greater or equal to len(list) + n.
func ErrorD2(list [][]error, index int) [][]error {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][]error, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// ErrorD3 ensures capacity of list is greater or equal to len(list) + n.
func ErrorD3(list [][][]error, index int) [][][]error {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][]error, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// ErrorD4 ensures capacity of list is greater or equal to len(list) + n.
func ErrorD4(list [][][][]error, index int) [][][][]error {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][]error, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// ErrorD5 ensures capacity of list is greater or equal to len(list) + n.
func ErrorD5(list [][][][][]error, index int) [][][][][]error {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][][][]error, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float32 ensures capacity of list is greater or equal to len(list) + n.
func Float32(list []float32, index int) []float32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]float32, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float32D2 ensures capacity of list is greater or equal to len(list) + n.
func Float32D2(list [][]float32, index int) [][]float32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][]float32, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float32D3 ensures capacity of list is greater or equal to len(list) + n.
func Float32D3(list [][][]float32, index int) [][][]float32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([][][]float32, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float32D4 ensures capacity of list is greater or equal to len(list) + n.
func Float32D4(list [][][][]float32, index int) [][][][]float32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float32D5 ensures capacity of list is greater or equal to len(list) + n.
func Float32D5(list [][][][][]float32, index int) [][][][][]float32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float64 ensures capacity of list is greater or equal to len(list) + n.
func Float64(list []float64, index int) []float64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float64D2 ensures capacity of list is greater or equal to len(list) + n.
func Float64D2(list [][]float64, index int) [][]float64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float64D3 ensures capacity of list is greater or equal to len(list) + n.
func Float64D3(list [][][]float64, index int) [][][]float64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float64D4 ensures capacity of list is greater or equal to len(list) + n.
func Float64D4(list [][][][]float64, index int) [][][][]float64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Float64D5 ensures capacity of list is greater or equal to len(list) + n.
func Float64D5(list [][][][][]float64, index int) [][][][][]float64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int ensures capacity of list is greater or equal to len(list) + n.
func Int(list []int, index int) []int {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// IntD2 ensures capacity of list is greater or equal to len(list) + n.
func IntD2(list [][]int, index int) [][]int {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// IntD3 ensures capacity of list is greater or equal to len(list) + n.
func IntD3(list [][][]int, index int) [][][]int {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// IntD4 ensures capacity of list is greater or equal to len(list) + n.
func IntD4(list [][][][]int, index int) [][][][]int {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// IntD5 ensures capacity of list is greater or equal to len(list) + n.
func IntD5(list [][][][][]int, index int) [][][][][]int {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int8 ensures capacity of list is greater or equal to len(list) + n.
func Int8(list []int8, index int) []int8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int8D2 ensures capacity of list is greater or equal to len(list) + n.
func Int8D2(list [][]int8, index int) [][]int8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int8D3 ensures capacity of list is greater or equal to len(list) + n.
func Int8D3(list [][][]int8, index int) [][][]int8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int8D4 ensures capacity of list is greater or equal to len(list) + n.
func Int8D4(list [][][][]int8, index int) [][][][]int8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int8D5 ensures capacity of list is greater or equal to len(list) + n.
func Int8D5(list [][][][][]int8, index int) [][][][][]int8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int16 ensures capacity of list is greater or equal to len(list) + n.
func Int16(list []int16, index int) []int16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int16D2 ensures capacity of list is greater or equal to len(list) + n.
func Int16D2(list [][]int16, index int) [][]int16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int16D3 ensures capacity of list is greater or equal to len(list) + n.
func Int16D3(list [][][]int16, index int) [][][]int16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int16D4 ensures capacity of list is greater or equal to len(list) + n.
func Int16D4(list [][][][]int16, index int) [][][][]int16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int16D5 ensures capacity of list is greater or equal to len(list) + n.
func Int16D5(list [][][][][]int16, index int) [][][][][]int16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int32 ensures capacity of list is greater or equal to len(list) + n.
func Int32(list []int32, index int) []int32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int32D2 ensures capacity of list is greater or equal to len(list) + n.
func Int32D2(list [][]int32, index int) [][]int32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int32D3 ensures capacity of list is greater or equal to len(list) + n.
func Int32D3(list [][][]int32, index int) [][][]int32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int32D4 ensures capacity of list is greater or equal to len(list) + n.
func Int32D4(list [][][][]int32, index int) [][][][]int32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int32D5 ensures capacity of list is greater or equal to len(list) + n.
func Int32D5(list [][][][][]int32, index int) [][][][][]int32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int64 ensures capacity of list is greater or equal to len(list) + n.
func Int64(list []int64, index int) []int64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int64D2 ensures capacity of list is greater or equal to len(list) + n.
func Int64D2(list [][]int64, index int) [][]int64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int64D3 ensures capacity of list is greater or equal to len(list) + n.
func Int64D3(list [][][]int64, index int) [][][]int64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int64D4 ensures capacity of list is greater or equal to len(list) + n.
func Int64D4(list [][][][]int64, index int) [][][][]int64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Int64D5 ensures capacity of list is greater or equal to len(list) + n.
func Int64D5(list [][][][][]int64, index int) [][][][][]int64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Interface ensures capacity of list is greater or equal to len(list) + n.
func Interface(list []interface{}, index int) []interface{} {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// InterfaceD2 ensures capacity of list is greater or equal to len(list) + n.
func InterfaceD2(list [][]interface{}, index int) [][]interface{} {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// InterfaceD3 ensures capacity of list is greater or equal to len(list) + n.
func InterfaceD3(list [][][]interface{}, index int) [][][]interface{} {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// InterfaceD4 ensures capacity of list is greater or equal to len(list) + n.
func InterfaceD4(list [][][][]interface{}, index int) [][][][]interface{} {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// InterfaceD5 ensures capacity of list is greater or equal to len(list) + n.
func InterfaceD5(list [][][][][]interface{}, index int) [][][][][]interface{} {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Pointer ensures capacity of list is greater or equal to len(list) + n.
func Pointer(list []unsafe.Pointer, index int) []unsafe.Pointer {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// PointerD2 ensures capacity of list is greater or equal to len(list) + n.
func PointerD2(list [][]unsafe.Pointer, index int) [][]unsafe.Pointer {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// PointerD3 ensures capacity of list is greater or equal to len(list) + n.
func PointerD3(list [][][]unsafe.Pointer, index int) [][][]unsafe.Pointer {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// PointerD4 ensures capacity of list is greater or equal to len(list) + n.
func PointerD4(list [][][][]unsafe.Pointer, index int) [][][][]unsafe.Pointer {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// PointerD5 ensures capacity of list is greater or equal to len(list) + n.
func PointerD5(list [][][][][]unsafe.Pointer, index int) [][][][][]unsafe.Pointer {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// Rune ensures capacity of list is greater or equal to len(list) + n.
func Rune(list []rune, index int) []rune {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// RuneD2 ensures capacity of list is greater or equal to len(list) + n.
func RuneD2(list [][]rune, index int) [][]rune {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// RuneD3 ensures capacity of list is greater or equal to len(list) + n.
func RuneD3(list [][][]rune, index int) [][][]rune {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// RuneD4 ensures capacity of list is greater or equal to len(list) + n.
func RuneD4(list [][][][]rune, index int) [][][][]rune {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// RuneD5 ensures capacity of list is greater or equal to len(list) + n.
func RuneD5(list [][][][][]rune, index int) [][][][][]rune {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// String ensures capacity of list is greater or equal to len(list) + n.
func String(list []string, index int) []string {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// StringD2 ensures capacity of list is greater or equal to len(list) + n.
func StringD2(list [][]string, index int) [][]string {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// StringD3 ensures capacity of list is greater or equal to len(list) + n.
func StringD3(list [][][]string, index int) [][][]string {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// StringD4 ensures capacity of list is greater or equal to len(list) + n.
func StringD4(list [][][][]string, index int) [][][][]string {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// StringD5 ensures capacity of list is greater or equal to len(list) + n.
func StringD5(list [][][][][]string, index int) [][][][][]string {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt ensures capacity of list is greater or equal to len(list) + n.
func UInt(list []uint, index int) []uint {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UIntD2 ensures capacity of list is greater or equal to len(list) + n.
func UIntD2(list [][]uint, index int) [][]uint {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UIntD3 ensures capacity of list is greater or equal to len(list) + n.
func UIntD3(list [][][]uint, index int) [][][]uint {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UIntD4 ensures capacity of list is greater or equal to len(list) + n.
func UIntD4(list [][][][]uint, index int) [][][][]uint {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UIntD5 ensures capacity of list is greater or equal to len(list) + n.
func UIntD5(list [][][][][]uint, index int) [][][][][]uint {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt8 ensures capacity of list is greater or equal to len(list) + n.
func UInt8(list []uint8, index int) []uint8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt8D2 ensures capacity of list is greater or equal to len(list) + n.
func UInt8D2(list [][]uint8, index int) [][]uint8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt8D3 ensures capacity of list is greater or equal to len(list) + n.
func UInt8D3(list [][][]uint8, index int) [][][]uint8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt8D4 ensures capacity of list is greater or equal to len(list) + n.
func UInt8D4(list [][][][]uint8, index int) [][][][]uint8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt8D5 ensures capacity of list is greater or equal to len(list) + n.
func UInt8D5(list [][][][][]uint8, index int) [][][][][]uint8 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt16 ensures capacity of list is greater or equal to len(list) + n.
func UInt16(list []uint16, index int) []uint16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt16D2 ensures capacity of list is greater or equal to len(list) + n.
func UInt16D2(list [][]uint16, index int) [][]uint16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt16D3 ensures capacity of list is greater or equal to len(list) + n.
func UInt16D3(list [][][]uint16, index int) [][][]uint16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt16D4 ensures capacity of list is greater or equal to len(list) + n.
func UInt16D4(list [][][][]uint16, index int) [][][][]uint16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt16D5 ensures capacity of list is greater or equal to len(list) + n.
func UInt16D5(list [][][][][]uint16, index int) [][][][][]uint16 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt32 ensures capacity of list is greater or equal to len(list) + n.
func UInt32(list []uint32, index int) []uint32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt32D2 ensures capacity of list is greater or equal to len(list) + n.
func UInt32D2(list [][]uint32, index int) [][]uint32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt32D3 ensures capacity of list is greater or equal to len(list) + n.
func UInt32D3(list [][][]uint32, index int) [][][]uint32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt32D4 ensures capacity of list is greater or equal to len(list) + n.
func UInt32D4(list [][][][]uint32, index int) [][][][]uint32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt32D5 ensures capacity of list is greater or equal to len(list) + n.
func UInt32D5(list [][][][][]uint32, index int) [][][][][]uint32 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt64 ensures capacity of list is greater or equal to len(list) + n.
func UInt64(list []uint64, index int) []uint64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt64D2 ensures capacity of list is greater or equal to len(list) + n.
func UInt64D2(list [][]uint64, index int) [][]uint64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt64D3 ensures capacity of list is greater or equal to len(list) + n.
func UInt64D3(list [][][]uint64, index int) [][][]uint64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt64D4 ensures capacity of list is greater or equal to len(list) + n.
func UInt64D4(list [][][][]uint64, index int) [][][][]uint64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}

// UInt64D5 ensures capacity of list is greater or equal to len(list) + n.
func UInt64D5(list [][][][][]uint64, index int) [][][][][]uint64 {
	if len(list) + n <= cap(list) {
		return list
	}
	newList := make([]bool, len(list), newCap(len(list), len(cap), n))
	copy(newList, list)
	return newList
}
