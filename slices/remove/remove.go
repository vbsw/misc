/*
 *       Copyright 2018, 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package remove provides the "remove" function for slices of basic types.
package remove

import (
	"unsafe"
)

// Bool removes an element from list.
func Bool(list []bool, index int) []bool {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// BoolD2 removes an element from list.
func BoolD2(list [][]bool, index int) [][]bool {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// BoolD3 removes an element from list.
func BoolD3(list [][][]bool, index int) [][][]bool {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// BoolD4 removes an element from list.
func BoolD4(list [][][][]bool, index int) [][][][]bool {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// BoolD5 removes an element from list.
func BoolD5(list [][][][][]bool, index int) [][][][][]bool {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// BoolN removes n elements from list.
func BoolN(list []bool, index, n int) []bool {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// BoolND2 removes n elements from list.
func BoolND2(list [][]bool, index, n int) [][]bool {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// BoolND3 removes n elements from list.
func BoolND3(list [][][]bool, index, n int) [][][]bool {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// BoolND4 removes n elements from list.
func BoolND4(list [][][][]bool, index, n int) [][][][]bool {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// BoolND5 removes n elements from list.
func BoolND5(list [][][][][]bool, index, n int) [][][][][]bool {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Byte removes an element from list.
func Byte(list []byte, index int) []byte {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ByteD2 removes an element from list.
func ByteD2(list [][]byte, index int) [][]byte {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ByteD3 removes an element from list.
func ByteD3(list [][][]byte, index int) [][][]byte {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ByteD4 removes an element from list.
func ByteD4(list [][][][]byte, index int) [][][][]byte {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ByteD5 removes an element from list.
func ByteD5(list [][][][][]byte, index int) [][][][][]byte {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ByteN removes n elements from list.
func ByteN(list []byte, index, n int) []byte {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// ByteND2 removes n elements from list.
func ByteND2(list [][]byte, index, n int) [][]byte {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// ByteND3 removes n elements from list.
func ByteND3(list [][][]byte, index, n int) [][][]byte {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// ByteND4 removes n elements from list.
func ByteND4(list [][][][]byte, index, n int) [][][][]byte {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// ByteND5 removes n elements from list.
func ByteND5(list [][][][][]byte, index, n int) [][][][][]byte {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex64 removes an element from list.
func Complex64(list []complex64, index int) []complex64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex64D2 removes an element from list.
func Complex64D2(list [][]complex64, index int) [][]complex64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex64D3 removes an element from list.
func Complex64D3(list [][][]complex64, index int) [][][]complex64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex64D4 removes an element from list.
func Complex64D4(list [][][][]complex64, index int) [][][][]complex64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex64D5 removes an element from list.
func Complex64D5(list [][][][][]complex64, index int) [][][][][]complex64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex64N removes n elements from list.
func Complex64N(list []complex64, index, n int) []complex64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex64ND2 removes n elements from list.
func Complex64ND2(list [][]complex64, index, n int) [][]complex64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex64ND3 removes n elements from list.
func Complex64ND3(list [][][]complex64, index, n int) [][][]complex64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex64ND4 removes n elements from list.
func Complex64ND4(list [][][][]complex64, index, n int) [][][][]complex64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex64ND5 removes n elements from list.
func Complex64ND5(list [][][][][]complex64, index, n int) [][][][][]complex64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex128 removes an element from list.
func Complex128(list []complex128, index int) []complex128 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex128D2 removes an element from list.
func Complex128D2(list [][]complex128, index int) [][]complex128 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex128D3 removes an element from list.
func Complex128D3(list [][][]complex128, index int) [][][]complex128 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex128D4 removes an element from list.
func Complex128D4(list [][][][]complex128, index int) [][][][]complex128 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex128D5 removes an element from list.
func Complex128D5(list [][][][][]complex128, index int) [][][][][]complex128 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Complex128N removes n elements from list.
func Complex128N(list []complex128, index, n int) []complex128 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex128ND2 removes n elements from list.
func Complex128ND2(list [][]complex128, index, n int) [][]complex128 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex128ND3 removes n elements from list.
func Complex128ND3(list [][][]complex128, index, n int) [][][]complex128 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex128ND4 removes n elements from list.
func Complex128ND4(list [][][][]complex128, index, n int) [][][][]complex128 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Complex128ND5 removes n elements from list.
func Complex128ND5(list [][][][][]complex128, index, n int) [][][][][]complex128 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Error removes an element from list.
func Error(list []error, index int) []error {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ErrorD2 removes an element from list.
func ErrorD2(list [][]error, index int) [][]error {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ErrorD3 removes an element from list.
func ErrorD3(list [][][]error, index int) [][][]error {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ErrorD4 removes an element from list.
func ErrorD4(list [][][][]error, index int) [][][][]error {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ErrorD5 removes an element from list.
func ErrorD5(list [][][][][]error, index int) [][][][][]error {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// ErrorN removes n elements from list.
func ErrorN(list []error, index, n int) []error {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// ErrorND2 removes n elements from list.
func ErrorND2(list [][]error, index, n int) [][]error {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// ErrorND3 removes n elements from list.
func ErrorND3(list [][][]error, index, n int) [][][]error {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// ErrorND4 removes n elements from list.
func ErrorND4(list [][][][]error, index, n int) [][][][]error {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// ErrorND5 removes n elements from list.
func ErrorND5(list [][][][][]error, index, n int) [][][][][]error {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float32 removes an element from list.
func Float32(list []float32, index int) []float32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float32D2 removes an element from list.
func Float32D2(list [][]float32, index int) [][]float32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float32D3 removes an element from list.
func Float32D3(list [][][]float32, index int) [][][]float32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float32D4 removes an element from list.
func Float32D4(list [][][][]float32, index int) [][][][]float32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float32D5 removes an element from list.
func Float32D5(list [][][][][]float32, index int) [][][][][]float32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float32N removes n elements from list.
func Float32N(list []float32, index, n int) []float32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float32ND2 removes n elements from list.
func Float32ND2(list [][]float32, index, n int) [][]float32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float32ND3 removes n elements from list.
func Float32ND3(list [][][]float32, index, n int) [][][]float32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float32ND4 removes n elements from list.
func Float32ND4(list [][][][]float32, index, n int) [][][][]float32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float32ND5 removes n elements from list.
func Float32ND5(list [][][][][]float32, index, n int) [][][][][]float32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float64 removes an element from list.
func Float64(list []float64, index int) []float64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float64D2 removes an element from list.
func Float64D2(list [][]float64, index int) [][]float64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float64D3 removes an element from list.
func Float64D3(list [][][]float64, index int) [][][]float64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float64D4 removes an element from list.
func Float64D4(list [][][][]float64, index int) [][][][]float64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float64D5 removes an element from list.
func Float64D5(list [][][][][]float64, index int) [][][][][]float64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Float64N removes n elements from list.
func Float64N(list []float64, index, n int) []float64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float64ND2 removes n elements from list.
func Float64ND2(list [][]float64, index, n int) [][]float64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float64ND3 removes n elements from list.
func Float64ND3(list [][][]float64, index, n int) [][][]float64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float64ND4 removes n elements from list.
func Float64ND4(list [][][][]float64, index, n int) [][][][]float64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Float64ND5 removes n elements from list.
func Float64ND5(list [][][][][]float64, index, n int) [][][][][]float64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int removes an element from list.
func Int(list []int, index int) []int {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// IntD2 removes an element from list.
func IntD2(list [][]int, index int) [][]int {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// IntD3 removes an element from list.
func IntD3(list [][][]int, index int) [][][]int {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// IntD4 removes an element from list.
func IntD4(list [][][][]int, index int) [][][][]int {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// IntD5 removes an element from list.
func IntD5(list [][][][][]int, index int) [][][][][]int {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// IntN removes n elements from list.
func IntN(list []int, index, n int) []int {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// IntND2 removes n elements from list.
func IntND2(list [][]int, index, n int) [][]int {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// IntND3 removes n elements from list.
func IntND3(list [][][]int, index, n int) [][][]int {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// IntND4 removes n elements from list.
func IntND4(list [][][][]int, index, n int) [][][][]int {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// IntND5 removes n elements from list.
func IntND5(list [][][][][]int, index, n int) [][][][][]int {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int8 removes an element from list.
func Int8(list []int8, index int) []int8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int8D2 removes an element from list.
func Int8D2(list [][]int8, index int) [][]int8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int8D3 removes an element from list.
func Int8D3(list [][][]int8, index int) [][][]int8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int8D4 removes an element from list.
func Int8D4(list [][][][]int8, index int) [][][][]int8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int8D5 removes an element from list.
func Int8D5(list [][][][][]int8, index int) [][][][][]int8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int8N removes n elements from list.
func Int8N(list []int8, index, n int) []int8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int8ND2 removes n elements from list.
func Int8ND2(list [][]int8, index, n int) [][]int8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int8ND3 removes n elements from list.
func Int8ND3(list [][][]int8, index, n int) [][][]int8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int8ND4 removes n elements from list.
func Int8ND4(list [][][][]int8, index, n int) [][][][]int8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int8ND5 removes n elements from list.
func Int8ND5(list [][][][][]int8, index, n int) [][][][][]int8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int16 removes an element from list.
func Int16(list []int16, index int) []int16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int16D2 removes an element from list.
func Int16D2(list [][]int16, index int) [][]int16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int16D3 removes an element from list.
func Int16D3(list [][][]int16, index int) [][][]int16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int16D4 removes an element from list.
func Int16D4(list [][][][]int16, index int) [][][][]int16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int16D5 removes an element from list.
func Int16D5(list [][][][][]int16, index int) [][][][][]int16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int16N removes n elements from list.
func Int16N(list []int16, index, n int) []int16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int16ND2 removes n elements from list.
func Int16ND2(list [][]int16, index, n int) [][]int16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int16ND3 removes n elements from list.
func Int16ND3(list [][][]int16, index, n int) [][][]int16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int16ND4 removes n elements from list.
func Int16ND4(list [][][][]int16, index, n int) [][][][]int16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int16ND5 removes n elements from list.
func Int16ND5(list [][][][][]int16, index, n int) [][][][][]int16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int32 removes an element from list.
func Int32(list []int32, index int) []int32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int32D2 removes an element from list.
func Int32D2(list [][]int32, index int) [][]int32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int32D3 removes an element from list.
func Int32D3(list [][][]int32, index int) [][][]int32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int32D4 removes an element from list.
func Int32D4(list [][][][]int32, index int) [][][][]int32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int32D5 removes an element from list.
func Int32D5(list [][][][][]int32, index int) [][][][][]int32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int32N removes n elements from list.
func Int32N(list []int32, index, n int) []int32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int32ND2 removes n elements from list.
func Int32ND2(list [][]int32, index, n int) [][]int32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int32ND3 removes n elements from list.
func Int32ND3(list [][][]int32, index, n int) [][][]int32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int32ND4 removes n elements from list.
func Int32ND4(list [][][][]int32, index, n int) [][][][]int32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int32ND5 removes n elements from list.
func Int32ND5(list [][][][][]int32, index, n int) [][][][][]int32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int64 removes an element from list.
func Int64(list []int64, index int) []int64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int64D2 removes an element from list.
func Int64D2(list [][]int64, index int) [][]int64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int64D3 removes an element from list.
func Int64D3(list [][][]int64, index int) [][][]int64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int64D4 removes an element from list.
func Int64D4(list [][][][]int64, index int) [][][][]int64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int64D5 removes an element from list.
func Int64D5(list [][][][][]int64, index int) [][][][][]int64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// Int64N removes n elements from list.
func Int64N(list []int64, index, n int) []int64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int64ND2 removes n elements from list.
func Int64ND2(list [][]int64, index, n int) [][]int64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int64ND3 removes n elements from list.
func Int64ND3(list [][][]int64, index, n int) [][][]int64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int64ND4 removes n elements from list.
func Int64ND4(list [][][][]int64, index, n int) [][][][]int64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Int64ND5 removes n elements from list.
func Int64ND5(list [][][][][]int64, index, n int) [][][][][]int64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Interface removes an element from list.
func Interface(list []interface{}, index int) []interface{} {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// InterfaceD2 removes an element from list.
func InterfaceD2(list [][]interface{}, index int) [][]interface{} {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// InterfaceD3 removes an element from list.
func InterfaceD3(list [][][]interface{}, index int) [][][]interface{} {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// InterfaceD4 removes an element from list.
func InterfaceD4(list [][][][]interface{}, index int) [][][][]interface{} {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// InterfaceD5 removes an element from list.
func InterfaceD5(list [][][][][]interface{}, index int) [][][][][]interface{} {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// InterfaceN removes n elements from list.
func InterfaceN(list []interface{}, index, n int) []interface{} {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// InterfaceND2 removes n elements from list.
func InterfaceND2(list [][]interface{}, index, n int) [][]interface{} {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// InterfaceND3 removes n elements from list.
func InterfaceND3(list [][][]interface{}, index, n int) [][][]interface{} {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// InterfaceND4 removes n elements from list.
func InterfaceND4(list [][][][]interface{}, index, n int) [][][][]interface{} {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// InterfaceND5 removes n elements from list.
func InterfaceND5(list [][][][][]interface{}, index, n int) [][][][][]interface{} {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Pointer removes an element from list.
func Pointer(list []unsafe.Pointer, index int) []unsafe.Pointer {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// PointerD2 removes an element from list.
func PointerD2(list [][]unsafe.Pointer, index int) [][]unsafe.Pointer {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// PointerD3 removes an element from list.
func PointerD3(list [][][]unsafe.Pointer, index int) [][][]unsafe.Pointer {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// PointerD4 removes an element from list.
func PointerD4(list [][][][]unsafe.Pointer, index int) [][][][]unsafe.Pointer {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// PointerD5 removes an element from list.
func PointerD5(list [][][][][]unsafe.Pointer, index int) [][][][][]unsafe.Pointer {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// PointerN removes n elements from list.
func PointerN(list []unsafe.Pointer, index, n int) []unsafe.Pointer {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// PointerND2 removes n elements from list.
func PointerND2(list [][]unsafe.Pointer, index, n int) [][]unsafe.Pointer {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// PointerND3 removes n elements from list.
func PointerND3(list [][][]unsafe.Pointer, index, n int) [][][]unsafe.Pointer {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// PointerND4 removes n elements from list.
func PointerND4(list [][][][]unsafe.Pointer, index, n int) [][][][]unsafe.Pointer {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// PointerND5 removes n elements from list.
func PointerND5(list [][][][][]unsafe.Pointer, index, n int) [][][][][]unsafe.Pointer {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// Rune removes an element from list.
func Rune(list []rune, index int) []rune {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// RuneD2 removes an element from list.
func RuneD2(list [][]rune, index int) [][]rune {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// RuneD3 removes an element from list.
func RuneD3(list [][][]rune, index int) [][][]rune {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// RuneD4 removes an element from list.
func RuneD4(list [][][][]rune, index int) [][][][]rune {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// RuneD5 removes an element from list.
func RuneD5(list [][][][][]rune, index int) [][][][][]rune {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// RuneN removes n elements from list.
func RuneN(list []rune, index, n int) []rune {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// RuneND2 removes n elements from list.
func RuneND2(list [][]rune, index, n int) [][]rune {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// RuneND3 removes n elements from list.
func RuneND3(list [][][]rune, index, n int) [][][]rune {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// RuneND4 removes n elements from list.
func RuneND4(list [][][][]rune, index, n int) [][][][]rune {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// RuneND5 removes n elements from list.
func RuneND5(list [][][][][]rune, index, n int) [][][][][]rune {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// String removes an element from list.
func String(list []string, index int) []string {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// StringD2 removes an element from list.
func StringD2(list [][]string, index int) [][]string {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// StringD3 removes an element from list.
func StringD3(list [][][]string, index int) [][][]string {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// StringD4 removes an element from list.
func StringD4(list [][][][]string, index int) [][][][]string {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// StringD5 removes an element from list.
func StringD5(list [][][][][]string, index int) [][][][][]string {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// StringN removes n elements from list.
func StringN(list []string, index, n int) []string {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// StringND2 removes n elements from list.
func StringND2(list [][]string, index, n int) [][]string {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// StringND3 removes n elements from list.
func StringND3(list [][][]string, index, n int) [][][]string {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// StringND4 removes n elements from list.
func StringND4(list [][][][]string, index, n int) [][][][]string {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// StringND5 removes n elements from list.
func StringND5(list [][][][][]string, index, n int) [][][][][]string {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt removes an element from list.
func UInt(list []uint, index int) []uint {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UIntD2 removes an element from list.
func UIntD2(list [][]uint, index int) [][]uint {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UIntD3 removes an element from list.
func UIntD3(list [][][]uint, index int) [][][]uint {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UIntD4 removes an element from list.
func UIntD4(list [][][][]uint, index int) [][][][]uint {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UIntD5 removes an element from list.
func UIntD5(list [][][][][]uint, index int) [][][][][]uint {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UIntN removes n elements from list.
func UIntN(list []uint, index, n int) []uint {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UIntND2 removes n elements from list.
func UIntND2(list [][]uint, index, n int) [][]uint {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UIntND3 removes n elements from list.
func UIntND3(list [][][]uint, index, n int) [][][]uint {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UIntND4 removes n elements from list.
func UIntND4(list [][][][]uint, index, n int) [][][][]uint {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UIntND5 removes n elements from list.
func UIntND5(list [][][][][]uint, index, n int) [][][][][]uint {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt8 removes an element from list.
func UInt8(list []uint8, index int) []uint8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt8D2 removes an element from list.
func UInt8D2(list [][]uint8, index int) [][]uint8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt8D3 removes an element from list.
func UInt8D3(list [][][]uint8, index int) [][][]uint8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt8D4 removes an element from list.
func UInt8D4(list [][][][]uint8, index int) [][][][]uint8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt8D5 removes an element from list.
func UInt8D5(list [][][][][]uint8, index int) [][][][][]uint8 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt8N removes n elements from list.
func UInt8N(list []uint8, index, n int) []uint8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt8ND2 removes n elements from list.
func UInt8ND2(list [][]uint8, index, n int) [][]uint8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt8ND3 removes n elements from list.
func UInt8ND3(list [][][]uint8, index, n int) [][][]uint8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt8ND4 removes n elements from list.
func UInt8ND4(list [][][][]uint8, index, n int) [][][][]uint8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt8ND5 removes n elements from list.
func UInt8ND5(list [][][][][]uint8, index, n int) [][][][][]uint8 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt16 removes an element from list.
func UInt16(list []uint16, index int) []uint16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt16D2 removes an element from list.
func UInt16D2(list [][]uint16, index int) [][]uint16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt16D3 removes an element from list.
func UInt16D3(list [][][]uint16, index int) [][][]uint16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt16D4 removes an element from list.
func UInt16D4(list [][][][]uint16, index int) [][][][]uint16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt16D5 removes an element from list.
func UInt16D5(list [][][][][]uint16, index int) [][][][][]uint16 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt16N removes n elements from list.
func UInt16N(list []uint16, index, n int) []uint16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt16ND2 removes n elements from list.
func UInt16ND2(list [][]uint16, index, n int) [][]uint16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt16ND3 removes n elements from list.
func UInt16ND3(list [][][]uint16, index, n int) [][][]uint16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt16ND4 removes n elements from list.
func UInt16ND4(list [][][][]uint16, index, n int) [][][][]uint16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt16ND5 removes n elements from list.
func UInt16ND5(list [][][][][]uint16, index, n int) [][][][][]uint16 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt32 removes an element from list.
func UInt32(list []uint32, index int) []uint32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt32D2 removes an element from list.
func UInt32D2(list [][]uint32, index int) [][]uint32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt32D3 removes an element from list.
func UInt32D3(list [][][]uint32, index int) [][][]uint32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt32D4 removes an element from list.
func UInt32D4(list [][][][]uint32, index int) [][][][]uint32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt32D5 removes an element from list.
func UInt32D5(list [][][][][]uint32, index int) [][][][][]uint32 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt32N removes n elements from list.
func UInt32N(list []uint32, index, n int) []uint32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt32ND2 removes n elements from list.
func UInt32ND2(list [][]uint32, index, n int) [][]uint32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt32ND3 removes n elements from list.
func UInt32ND3(list [][][]uint32, index, n int) [][][]uint32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt32ND4 removes n elements from list.
func UInt32ND4(list [][][][]uint32, index, n int) [][][][]uint32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt32ND5 removes n elements from list.
func UInt32ND5(list [][][][][]uint32, index, n int) [][][][][]uint32 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt64 removes an element from list.
func UInt64(list []uint64, index int) []uint64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt64D2 removes an element from list.
func UInt64D2(list [][]uint64, index int) [][]uint64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt64D3 removes an element from list.
func UInt64D3(list [][][]uint64, index int) [][][]uint64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt64D4 removes an element from list.
func UInt64D4(list [][][][]uint64, index int) [][][][]uint64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt64D5 removes an element from list.
func UInt64D5(list [][][][][]uint64, index int) [][][][][]uint64 {
	copy(list[index:], list[index+1:])
	return list[:len(list)-1]
}

// UInt64N removes n elements from list.
func UInt64N(list []uint64, index, n int) []uint64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt64ND2 removes n elements from list.
func UInt64ND2(list [][]uint64, index, n int) [][]uint64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt64ND3 removes n elements from list.
func UInt64ND3(list [][][]uint64, index, n int) [][][]uint64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt64ND4 removes n elements from list.
func UInt64ND4(list [][][][]uint64, index, n int) [][][][]uint64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}

// UInt64ND5 removes n elements from list.
func UInt64ND5(list [][][][][]uint64, index, n int) [][][][][]uint64 {
	copy(list[index:], list[index+n:])
	return list[:len(list)-n]
}
