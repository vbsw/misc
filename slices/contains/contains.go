/*
 *       Copyright 2018, 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package contains provides the "contains" function for slices of basic types.
package contains

import (
	"unsafe"
)

// Bool returns true, if list contains value.
func Bool(list []bool, value bool) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// BoolD2 returns true, if list contains value.
func BoolD2(list [][]bool, value bool) bool {
	for _, listValue := range list {
		if Bool(listValue, value) {
			return true
		}
	}
	return false
}

// BoolD3 returns true, if list contains value.
func BoolD3(list [][][]bool, value bool) bool {
	for _, listValue := range list {
		if BoolD2(listValue, value) {
			return true
		}
	}
	return false
}

// BoolD4 returns true, if list contains value.
func BoolD4(list [][][][]bool, value bool) bool {
	for _, listValue := range list {
		if BoolD3(listValue, value) {
			return true
		}
	}
	return false
}

// BoolD5 returns true, if list contains value.
func BoolD5(list [][][][][]bool, value bool) bool {
	for _, listValue := range list {
		if BoolD4(listValue, value) {
			return true
		}
	}
	return false
}

// Byte returns true, if list contains value.
func Byte(list []byte, value byte) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// ByteD2 returns true, if list contains value.
func ByteD2(list [][]byte, value byte) bool {
	for _, listValue := range list {
		if Byte(listValue, value) {
			return true
		}
	}
	return false
}

// ByteD3 returns true, if list contains value.
func ByteD3(list [][][]byte, value byte) bool {
	for _, listValue := range list {
		if ByteD2(listValue, value) {
			return true
		}
	}
	return false
}

// ByteD4 returns true, if list contains value.
func ByteD4(list [][][][]byte, value byte) bool {
	for _, listValue := range list {
		if ByteD3(listValue, value) {
			return true
		}
	}
	return false
}

// ByteD5 returns true, if list contains value.
func ByteD5(list [][][][][]byte, value byte) bool {
	for _, listValue := range list {
		if ByteD4(listValue, value) {
			return true
		}
	}
	return false
}

// Complex64 returns true, if list contains value.
func Complex64(list []complex64, value complex64) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Complex64D2 returns true, if list contains value.
func Complex64D2(list [][]complex64, value complex64) bool {
	for _, listValue := range list {
		if Complex64(listValue, value) {
			return true
		}
	}
	return false
}

// Complex64D3 returns true, if list contains value.
func Complex64D3(list [][][]complex64, value complex64) bool {
	for _, listValue := range list {
		if Complex64D2(listValue, value) {
			return true
		}
	}
	return false
}

// Complex64D4 returns true, if list contains value.
func Complex64D4(list [][][][]complex64, value complex64) bool {
	for _, listValue := range list {
		if Complex64D3(listValue, value) {
			return true
		}
	}
	return false
}

// Complex64D5 returns true, if list contains value.
func Complex64D5(list [][][][][]complex64, value complex64) bool {
	for _, listValue := range list {
		if Complex64D4(listValue, value) {
			return true
		}
	}
	return false
}

// Complex128 returns true, if list contains value.
func Complex128(list []complex128, value complex128) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Complex128D2 returns true, if list contains value.
func Complex128D2(list [][]complex128, value complex128) bool {
	for _, listValue := range list {
		if Complex128(listValue, value) {
			return true
		}
	}
	return false
}

// Complex128D3 returns true, if list contains value.
func Complex128D3(list [][][]complex128, value complex128) bool {
	for _, listValue := range list {
		if Complex128D2(listValue, value) {
			return true
		}
	}
	return false
}

// Complex128D4 returns true, if list contains value.
func Complex128D4(list [][][][]complex128, value complex128) bool {
	for _, listValue := range list {
		if Complex128D3(listValue, value) {
			return true
		}
	}
	return false
}

// Complex128D5 returns true, if list contains value.
func Complex128D5(list [][][][][]complex128, value complex128) bool {
	for _, listValue := range list {
		if Complex128D4(listValue, value) {
			return true
		}
	}
	return false
}

// Error returns true, if list contains value.
func Error(list []error, value error) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// ErrorD2 returns true, if list contains value.
func ErrorD2(list [][]error, value error) bool {
	for _, listValue := range list {
		if Error(listValue, value) {
			return true
		}
	}
	return false
}

// ErrorD3 returns true, if list contains value.
func ErrorD3(list [][][]error, value error) bool {
	for _, listValue := range list {
		if ErrorD2(listValue, value) {
			return true
		}
	}
	return false
}

// ErrorD4 returns true, if list contains value.
func ErrorD4(list [][][][]error, value error) bool {
	for _, listValue := range list {
		if ErrorD3(listValue, value) {
			return true
		}
	}
	return false
}

// ErrorD5 returns true, if list contains value.
func ErrorD5(list [][][][][]error, value error) bool {
	for _, listValue := range list {
		if ErrorD4(listValue, value) {
			return true
		}
	}
	return false
}

// Float32 returns true, if list contains value.
func Float32(list []float32, value float32) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Float32D2 returns true, if list contains value.
func Float32D2(list [][]float32, value float32) bool {
	for _, listValue := range list {
		if Float32(listValue, value) {
			return true
		}
	}
	return false
}

// Float32D3 returns true, if list contains value.
func Float32D3(list [][][]float32, value float32) bool {
	for _, listValue := range list {
		if Float32D2(listValue, value) {
			return true
		}
	}
	return false
}

// Float32D4 returns true, if list contains value.
func Float32D4(list [][][][]float32, value float32) bool {
	for _, listValue := range list {
		if Float32D3(listValue, value) {
			return true
		}
	}
	return false
}

// Float32D5 returns true, if list contains value.
func Float32D5(list [][][][][]float32, value float32) bool {
	for _, listValue := range list {
		if Float32D4(listValue, value) {
			return true
		}
	}
	return false
}

// Float64 returns true, if list contains value.
func Float64(list []float64, value float64) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Float64D2 returns true, if list contains value.
func Float64D2(list [][]float64, value float64) bool {
	for _, listValue := range list {
		if Float64(listValue, value) {
			return true
		}
	}
	return false
}

// Float64D3 returns true, if list contains value.
func Float64D3(list [][][]float64, value float64) bool {
	for _, listValue := range list {
		if Float64D2(listValue, value) {
			return true
		}
	}
	return false
}

// Float64D4 returns true, if list contains value.
func Float64D4(list [][][][]float64, value float64) bool {
	for _, listValue := range list {
		if Float64D3(listValue, value) {
			return true
		}
	}
	return false
}

// Float64D5 returns true, if list contains value.
func Float64D5(list [][][][][]float64, value float64) bool {
	for _, listValue := range list {
		if Float64D4(listValue, value) {
			return true
		}
	}
	return false
}

// Int returns true, if list contains value.
func Int(list []int, value int) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// IntD2 returns true, if list contains value.
func IntD2(list [][]int, value int) bool {
	for _, listValue := range list {
		if Int(listValue, value) {
			return true
		}
	}
	return false
}

// IntD3 returns true, if list contains value.
func IntD3(list [][][]int, value int) bool {
	for _, listValue := range list {
		if IntD2(listValue, value) {
			return true
		}
	}
	return false
}

// IntD4 returns true, if list contains value.
func IntD4(list [][][][]int, value int) bool {
	for _, listValue := range list {
		if IntD3(listValue, value) {
			return true
		}
	}
	return false
}

// IntD5 returns true, if list contains value.
func IntD5(list [][][][][]int, value int) bool {
	for _, listValue := range list {
		if IntD4(listValue, value) {
			return true
		}
	}
	return false
}

// Int8 returns true, if list contains value.
func Int8(list []int8, value int8) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Int8D2 returns true, if list contains value.
func Int8D2(list [][]int8, value int8) bool {
	for _, listValue := range list {
		if Int8(listValue, value) {
			return true
		}
	}
	return false
}

// Int8D3 returns true, if list contains value.
func Int8D3(list [][][]int8, value int8) bool {
	for _, listValue := range list {
		if Int8D2(listValue, value) {
			return true
		}
	}
	return false
}

// Int8D4 returns true, if list contains value.
func Int8D4(list [][][][]int8, value int8) bool {
	for _, listValue := range list {
		if Int8D3(listValue, value) {
			return true
		}
	}
	return false
}

// Int8D5 returns true, if list contains value.
func Int8D5(list [][][][][]int8, value int8) bool {
	for _, listValue := range list {
		if Int8D4(listValue, value) {
			return true
		}
	}
	return false
}

// Int16 returns true, if list contains value.
func Int16(list []int16, value int16) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Int16D2 returns true, if list contains value.
func Int16D2(list [][]int16, value int16) bool {
	for _, listValue := range list {
		if Int16(listValue, value) {
			return true
		}
	}
	return false
}

// Int16D3 returns true, if list contains value.
func Int16D3(list [][][]int16, value int16) bool {
	for _, listValue := range list {
		if Int16D2(listValue, value) {
			return true
		}
	}
	return false
}

// Int16D4 returns true, if list contains value.
func Int16D4(list [][][][]int16, value int16) bool {
	for _, listValue := range list {
		if Int16D3(listValue, value) {
			return true
		}
	}
	return false
}

// Int16D5 returns true, if list contains value.
func Int16D5(list [][][][][]int16, value int16) bool {
	for _, listValue := range list {
		if Int16D4(listValue, value) {
			return true
		}
	}
	return false
}

// Int32 returns true, if list contains value.
func Int32(list []int32, value int32) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Int32D2 returns true, if list contains value.
func Int32D2(list [][]int32, value int32) bool {
	for _, listValue := range list {
		if Int32(listValue, value) {
			return true
		}
	}
	return false
}

// Int32D3 returns true, if list contains value.
func Int32D3(list [][][]int32, value int32) bool {
	for _, listValue := range list {
		if Int32D2(listValue, value) {
			return true
		}
	}
	return false
}

// Int32D4 returns true, if list contains value.
func Int32D4(list [][][][]int32, value int32) bool {
	for _, listValue := range list {
		if Int32D3(listValue, value) {
			return true
		}
	}
	return false
}

// Int32D5 returns true, if list contains value.
func Int32D5(list [][][][][]int32, value int32) bool {
	for _, listValue := range list {
		if Int32D4(listValue, value) {
			return true
		}
	}
	return false
}

// Int64 returns true, if list contains value.
func Int64(list []int64, value int64) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// Int64D2 returns true, if list contains value.
func Int64D2(list [][]int64, value int64) bool {
	for _, listValue := range list {
		if Int64(listValue, value) {
			return true
		}
	}
	return false
}

// Int64D3 returns true, if list contains value.
func Int64D3(list [][][]int64, value int64) bool {
	for _, listValue := range list {
		if Int64D2(listValue, value) {
			return true
		}
	}
	return false
}

// Int64D4 returns true, if list contains value.
func Int64D4(list [][][][]int64, value int64) bool {
	for _, listValue := range list {
		if Int64D3(listValue, value) {
			return true
		}
	}
	return false
}

// Int64D5 returns true, if list contains value.
func Int64D5(list [][][][][]int64, value int64) bool {
	for _, listValue := range list {
		if Int64D4(listValue, value) {
			return true
		}
	}
	return false
}

// Interface returns true, if list contains value.
func Interface(list []interface{}, value interface{}) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// InterfaceD2 returns true, if list contains value.
func InterfaceD2(list [][]interface{}, value interface{}) bool {
	for _, listValue := range list {
		if Interface(listValue, value) {
			return true
		}
	}
	return false
}

// InterfaceD3 returns true, if list contains value.
func InterfaceD3(list [][][]interface{}, value interface{}) bool {
	for _, listValue := range list {
		if InterfaceD2(listValue, value) {
			return true
		}
	}
	return false
}

// InterfaceD4 returns true, if list contains value.
func InterfaceD4(list [][][][]interface{}, value interface{}) bool {
	for _, listValue := range list {
		if InterfaceD3(listValue, value) {
			return true
		}
	}
	return false
}

// InterfaceD5 returns true, if list contains value.
func InterfaceD5(list [][][][][]interface{}, value interface{}) bool {
	for _, listValue := range list {
		if InterfaceD4(listValue, value) {
			return true
		}
	}
	return false
}

// Pointer returns true, if list contains value.
func Pointer(list []unsafe.Pointer, value unsafe.Pointer) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// PointerD2 returns true, if list contains value.
func PointerD2(list [][]unsafe.Pointer, value unsafe.Pointer) bool {
	for _, listValue := range list {
		if Pointer(listValue, value) {
			return true
		}
	}
	return false
}

// PointerD3 returns true, if list contains value.
func PointerD3(list [][][]unsafe.Pointer, value unsafe.Pointer) bool {
	for _, listValue := range list {
		if PointerD2(listValue, value) {
			return true
		}
	}
	return false
}

// PointerD4 returns true, if list contains value.
func PointerD4(list [][][][]unsafe.Pointer, value unsafe.Pointer) bool {
	for _, listValue := range list {
		if PointerD3(listValue, value) {
			return true
		}
	}
	return false
}

// PointerD5 returns true, if list contains value.
func PointerD5(list [][][][][]unsafe.Pointer, value unsafe.Pointer) bool {
	for _, listValue := range list {
		if PointerD4(listValue, value) {
			return true
		}
	}
	return false
}

// Rune returns true, if list contains value.
func Rune(list []rune, value rune) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// RuneD2 returns true, if list contains value.
func RuneD2(list [][]rune, value rune) bool {
	for _, listValue := range list {
		if Rune(listValue, value) {
			return true
		}
	}
	return false
}

// RuneD3 returns true, if list contains value.
func RuneD3(list [][][]rune, value rune) bool {
	for _, listValue := range list {
		if RuneD2(listValue, value) {
			return true
		}
	}
	return false
}

// RuneD4 returns true, if list contains value.
func RuneD4(list [][][][]rune, value rune) bool {
	for _, listValue := range list {
		if RuneD3(listValue, value) {
			return true
		}
	}
	return false
}

// RuneD5 returns true, if list contains value.
func RuneD5(list [][][][][]rune, value rune) bool {
	for _, listValue := range list {
		if RuneD4(listValue, value) {
			return true
		}
	}
	return false
}

// String returns true, if list contains value.
func String(list []string, value string) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// StringD2 returns true, if list contains value.
func StringD2(list [][]string, value string) bool {
	for _, listValue := range list {
		if String(listValue, value) {
			return true
		}
	}
	return false
}

// StringD3 returns true, if list contains value.
func StringD3(list [][][]string, value string) bool {
	for _, listValue := range list {
		if StringD2(listValue, value) {
			return true
		}
	}
	return false
}

// StringD4 returns true, if list contains value.
func StringD4(list [][][][]string, value string) bool {
	for _, listValue := range list {
		if StringD3(listValue, value) {
			return true
		}
	}
	return false
}

// StringD5 returns true, if list contains value.
func StringD5(list [][][][][]string, value string) bool {
	for _, listValue := range list {
		if StringD4(listValue, value) {
			return true
		}
	}
	return false
}

// UInt returns true, if list contains value.
func UInt(list []uint, value uint) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UIntD2 returns true, if list contains value.
func UIntD2(list [][]uint, value uint) bool {
	for _, listValue := range list {
		if UInt(listValue, value) {
			return true
		}
	}
	return false
}

// UIntD3 returns true, if list contains value.
func UIntD3(list [][][]uint, value uint) bool {
	for _, listValue := range list {
		if UIntD2(listValue, value) {
			return true
		}
	}
	return false
}

// UIntD4 returns true, if list contains value.
func UIntD4(list [][][][]uint, value uint) bool {
	for _, listValue := range list {
		if UIntD3(listValue, value) {
			return true
		}
	}
	return false
}

// UIntD5 returns true, if list contains value.
func UIntD5(list [][][][][]uint, value uint) bool {
	for _, listValue := range list {
		if UIntD4(listValue, value) {
			return true
		}
	}
	return false
}

// UInt8 returns true, if list contains value.
func UInt8(list []uint8, value uint8) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UInt8D2 returns true, if list contains value.
func UInt8D2(list [][]uint8, value uint8) bool {
	for _, listValue := range list {
		if UInt8(listValue, value) {
			return true
		}
	}
	return false
}

// UInt8D3 returns true, if list contains value.
func UInt8D3(list [][][]uint8, value uint8) bool {
	for _, listValue := range list {
		if UInt8D2(listValue, value) {
			return true
		}
	}
	return false
}

// UInt8D4 returns true, if list contains value.
func UInt8D4(list [][][][]uint8, value uint8) bool {
	for _, listValue := range list {
		if UInt8D3(listValue, value) {
			return true
		}
	}
	return false
}

// UInt8D5 returns true, if list contains value.
func UInt8D5(list [][][][][]uint8, value uint8) bool {
	for _, listValue := range list {
		if UInt8D4(listValue, value) {
			return true
		}
	}
	return false
}

// UInt16 returns true, if list contains value.
func UInt16(list []uint16, value uint16) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UInt16D2 returns true, if list contains value.
func UInt16D2(list [][]uint16, value uint16) bool {
	for _, listValue := range list {
		if UInt16(listValue, value) {
			return true
		}
	}
	return false
}

// UInt16D3 returns true, if list contains value.
func UInt16D3(list [][][]uint16, value uint16) bool {
	for _, listValue := range list {
		if UInt16D2(listValue, value) {
			return true
		}
	}
	return false
}

// UInt16D4 returns true, if list contains value.
func UInt16D4(list [][][][]uint16, value uint16) bool {
	for _, listValue := range list {
		if UInt16D3(listValue, value) {
			return true
		}
	}
	return false
}

// UInt16D5 returns true, if list contains value.
func UInt16D5(list [][][][][]uint16, value uint16) bool {
	for _, listValue := range list {
		if UInt16D4(listValue, value) {
			return true
		}
	}
	return false
}

// UInt32 returns true, if list contains value.
func UInt32(list []uint32, value uint32) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UInt32D2 returns true, if list contains value.
func UInt32D2(list [][]uint32, value uint32) bool {
	for _, listValue := range list {
		if UInt32(listValue, value) {
			return true
		}
	}
	return false
}

// UInt32D3 returns true, if list contains value.
func UInt32D3(list [][][]uint32, value uint32) bool {
	for _, listValue := range list {
		if UInt32D2(listValue, value) {
			return true
		}
	}
	return false
}

// UInt32D4 returns true, if list contains value.
func UInt32D4(list [][][][]uint32, value uint32) bool {
	for _, listValue := range list {
		if UInt32D3(listValue, value) {
			return true
		}
	}
	return false
}

// UInt32D5 returns true, if list contains value.
func UInt32D5(list [][][][][]uint32, value uint32) bool {
	for _, listValue := range list {
		if UInt32D4(listValue, value) {
			return true
		}
	}
	return false
}

// UInt64 returns true, if list contains value.
func UInt64(list []uint64, value uint64) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}
	return false
}

// UInt64D2 returns true, if list contains value.
func UInt64D2(list [][]uint64, value uint64) bool {
	for _, listValue := range list {
		if UInt64(listValue, value) {
			return true
		}
	}
	return false
}

// UInt64D3 returns true, if list contains value.
func UInt64D3(list [][][]uint64, value uint64) bool {
	for _, listValue := range list {
		if UInt64D2(listValue, value) {
			return true
		}
	}
	return false
}

// UInt64D4 returns true, if list contains value.
func UInt64D4(list [][][][]uint64, value uint64) bool {
	for _, listValue := range list {
		if UInt64D3(listValue, value) {
			return true
		}
	}
	return false
}

// UInt64D5 returns true, if list contains value.
func UInt64D5(list [][][][][]uint64, value uint64) bool {
	for _, listValue := range list {
		if UInt64D4(listValue, value) {
			return true
		}
	}
	return false
}
