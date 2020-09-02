/*
 *       Copyright 2018, 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package clone provides the "clone" function for slices of basic types.
package clone

import (
	"unsafe"
)

// Bool makes a copy of the given list. The copy is returned.
func Bool(list []bool) []bool {
	listCopy := make([]bool, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// BoolD2 makes a shallow copy of the given list. The copy is returned.
func BoolD2(list [][]bool) [][]bool {
	listCopy := make([][]bool, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// BoolD2DeepCopy makes a deep copy of the given list. The copy is returned.
func BoolD2DeepCopy(list [][]bool) [][]bool {
	listCopy := make([][]bool, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Bool(li)
	}
	return listCopy
}

// BoolD3 makes a shallow copy of the given list. The copy is returned.
func BoolD3(list [][][]bool) [][][]bool {
	listCopy := make([][][]bool, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// BoolD3DeepCopy makes a deep copy of the given list. The copy is returned.
func BoolD3DeepCopy(list [][][]bool) [][][]bool {
	listCopy := make([][][]bool, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = BoolD2DeepCopy(li)
	}
	return listCopy
}

// BoolD4 makes a shallow copy of the given list. The copy is returned.
func BoolD4(list [][][][]bool) [][][][]bool {
	listCopy := make([][][][]bool, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// BoolD4DeepCopy makes a deep copy of the given list. The copy is returned.
func BoolD4DeepCopy(list [][][][]bool) [][][][]bool {
	listCopy := make([][][][]bool, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = BoolD3DeepCopy(li)
	}
	return listCopy
}

// BoolD5 makes a shallow copy of the given list. The copy is returned.
func BoolD5(list [][][][][]bool) [][][][][]bool {
	listCopy := make([][][][][]bool, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// BoolD5DeepCopy makes a deep copy of the given list. The copy is returned.
func BoolD5DeepCopy(list [][][][][]bool) [][][][][]bool {
	listCopy := make([][][][][]bool, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = BoolD4DeepCopy(li)
	}
	return listCopy
}

// Byte makes a copy of the given list. The copy is returned.
func Byte(list []byte) []byte {
	listCopy := make([]byte, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ByteD2 makes a shallow copy of the given list. The copy is returned.
func ByteD2(list [][]byte) [][]byte {
	listCopy := make([][]byte, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ByteD2DeepCopy makes a deep copy of the given list. The copy is returned.
func ByteD2DeepCopy(list [][]byte) [][]byte {
	listCopy := make([][]byte, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Byte(li)
	}
	return listCopy
}

// ByteD3 makes a shallow copy of the given list. The copy is returned.
func ByteD3(list [][][]byte) [][][]byte {
	listCopy := make([][][]byte, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ByteD3DeepCopy makes a deep copy of the given list. The copy is returned.
func ByteD3DeepCopy(list [][][]byte) [][][]byte {
	listCopy := make([][][]byte, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = ByteD2DeepCopy(li)
	}
	return listCopy
}

// ByteD4 makes a shallow copy of the given list. The copy is returned.
func ByteD4(list [][][][]byte) [][][][]byte {
	listCopy := make([][][][]byte, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ByteD4DeepCopy makes a deep copy of the given list. The copy is returned.
func ByteD4DeepCopy(list [][][][]byte) [][][][]byte {
	listCopy := make([][][][]byte, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = ByteD3DeepCopy(li)
	}
	return listCopy
}

// ByteD5 makes a shallow copy of the given list. The copy is returned.
func ByteD5(list [][][][][]byte) [][][][][]byte {
	listCopy := make([][][][][]byte, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ByteD5DeepCopy makes a deep copy of the given list. The copy is returned.
func ByteD5DeepCopy(list [][][][][]byte) [][][][][]byte {
	listCopy := make([][][][][]byte, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = ByteD4DeepCopy(li)
	}
	return listCopy
}

// Complex64 makes a copy of the given list. The copy is returned.
func Complex64(list []complex64) []complex64 {
	listCopy := make([]complex64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex64D2 makes a shallow copy of the given list. The copy is returned.
func Complex64D2(list [][]complex64) [][]complex64 {
	listCopy := make([][]complex64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex64D2DeepCopy makes a deep copy of the given list. The copy is returned.
func Complex64D2DeepCopy(list [][]complex64) [][]complex64 {
	listCopy := make([][]complex64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Complex64(li)
	}
	return listCopy
}

// Complex64D3 makes a shallow copy of the given list. The copy is returned.
func Complex64D3(list [][][]complex64) [][][]complex64 {
	listCopy := make([][][]complex64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex64D3DeepCopy makes a deep copy of the given list. The copy is returned.
func Complex64D3DeepCopy(list [][][]complex64) [][][]complex64 {
	listCopy := make([][][]complex64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Complex64D2DeepCopy(li)
	}
	return listCopy
}

// Complex64D4 makes a shallow copy of the given list. The copy is returned.
func Complex64D4(list [][][][]complex64) [][][][]complex64 {
	listCopy := make([][][][]complex64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex64D4DeepCopy makes a deep copy of the given list. The copy is returned.
func Complex64D4DeepCopy(list [][][][]complex64) [][][][]complex64 {
	listCopy := make([][][][]complex64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Complex64D3DeepCopy(li)
	}
	return listCopy
}

// Complex64D5 makes a shallow copy of the given list. The copy is returned.
func Complex64D5(list [][][][][]complex64) [][][][][]complex64 {
	listCopy := make([][][][][]complex64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex64D5DeepCopy makes a deep copy of the given list. The copy is returned.
func Complex64D5DeepCopy(list [][][][][]complex64) [][][][][]complex64 {
	listCopy := make([][][][][]complex64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Complex64D4DeepCopy(li)
	}
	return listCopy
}

// Complex128 makes a copy of the given list. The copy is returned.
func Complex128(list []complex128) []complex128 {
	listCopy := make([]complex128, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex128D2 makes a shallow copy of the given list. The copy is returned.
func Complex128D2(list [][]complex128) [][]complex128 {
	listCopy := make([][]complex128, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex128D2DeepCopy makes a deep copy of the given list. The copy is returned.
func Complex128D2DeepCopy(list [][]complex128) [][]complex128 {
	listCopy := make([][]complex128, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Complex128(li)
	}
	return listCopy
}

// Complex128D3 makes a shallow copy of the given list. The copy is returned.
func Complex128D3(list [][][]complex128) [][][]complex128 {
	listCopy := make([][][]complex128, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex128D3DeepCopy makes a deep copy of the given list. The copy is returned.
func Complex128D3DeepCopy(list [][][]complex128) [][][]complex128 {
	listCopy := make([][][]complex128, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Complex128D2DeepCopy(li)
	}
	return listCopy
}

// Complex128D4 makes a shallow copy of the given list. The copy is returned.
func Complex128D4(list [][][][]complex128) [][][][]complex128 {
	listCopy := make([][][][]complex128, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex128D4DeepCopy makes a deep copy of the given list. The copy is returned.
func Complex128D4DeepCopy(list [][][][]complex128) [][][][]complex128 {
	listCopy := make([][][][]complex128, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Complex128D3DeepCopy(li)
	}
	return listCopy
}

// Complex128D5 makes a shallow copy of the given list. The copy is returned.
func Complex128D5(list [][][][][]complex128) [][][][][]complex128 {
	listCopy := make([][][][][]complex128, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Complex128D5DeepCopy makes a deep copy of the given list. The copy is returned.
func Complex128D5DeepCopy(list [][][][][]complex128) [][][][][]complex128 {
	listCopy := make([][][][][]complex128, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Complex128D4DeepCopy(li)
	}
	return listCopy
}

// Error makes a copy of the given list. The copy is returned.
func Error(list []error) []error {
	listCopy := make([]error, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ErrorD2 makes a shallow copy of the given list. The copy is returned.
func ErrorD2(list [][]error) [][]error {
	listCopy := make([][]error, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ErrorD2DeepCopy makes a deep copy of the given list. The copy is returned.
func ErrorD2DeepCopy(list [][]error) [][]error {
	listCopy := make([][]error, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Error(li)
	}
	return listCopy
}

// ErrorD3 makes a shallow copy of the given list. The copy is returned.
func ErrorD3(list [][][]error) [][][]error {
	listCopy := make([][][]error, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ErrorD3DeepCopy makes a deep copy of the given list. The copy is returned.
func ErrorD3DeepCopy(list [][][]error) [][][]error {
	listCopy := make([][][]error, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = ErrorD2DeepCopy(li)
	}
	return listCopy
}

// ErrorD4 makes a shallow copy of the given list. The copy is returned.
func ErrorD4(list [][][][]error) [][][][]error {
	listCopy := make([][][][]error, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ErrorD4DeepCopy makes a deep copy of the given list. The copy is returned.
func ErrorD4DeepCopy(list [][][][]error) [][][][]error {
	listCopy := make([][][][]error, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = ErrorD3DeepCopy(li)
	}
	return listCopy
}

// ErrorD5 makes a shallow copy of the given list. The copy is returned.
func ErrorD5(list [][][][][]error) [][][][][]error {
	listCopy := make([][][][][]error, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// ErrorD5DeepCopy makes a deep copy of the given list. The copy is returned.
func ErrorD5DeepCopy(list [][][][][]error) [][][][][]error {
	listCopy := make([][][][][]error, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = ErrorD4DeepCopy(li)
	}
	return listCopy
}

// Float32 makes a copy of the given list. The copy is returned.
func Float32(list []float32) []float32 {
	listCopy := make([]float32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float32D2 makes a shallow copy of the given list. The copy is returned.
func Float32D2(list [][]float32) [][]float32 {
	listCopy := make([][]float32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float32D2DeepCopy makes a deep copy of the given list. The copy is returned.
func Float32D2DeepCopy(list [][]float32) [][]float32 {
	listCopy := make([][]float32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Float32(li)
	}
	return listCopy
}

// Float32D3 makes a shallow copy of the given list. The copy is returned.
func Float32D3(list [][][]float32) [][][]float32 {
	listCopy := make([][][]float32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float32D3DeepCopy makes a deep copy of the given list. The copy is returned.
func Float32D3DeepCopy(list [][][]float32) [][][]float32 {
	listCopy := make([][][]float32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Float32D2DeepCopy(li)
	}
	return listCopy
}

// Float32D4 makes a shallow copy of the given list. The copy is returned.
func Float32D4(list [][][][]float32) [][][][]float32 {
	listCopy := make([][][][]float32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float32D4DeepCopy makes a deep copy of the given list. The copy is returned.
func Float32D4DeepCopy(list [][][][]float32) [][][][]float32 {
	listCopy := make([][][][]float32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Float32D3DeepCopy(li)
	}
	return listCopy
}

// Float32D5 makes a shallow copy of the given list. The copy is returned.
func Float32D5(list [][][][][]float32) [][][][][]float32 {
	listCopy := make([][][][][]float32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float32D5DeepCopy makes a deep copy of the given list. The copy is returned.
func Float32D5DeepCopy(list [][][][][]float32) [][][][][]float32 {
	listCopy := make([][][][][]float32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Float32D4DeepCopy(li)
	}
	return listCopy
}

// Float64 makes a copy of the given list. The copy is returned.
func Float64(list []float64) []float64 {
	listCopy := make([]float64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float64D2 makes a shallow copy of the given list. The copy is returned.
func Float64D2(list [][]float64) [][]float64 {
	listCopy := make([][]float64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float64D2DeepCopy makes a deep copy of the given list. The copy is returned.
func Float64D2DeepCopy(list [][]float64) [][]float64 {
	listCopy := make([][]float64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Float64(li)
	}
	return listCopy
}

// Float64D3 makes a shallow copy of the given list. The copy is returned.
func Float64D3(list [][][]float64) [][][]float64 {
	listCopy := make([][][]float64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float64D3DeepCopy makes a deep copy of the given list. The copy is returned.
func Float64D3DeepCopy(list [][][]float64) [][][]float64 {
	listCopy := make([][][]float64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Float64D2DeepCopy(li)
	}
	return listCopy
}

// Float64D4 makes a shallow copy of the given list. The copy is returned.
func Float64D4(list [][][][]float64) [][][][]float64 {
	listCopy := make([][][][]float64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float64D4DeepCopy makes a deep copy of the given list. The copy is returned.
func Float64D4DeepCopy(list [][][][]float64) [][][][]float64 {
	listCopy := make([][][][]float64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Float64D3DeepCopy(li)
	}
	return listCopy
}

// Float64D5 makes a shallow copy of the given list. The copy is returned.
func Float64D5(list [][][][][]float64) [][][][][]float64 {
	listCopy := make([][][][][]float64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Float64D5DeepCopy makes a deep copy of the given list. The copy is returned.
func Float64D5DeepCopy(list [][][][][]float64) [][][][][]float64 {
	listCopy := make([][][][][]float64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Float64D4DeepCopy(li)
	}
	return listCopy
}

// Int makes a copy of the given list. The copy is returned.
func Int(list []int) []int {
	listCopy := make([]int, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// IntD2 makes a shallow copy of the given list. The copy is returned.
func IntD2(list [][]int) [][]int {
	listCopy := make([][]int, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// IntD2DeepCopy makes a deep copy of the given list. The copy is returned.
func IntD2DeepCopy(list [][]int) [][]int {
	listCopy := make([][]int, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int(li)
	}
	return listCopy
}

// IntD3 makes a shallow copy of the given list. The copy is returned.
func IntD3(list [][][]int) [][][]int {
	listCopy := make([][][]int, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// IntD3DeepCopy makes a deep copy of the given list. The copy is returned.
func IntD3DeepCopy(list [][][]int) [][][]int {
	listCopy := make([][][]int, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = IntD2DeepCopy(li)
	}
	return listCopy
}

// IntD4 makes a shallow copy of the given list. The copy is returned.
func IntD4(list [][][][]int) [][][][]int {
	listCopy := make([][][][]int, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// IntD4DeepCopy makes a deep copy of the given list. The copy is returned.
func IntD4DeepCopy(list [][][][]int) [][][][]int {
	listCopy := make([][][][]int, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = IntD3DeepCopy(li)
	}
	return listCopy
}

// IntD5 makes a shallow copy of the given list. The copy is returned.
func IntD5(list [][][][][]int) [][][][][]int {
	listCopy := make([][][][][]int, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// IntD5DeepCopy makes a deep copy of the given list. The copy is returned.
func IntD5DeepCopy(list [][][][][]int) [][][][][]int {
	listCopy := make([][][][][]int, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = IntD4DeepCopy(li)
	}
	return listCopy
}

// Int8 makes a copy of the given list. The copy is returned.
func Int8(list []int8) []int8 {
	listCopy := make([]int8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int8D2 makes a shallow copy of the given list. The copy is returned.
func Int8D2(list [][]int8) [][]int8 {
	listCopy := make([][]int8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int8D2DeepCopy makes a deep copy of the given list. The copy is returned.
func Int8D2DeepCopy(list [][]int8) [][]int8 {
	listCopy := make([][]int8, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int8(li)
	}
	return listCopy
}

// Int8D3 makes a shallow copy of the given list. The copy is returned.
func Int8D3(list [][][]int8) [][][]int8 {
	listCopy := make([][][]int8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int8D3DeepCopy makes a deep copy of the given list. The copy is returned.
func Int8D3DeepCopy(list [][][]int8) [][][]int8 {
	listCopy := make([][][]int8, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int8D2DeepCopy(li)
	}
	return listCopy
}

// Int8D4 makes a shallow copy of the given list. The copy is returned.
func Int8D4(list [][][][]int8) [][][][]int8 {
	listCopy := make([][][][]int8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int8D4DeepCopy makes a deep copy of the given list. The copy is returned.
func Int8D4DeepCopy(list [][][][]int8) [][][][]int8 {
	listCopy := make([][][][]int8, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int8D3DeepCopy(li)
	}
	return listCopy
}

// Int8D5 makes a shallow copy of the given list. The copy is returned.
func Int8D5(list [][][][][]int8) [][][][][]int8 {
	listCopy := make([][][][][]int8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int8D5DeepCopy makes a deep copy of the given list. The copy is returned.
func Int8D5DeepCopy(list [][][][][]int8) [][][][][]int8 {
	listCopy := make([][][][][]int8, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int8D4DeepCopy(li)
	}
	return listCopy
}

// Int16 makes a copy of the given list. The copy is returned.
func Int16(list []int16) []int16 {
	listCopy := make([]int16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int16D2 makes a shallow copy of the given list. The copy is returned.
func Int16D2(list [][]int16) [][]int16 {
	listCopy := make([][]int16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int16D2DeepCopy makes a deep copy of the given list. The copy is returned.
func Int16D2DeepCopy(list [][]int16) [][]int16 {
	listCopy := make([][]int16, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int16(li)
	}
	return listCopy
}

// Int16D3 makes a shallow copy of the given list. The copy is returned.
func Int16D3(list [][][]int16) [][][]int16 {
	listCopy := make([][][]int16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int16D3DeepCopy makes a deep copy of the given list. The copy is returned.
func Int16D3DeepCopy(list [][][]int16) [][][]int16 {
	listCopy := make([][][]int16, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int16D2DeepCopy(li)
	}
	return listCopy
}

// Int16D4 makes a shallow copy of the given list. The copy is returned.
func Int16D4(list [][][][]int16) [][][][]int16 {
	listCopy := make([][][][]int16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int16D4DeepCopy makes a deep copy of the given list. The copy is returned.
func Int16D4DeepCopy(list [][][][]int16) [][][][]int16 {
	listCopy := make([][][][]int16, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int16D3DeepCopy(li)
	}
	return listCopy
}

// Int16D5 makes a shallow copy of the given list. The copy is returned.
func Int16D5(list [][][][][]int16) [][][][][]int16 {
	listCopy := make([][][][][]int16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int16D5DeepCopy makes a deep copy of the given list. The copy is returned.
func Int16D5DeepCopy(list [][][][][]int16) [][][][][]int16 {
	listCopy := make([][][][][]int16, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int16D4DeepCopy(li)
	}
	return listCopy
}

// Int32 makes a copy of the given list. The copy is returned.
func Int32(list []int32) []int32 {
	listCopy := make([]int32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int32D2 makes a shallow copy of the given list. The copy is returned.
func Int32D2(list [][]int32) [][]int32 {
	listCopy := make([][]int32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int32D2DeepCopy makes a deep copy of the given list. The copy is returned.
func Int32D2DeepCopy(list [][]int32) [][]int32 {
	listCopy := make([][]int32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int32(li)
	}
	return listCopy
}

// Int32D3 makes a shallow copy of the given list. The copy is returned.
func Int32D3(list [][][]int32) [][][]int32 {
	listCopy := make([][][]int32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int32D3DeepCopy makes a deep copy of the given list. The copy is returned.
func Int32D3DeepCopy(list [][][]int32) [][][]int32 {
	listCopy := make([][][]int32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int32D2DeepCopy(li)
	}
	return listCopy
}

// Int32D4 makes a shallow copy of the given list. The copy is returned.
func Int32D4(list [][][][]int32) [][][][]int32 {
	listCopy := make([][][][]int32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int32D4DeepCopy makes a deep copy of the given list. The copy is returned.
func Int32D4DeepCopy(list [][][][]int32) [][][][]int32 {
	listCopy := make([][][][]int32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int32D3DeepCopy(li)
	}
	return listCopy
}

// Int32D5 makes a shallow copy of the given list. The copy is returned.
func Int32D5(list [][][][][]int32) [][][][][]int32 {
	listCopy := make([][][][][]int32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int32D5DeepCopy makes a deep copy of the given list. The copy is returned.
func Int32D5DeepCopy(list [][][][][]int32) [][][][][]int32 {
	listCopy := make([][][][][]int32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int32D4DeepCopy(li)
	}
	return listCopy
}

// Int64 makes a copy of the given list. The copy is returned.
func Int64(list []int64) []int64 {
	listCopy := make([]int64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int64D2 makes a shallow copy of the given list. The copy is returned.
func Int64D2(list [][]int64) [][]int64 {
	listCopy := make([][]int64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int64D2DeepCopy makes a deep copy of the given list. The copy is returned.
func Int64D2DeepCopy(list [][]int64) [][]int64 {
	listCopy := make([][]int64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int64(li)
	}
	return listCopy
}

// Int64D3 makes a shallow copy of the given list. The copy is returned.
func Int64D3(list [][][]int64) [][][]int64 {
	listCopy := make([][][]int64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int64D3DeepCopy makes a deep copy of the given list. The copy is returned.
func Int64D3DeepCopy(list [][][]int64) [][][]int64 {
	listCopy := make([][][]int64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int64D2DeepCopy(li)
	}
	return listCopy
}

// Int64D4 makes a shallow copy of the given list. The copy is returned.
func Int64D4(list [][][][]int64) [][][][]int64 {
	listCopy := make([][][][]int64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int64D4DeepCopy makes a deep copy of the given list. The copy is returned.
func Int64D4DeepCopy(list [][][][]int64) [][][][]int64 {
	listCopy := make([][][][]int64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int64D3DeepCopy(li)
	}
	return listCopy
}

// Int64D5 makes a shallow copy of the given list. The copy is returned.
func Int64D5(list [][][][][]int64) [][][][][]int64 {
	listCopy := make([][][][][]int64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// Int64D5DeepCopy makes a deep copy of the given list. The copy is returned.
func Int64D5DeepCopy(list [][][][][]int64) [][][][][]int64 {
	listCopy := make([][][][][]int64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Int64D4DeepCopy(li)
	}
	return listCopy
}

// Interface makes a copy of the given list. The copy is returned.
// The object referenced by interface is not cloned.
func Interface(list []interface{}) []interface{} {
	listCopy := make([]interface{}, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// InterfaceD2 makes a shallow copy of the given list. The copy is returned.
func InterfaceD2(list [][]interface{}) [][]interface{} {
	listCopy := make([][]interface{}, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// InterfaceD2DeepCopy makes a deep copy of the given list. The copy is returned.
// The object referenced by interface is not cloned.
func InterfaceD2DeepCopy(list [][]interface{}) [][]interface{} {
	listCopy := make([][]interface{}, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Interface(li)
	}
	return listCopy
}

// InterfaceD3 makes a shallow copy of the given list. The copy is returned.
func InterfaceD3(list [][][]interface{}) [][][]interface{} {
	listCopy := make([][][]interface{}, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// InterfaceD3DeepCopy makes a deep copy of the given list. The copy is returned.
// The object referenced by interface is not cloned.
func InterfaceD3DeepCopy(list [][][]interface{}) [][][]interface{} {
	listCopy := make([][][]interface{}, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = InterfaceD2DeepCopy(li)
	}
	return listCopy
}

// InterfaceD4 makes a shallow copy of the given list. The copy is returned.
func InterfaceD4(list [][][][]interface{}) [][][][]interface{} {
	listCopy := make([][][][]interface{}, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// InterfaceD4DeepCopy makes a deep copy of the given list. The copy is returned.
// The object referenced by interface is not cloned.
func InterfaceD4DeepCopy(list [][][][]interface{}) [][][][]interface{} {
	listCopy := make([][][][]interface{}, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = InterfaceD3DeepCopy(li)
	}
	return listCopy
}

// InterfaceD5 makes a shallow copy of the given list. The copy is returned.
func InterfaceD5(list [][][][][]interface{}) [][][][][]interface{} {
	listCopy := make([][][][][]interface{}, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// InterfaceD5DeepCopy makes a deep copy of the given list. The copy is returned.
// The object referenced by interface is not cloned.
func InterfaceD5DeepCopy(list [][][][][]interface{}) [][][][][]interface{} {
	listCopy := make([][][][][]interface{}, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = InterfaceD4DeepCopy(li)
	}
	return listCopy
}

// Pointer makes a copy of the given list. The copy is returned.
func Pointer(list []unsafe.Pointer) []unsafe.Pointer {
	listCopy := make([]unsafe.Pointer, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// PointerD2 makes a shallow copy of the given list. The copy is returned.
func PointerD2(list [][]unsafe.Pointer) [][]unsafe.Pointer {
	listCopy := make([][]unsafe.Pointer, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// PointerD2DeepCopy makes a deep copy of the given list. The copy is returned.
// The data the unsafe pointer is pointing to is not cloned.
func PointerD2DeepCopy(list [][]unsafe.Pointer) [][]unsafe.Pointer {
	listCopy := make([][]unsafe.Pointer, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Pointer(li)
	}
	return listCopy
}

// PointerD3 makes a shallow copy of the given list. The copy is returned.
func PointerD3(list [][][]unsafe.Pointer) [][][]unsafe.Pointer {
	listCopy := make([][][]unsafe.Pointer, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// PointerD3DeepCopy makes a deep copy of the given list. The copy is returned.
// The data the unsafe pointer is pointing to is not cloned.
func PointerD3DeepCopy(list [][][]unsafe.Pointer) [][][]unsafe.Pointer {
	listCopy := make([][][]unsafe.Pointer, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = PointerD2DeepCopy(li)
	}
	return listCopy
}

// PointerD4 makes a shallow copy of the given list. The copy is returned.
func PointerD4(list [][][][]unsafe.Pointer) [][][][]unsafe.Pointer {
	listCopy := make([][][][]unsafe.Pointer, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// PointerD4DeepCopy makes a deep copy of the given list. The copy is returned.
// The data the unsafe pointer is pointing to is not cloned.
func PointerD4DeepCopy(list [][][][]unsafe.Pointer) [][][][]unsafe.Pointer {
	listCopy := make([][][][]unsafe.Pointer, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = PointerD3DeepCopy(li)
	}
	return listCopy
}

// PointerD5 makes a shallow copy of the given list. The copy is returned.
func PointerD5(list [][][][][]unsafe.Pointer) [][][][][]unsafe.Pointer {
	listCopy := make([][][][][]unsafe.Pointer, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// PointerD5DeepCopy makes a deep copy of the given list. The copy is returned.
// The data the unsafe pointer is pointing to is not cloned.
func PointerD5DeepCopy(list [][][][][]unsafe.Pointer) [][][][][]unsafe.Pointer {
	listCopy := make([][][][][]unsafe.Pointer, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = PointerD4DeepCopy(li)
	}
	return listCopy
}

// Rune makes a copy of the given list. The copy is returned.
func Rune(list []rune) []rune {
	listCopy := make([]rune, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// RuneD2 makes a shallow copy of the given list. The copy is returned.
func RuneD2(list [][]rune) [][]rune {
	listCopy := make([][]rune, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// RuneD2DeepCopy makes a deep copy of the given list. The copy is returned.
func RuneD2DeepCopy(list [][]rune) [][]rune {
	listCopy := make([][]rune, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = Rune(li)
	}
	return listCopy
}

// RuneD3 makes a shallow copy of the given list. The copy is returned.
func RuneD3(list [][][]rune) [][][]rune {
	listCopy := make([][][]rune, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// RuneD3DeepCopy makes a deep copy of the given list. The copy is returned.
func RuneD3DeepCopy(list [][][]rune) [][][]rune {
	listCopy := make([][][]rune, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = RuneD2DeepCopy(li)
	}
	return listCopy
}

// RuneD4 makes a shallow copy of the given list. The copy is returned.
func RuneD4(list [][][][]rune) [][][][]rune {
	listCopy := make([][][][]rune, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// RuneD4DeepCopy makes a deep copy of the given list. The copy is returned.
func RuneD4DeepCopy(list [][][][]rune) [][][][]rune {
	listCopy := make([][][][]rune, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = RuneD3DeepCopy(li)
	}
	return listCopy
}

// RuneD5 makes a shallow copy of the given list. The copy is returned.
func RuneD5(list [][][][][]rune) [][][][][]rune {
	listCopy := make([][][][][]rune, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// RuneD5DeepCopy makes a deep copy of the given list. The copy is returned.
func RuneD5DeepCopy(list [][][][][]rune) [][][][][]rune {
	listCopy := make([][][][][]rune, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = RuneD4DeepCopy(li)
	}
	return listCopy
}

// String makes a copy of the given list. The copy is returned.
func String(list []string) []string {
	listCopy := make([]string, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// StringD2 makes a shallow copy of the given list. The copy is returned.
func StringD2(list [][]string) [][]string {
	listCopy := make([][]string, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// StringD2DeepCopy makes a deep copy of the given list. The copy is returned.
func StringD2DeepCopy(list [][]string) [][]string {
	listCopy := make([][]string, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = String(li)
	}
	return listCopy
}

// StringD3 makes a shallow copy of the given list. The copy is returned.
func StringD3(list [][][]string) [][][]string {
	listCopy := make([][][]string, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// StringD3DeepCopy makes a deep copy of the given list. The copy is returned.
func StringD3DeepCopy(list [][][]string) [][][]string {
	listCopy := make([][][]string, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = StringD2DeepCopy(li)
	}
	return listCopy
}

// StringD4 makes a shallow copy of the given list. The copy is returned.
func StringD4(list [][][][]string) [][][][]string {
	listCopy := make([][][][]string, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// StringD4DeepCopy makes a deep copy of the given list. The copy is returned.
func StringD4DeepCopy(list [][][][]string) [][][][]string {
	listCopy := make([][][][]string, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = StringD3DeepCopy(li)
	}
	return listCopy
}

// StringD5 makes a shallow copy of the given list. The copy is returned.
func StringD5(list [][][][][]string) [][][][][]string {
	listCopy := make([][][][][]string, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// StringD5DeepCopy makes a deep copy of the given list. The copy is returned.
func StringD5DeepCopy(list [][][][][]string) [][][][][]string {
	listCopy := make([][][][][]string, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = StringD4DeepCopy(li)
	}
	return listCopy
}

// UInt makes a copy of the given list. The copy is returned.
func UInt(list []uint) []uint {
	listCopy := make([]uint, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UIntD2 makes a shallow copy of the given list. The copy is returned.
func UIntD2(list [][]uint) [][]uint {
	listCopy := make([][]uint, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UIntD2DeepCopy makes a deep copy of the given list. The copy is returned.
func UIntD2DeepCopy(list [][]uint) [][]uint {
	listCopy := make([][]uint, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt(li)
	}
	return listCopy
}

// UIntD3 makes a shallow copy of the given list. The copy is returned.
func UIntD3(list [][][]uint) [][][]uint {
	listCopy := make([][][]uint, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UIntD3DeepCopy makes a deep copy of the given list. The copy is returned.
func UIntD3DeepCopy(list [][][]uint) [][][]uint {
	listCopy := make([][][]uint, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UIntD2DeepCopy(li)
	}
	return listCopy
}

// UIntD4 makes a shallow copy of the given list. The copy is returned.
func UIntD4(list [][][][]uint) [][][][]uint {
	listCopy := make([][][][]uint, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UIntD4DeepCopy makes a deep copy of the given list. The copy is returned.
func UIntD4DeepCopy(list [][][][]uint) [][][][]uint {
	listCopy := make([][][][]uint, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UIntD3DeepCopy(li)
	}
	return listCopy
}

// UIntD5 makes a shallow copy of the given list. The copy is returned.
func UIntD5(list [][][][][]uint) [][][][][]uint {
	listCopy := make([][][][][]uint, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UIntD5DeepCopy makes a deep copy of the given list. The copy is returned.
func UIntD5DeepCopy(list [][][][][]uint) [][][][][]uint {
	listCopy := make([][][][][]uint, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UIntD4DeepCopy(li)
	}
	return listCopy
}

// UInt8 makes a copy of the given list. The copy is returned.
func UInt8(list []uint8) []uint8 {
	listCopy := make([]uint8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt8D2 makes a shallow copy of the given list. The copy is returned.
func UInt8D2(list [][]uint8) [][]uint8 {
	listCopy := make([][]uint8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt8D2DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt8D2DeepCopy(list [][]uint8) [][]uint8 {
	listCopy := make([][]uint8, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt8(li)
	}
	return listCopy
}

// UInt8D3 makes a shallow copy of the given list. The copy is returned.
func UInt8D3(list [][][]uint8) [][][]uint8 {
	listCopy := make([][][]uint8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt8D3DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt8D3DeepCopy(list [][][]uint8) [][][]uint8 {
	listCopy := make([][][]uint8, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt8D2DeepCopy(li)
	}
	return listCopy
}

// UInt8D4 makes a shallow copy of the given list. The copy is returned.
func UInt8D4(list [][][][]uint8) [][][][]uint8 {
	listCopy := make([][][][]uint8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt8D4DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt8D4DeepCopy(list [][][][]uint8) [][][][]uint8 {
	listCopy := make([][][][]uint8, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt8D3DeepCopy(li)
	}
	return listCopy
}

// UInt8D5 makes a shallow copy of the given list. The copy is returned.
func UInt8D5(list [][][][][]uint8) [][][][][]uint8 {
	listCopy := make([][][][][]uint8, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt8D5DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt8D5DeepCopy(list [][][][][]uint8) [][][][][]uint8 {
	listCopy := make([][][][][]uint8, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt8D4DeepCopy(li)
	}
	return listCopy
}

// UInt16 makes a copy of the given list. The copy is returned.
func UInt16(list []uint16) []uint16 {
	listCopy := make([]uint16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt16D2 makes a shallow copy of the given list. The copy is returned.
func UInt16D2(list [][]uint16) [][]uint16 {
	listCopy := make([][]uint16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt16D2DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt16D2DeepCopy(list [][]uint16) [][]uint16 {
	listCopy := make([][]uint16, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt16(li)
	}
	return listCopy
}

// UInt16D3 makes a shallow copy of the given list. The copy is returned.
func UInt16D3(list [][][]uint16) [][][]uint16 {
	listCopy := make([][][]uint16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt16D3DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt16D3DeepCopy(list [][][]uint16) [][][]uint16 {
	listCopy := make([][][]uint16, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt16D2DeepCopy(li)
	}
	return listCopy
}

// UInt16D4 makes a shallow copy of the given list. The copy is returned.
func UInt16D4(list [][][][]uint16) [][][][]uint16 {
	listCopy := make([][][][]uint16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt16D4DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt16D4DeepCopy(list [][][][]uint16) [][][][]uint16 {
	listCopy := make([][][][]uint16, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt16D3DeepCopy(li)
	}
	return listCopy
}

// UInt16D5 makes a shallow copy of the given list. The copy is returned.
func UInt16D5(list [][][][][]uint16) [][][][][]uint16 {
	listCopy := make([][][][][]uint16, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt16D5DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt16D5DeepCopy(list [][][][][]uint16) [][][][][]uint16 {
	listCopy := make([][][][][]uint16, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt16D4DeepCopy(li)
	}
	return listCopy
}

// UInt32 makes a copy of the given list. The copy is returned.
func UInt32(list []uint32) []uint32 {
	listCopy := make([]uint32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt32D2 makes a shallow copy of the given list. The copy is returned.
func UInt32D2(list [][]uint32) [][]uint32 {
	listCopy := make([][]uint32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt32D2DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt32D2DeepCopy(list [][]uint32) [][]uint32 {
	listCopy := make([][]uint32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt32(li)
	}
	return listCopy
}

// UInt32D3 makes a shallow copy of the given list. The copy is returned.
func UInt32D3(list [][][]uint32) [][][]uint32 {
	listCopy := make([][][]uint32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt32D3DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt32D3DeepCopy(list [][][]uint32) [][][]uint32 {
	listCopy := make([][][]uint32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt32D2DeepCopy(li)
	}
	return listCopy
}

// UInt32D4 makes a shallow copy of the given list. The copy is returned.
func UInt32D4(list [][][][]uint32) [][][][]uint32 {
	listCopy := make([][][][]uint32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt32D4DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt32D4DeepCopy(list [][][][]uint32) [][][][]uint32 {
	listCopy := make([][][][]uint32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt32D3DeepCopy(li)
	}
	return listCopy
}

// UInt32D5 makes a shallow copy of the given list. The copy is returned.
func UInt32D5(list [][][][][]uint32) [][][][][]uint32 {
	listCopy := make([][][][][]uint32, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt32D5DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt32D5DeepCopy(list [][][][][]uint32) [][][][][]uint32 {
	listCopy := make([][][][][]uint32, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt32D4DeepCopy(li)
	}
	return listCopy
}

// UInt64 makes a copy of the given list. The copy is returned.
func UInt64(list []uint64) []uint64 {
	listCopy := make([]uint64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt64D2 makes a shallow copy of the given list. The copy is returned.
func UInt64D2(list [][]uint64) [][]uint64 {
	listCopy := make([][]uint64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt64D2DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt64D2DeepCopy(list [][]uint64) [][]uint64 {
	listCopy := make([][]uint64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt64(li)
	}
	return listCopy
}

// UInt64D3 makes a shallow copy of the given list. The copy is returned.
func UInt64D3(list [][][]uint64) [][][]uint64 {
	listCopy := make([][][]uint64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt64D3DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt64D3DeepCopy(list [][][]uint64) [][][]uint64 {
	listCopy := make([][][]uint64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt64D2DeepCopy(li)
	}
	return listCopy
}

// UInt64D4 makes a shallow copy of the given list. The copy is returned.
func UInt64D4(list [][][][]uint64) [][][][]uint64 {
	listCopy := make([][][][]uint64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt64D4DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt64D4DeepCopy(list [][][][]uint64) [][][][]uint64 {
	listCopy := make([][][][]uint64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt64D3DeepCopy(li)
	}
	return listCopy
}

// UInt64D5 makes a shallow copy of the given list. The copy is returned.
func UInt64D5(list [][][][][]uint64) [][][][][]uint64 {
	listCopy := make([][][][][]uint64, len(list), cap(list))
	copy(listCopy, list)
	return listCopy
}

// UInt64D5DeepCopy makes a deep copy of the given list. The copy is returned.
func UInt64D5DeepCopy(list [][][][][]uint64) [][][][][]uint64 {
	listCopy := make([][][][][]uint64, len(list), cap(list))
	for i, li := range list {
		listCopy[i] = UInt64D4DeepCopy(li)
	}
	return listCopy
}
