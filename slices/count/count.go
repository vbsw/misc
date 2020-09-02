/*
 *       Copyright 2019, 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package count provides the "count" function for slices of basic types.
package count

import (
	"unsafe"
)

// Bool returns count of value in list.
func Bool(list []bool, value bool) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// BoolD2 returns count of value in list.
func BoolD2(list [][]bool, value bool) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Bool(listValue, value)
	}
	return valueCount
}

// BoolD3 returns count of value in list.
func BoolD3(list [][][]bool, value bool) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += BoolD2(listValue, value)
	}
	return valueCount
}

// BoolD4 returns count of value in list.
func BoolD4(list [][][][]bool, value bool) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += BoolD3(listValue, value)
	}
	return valueCount
}

// BoolD5 returns count of value in list.
func BoolD5(list [][][][][]bool, value bool) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += BoolD4(listValue, value)
	}
	return valueCount
}

// Byte returns count of value in list.
func Byte(list []byte, value byte) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// ByteD2 returns count of value in list.
func ByteD2(list [][]byte, value byte) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Byte(listValue, value)
	}
	return valueCount
}

// ByteD3 returns count of value in list.
func ByteD3(list [][][]byte, value byte) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += ByteD2(listValue, value)
	}
	return valueCount
}

// ByteD4 returns count of value in list.
func ByteD4(list [][][][]byte, value byte) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += ByteD3(listValue, value)
	}
	return valueCount
}

// ByteD5 returns count of value in list.
func ByteD5(list [][][][][]byte, value byte) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += ByteD4(listValue, value)
	}
	return valueCount
}

// Complex64 returns count of value in list.
func Complex64(list []complex64, value complex64) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// Complex64D2 returns count of value in list.
func Complex64D2(list [][]complex64, value complex64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Complex64(listValue, value)
	}
	return valueCount
}

// Complex64D3 returns count of value in list.
func Complex64D3(list [][][]complex64, value complex64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Complex64D2(listValue, value)
	}
	return valueCount
}

// Complex64D4 returns count of value in list.
func Complex64D4(list [][][][]complex64, value complex64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Complex64D3(listValue, value)
	}
	return valueCount
}

// Complex64D5 returns count of value in list.
func Complex64D5(list [][][][][]complex64, value complex64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Complex64D4(listValue, value)
	}
	return valueCount
}

// Complex128 returns count of value in list.
func Complex128(list []complex128, value complex128) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// Complex128D2 returns count of value in list.
func Complex128D2(list [][]complex128, value complex128) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Complex128(listValue, value)
	}
	return valueCount
}

// Complex128D3 returns count of value in list.
func Complex128D3(list [][][]complex128, value complex128) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Complex128D2(listValue, value)
	}
	return valueCount
}

// Complex128D4 returns count of value in list.
func Complex128D4(list [][][][]complex128, value complex128) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Complex128D3(listValue, value)
	}
	return valueCount
}

// Complex128D5 returns count of value in list.
func Complex128D5(list [][][][][]complex128, value complex128) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Complex128D4(listValue, value)
	}
	return valueCount
}

// Error returns count of value in list.
func Error(list []error, value error) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// ErrorD2 returns count of value in list.
func ErrorD2(list [][]error, value error) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Error(listValue, value)
	}
	return valueCount
}

// ErrorD3 returns count of value in list.
func ErrorD3(list [][][]error, value error) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += ErrorD2(listValue, value)
	}
	return valueCount
}

// ErrorD4 returns count of value in list.
func ErrorD4(list [][][][]error, value error) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += ErrorD3(listValue, value)
	}
	return valueCount
}

// ErrorD5 returns count of value in list.
func ErrorD5(list [][][][][]error, value error) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += ErrorD4(listValue, value)
	}
	return valueCount
}

// Float32 returns count of value in list.
func Float32(list []float32, value float32) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// Float32D2 returns count of value in list.
func Float32D2(list [][]float32, value float32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float32(listValue, value)
	}
	return valueCount
}

// Float32D3 returns count of value in list.
func Float32D3(list [][][]float32, value float32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float32D2(listValue, value)
	}
	return valueCount
}

// Float32D4 returns count of value in list.
func Float32D4(list [][][][]float32, value float32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float32D3(listValue, value)
	}
	return valueCount
}

// Float32D5 returns count of value in list.
func Float32D5(list [][][][][]float32, value float32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float32D4(listValue, value)
	}
	return valueCount
}

// Float64 returns count of value in list.
func Float64(list []float64, value float64) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// Float64D2 returns count of value in list.
func Float64D2(list [][]float64, value float64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float64(listValue, value)
	}
	return valueCount
}

// Float64D3 returns count of value in list.
func Float64D3(list [][][]float64, value float64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float64D2(listValue, value)
	}
	return valueCount
}

// Float64D4 returns count of value in list.
func Float64D4(list [][][][]float64, value float64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float64D3(listValue, value)
	}
	return valueCount
}

// Float64D5 returns count of value in list.
func Float64D5(list [][][][][]float64, value float64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Float64D4(listValue, value)
	}
	return valueCount
}

// Int returns count of value in list.
func Int(list []int, value int) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// IntD2 returns count of value in list.
func IntD2(list [][]int, value int) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int(listValue, value)
	}
	return valueCount
}

// IntD3 returns count of value in list.
func IntD3(list [][][]int, value int) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += IntD2(listValue, value)
	}
	return valueCount
}

// IntD4 returns count of value in list.
func IntD4(list [][][][]int, value int) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += IntD3(listValue, value)
	}
	return valueCount
}

// IntD5 returns count of value in list.
func IntD5(list [][][][][]int, value int) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += IntD4(listValue, value)
	}
	return valueCount
}

// Int8 returns count of value in list.
func Int8(list []int8, value int8) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// Int8D2 returns count of value in list.
func Int8D2(list [][]int8, value int8) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int8(listValue, value)
	}
	return valueCount
}

// Int8D3 returns count of value in list.
func Int8D3(list [][][]int8, value int8) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int8D2(listValue, value)
	}
	return valueCount
}

// Int8D4 returns count of value in list.
func Int8D4(list [][][][]int8, value int8) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int8D3(listValue, value)
	}
	return valueCount
}

// Int8D5 returns count of value in list.
func Int8D5(list [][][][][]int8, value int8) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int8D4(listValue, value)
	}
	return valueCount
}

// Int16 returns count of value in list.
func Int16(list []int16, value int16) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// Int16D2 returns count of value in list.
func Int16D2(list [][]int16, value int16) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int16(listValue, value)
	}
	return valueCount
}

// Int16D3 returns count of value in list.
func Int16D3(list [][][]int16, value int16) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int16D2(listValue, value)
	}
	return valueCount
}

// Int16D4 returns count of value in list.
func Int16D4(list [][][][]int16, value int16) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int16D3(listValue, value)
	}
	return valueCount
}

// Int16D5 returns count of value in list.
func Int16D5(list [][][][][]int16, value int16) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int16D4(listValue, value)
	}
	return valueCount
}

// Int32 returns count of value in list.
func Int32(list []int32, value int32) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// Int32D2 returns count of value in list.
func Int32D2(list [][]int32, value int32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int32(listValue, value)
	}
	return valueCount
}

// Int32D3 returns count of value in list.
func Int32D3(list [][][]int32, value int32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int32D2(listValue, value)
	}
	return valueCount
}

// Int32D4 returns count of value in list.
func Int32D4(list [][][][]int32, value int32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int32D3(listValue, value)
	}
	return valueCount
}

// Int32D5 returns count of value in list.
func Int32D5(list [][][][][]int32, value int32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int32D4(listValue, value)
	}
	return valueCount
}

// Int64 returns count of value in list.
func Int64(list []int64, value int64) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// Int64D2 returns count of value in list.
func Int64D2(list [][]int64, value int64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int64(listValue, value)
	}
	return valueCount
}

// Int64D3 returns count of value in list.
func Int64D3(list [][][]int64, value int64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int64D2(listValue, value)
	}
	return valueCount
}

// Int64D4 returns count of value in list.
func Int64D4(list [][][][]int64, value int64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int64D3(listValue, value)
	}
	return valueCount
}

// Int64D5 returns count of value in list.
func Int64D5(list [][][][][]int64, value int64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Int64D4(listValue, value)
	}
	return valueCount
}

// Interface returns count of value in list.
func Interface(list []interface{}, value interface{}) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// InterfaceD2 returns count of value in list.
func InterfaceD2(list [][]interface{}, value interface{}) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Interface(listValue, value)
	}
	return valueCount
}

// InterfaceD3 returns count of value in list.
func InterfaceD3(list [][][]interface{}, value interface{}) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += InterfaceD2(listValue, value)
	}
	return valueCount
}

// InterfaceD4 returns count of value in list.
func InterfaceD4(list [][][][]interface{}, value interface{}) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += InterfaceD3(listValue, value)
	}
	return valueCount
}

// InterfaceD5 returns count of value in list.
func InterfaceD5(list [][][][][]interface{}, value interface{}) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += InterfaceD4(listValue, value)
	}
	return valueCount
}

// Pointer returns count of value in list.
func Pointer(list []unsafe.Pointer, value unsafe.Pointer) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// PointerD2 returns count of value in list.
func PointerD2(list [][]unsafe.Pointer, value unsafe.Pointer) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Pointer(listValue, value)
	}
	return valueCount
}

// PointerD3 returns count of value in list.
func PointerD3(list [][][]unsafe.Pointer, value unsafe.Pointer) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += PointerD2(listValue, value)
	}
	return valueCount
}

// PointerD4 returns count of value in list.
func PointerD4(list [][][][]unsafe.Pointer, value unsafe.Pointer) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += PointerD3(listValue, value)
	}
	return valueCount
}

// PointerD5 returns count of value in list.
func PointerD5(list [][][][][]unsafe.Pointer, value unsafe.Pointer) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += PointerD4(listValue, value)
	}
	return valueCount
}

// Rune returns count of value in list.
func Rune(list []rune, value rune) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// RuneD2 returns count of value in list.
func RuneD2(list [][]rune, value rune) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += Rune(listValue, value)
	}
	return valueCount
}

// RuneD3 returns count of value in list.
func RuneD3(list [][][]rune, value rune) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += RuneD2(listValue, value)
	}
	return valueCount
}

// RuneD4 returns count of value in list.
func RuneD4(list [][][][]rune, value rune) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += RuneD3(listValue, value)
	}
	return valueCount
}

// RuneD5 returns count of value in list.
func RuneD5(list [][][][][]rune, value rune) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += RuneD4(listValue, value)
	}
	return valueCount
}

// String returns count of value in list.
func String(list []string, value string) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// StringD2 returns count of value in list.
func StringD2(list [][]string, value string) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += String(listValue, value)
	}
	return valueCount
}

// StringD3 returns count of value in list.
func StringD3(list [][][]string, value string) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += StringD2(listValue, value)
	}
	return valueCount
}

// StringD4 returns count of value in list.
func StringD4(list [][][][]string, value string) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += StringD3(listValue, value)
	}
	return valueCount
}

// StringD5 returns count of value in list.
func StringD5(list [][][][][]string, value string) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += StringD4(listValue, value)
	}
	return valueCount
}

// UInt returns count of value in list.
func UInt(list []uint, value uint) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// UIntD2 returns count of value in list.
func UIntD2(list [][]uint, value uint) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt(listValue, value)
	}
	return valueCount
}

// UIntD3 returns count of value in list.
func UIntD3(list [][][]uint, value uint) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UIntD2(listValue, value)
	}
	return valueCount
}

// UIntD4 returns count of value in list.
func UIntD4(list [][][][]uint, value uint) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UIntD3(listValue, value)
	}
	return valueCount
}

// UIntD5 returns count of value in list.
func UIntD5(list [][][][][]uint, value uint) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UIntD4(listValue, value)
	}
	return valueCount
}

// UInt8 returns count of value in list.
func UInt8(list []uint8, value uint8) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// UInt8D2 returns count of value in list.
func UInt8D2(list [][]uint8, value uint8) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt8(listValue, value)
	}
	return valueCount
}

// UInt8D3 returns count of value in list.
func UInt8D3(list [][][]uint8, value uint8) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt8D2(listValue, value)
	}
	return valueCount
}

// UInt8D4 returns count of value in list.
func UInt8D4(list [][][][]uint8, value uint8) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt8D3(listValue, value)
	}
	return valueCount
}

// UInt8D5 returns count of value in list.
func UInt8D5(list [][][][][]uint8, value uint8) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt8D4(listValue, value)
	}
	return valueCount
}

// UInt16 returns count of value in list.
func UInt16(list []uint16, value uint16) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// UInt16D2 returns count of value in list.
func UInt16D2(list [][]uint16, value uint16) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt16(listValue, value)
	}
	return valueCount
}

// UInt16D3 returns count of value in list.
func UInt16D3(list [][][]uint16, value uint16) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt16D2(listValue, value)
	}
	return valueCount
}

// UInt16D4 returns count of value in list.
func UInt16D4(list [][][][]uint16, value uint16) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt16D3(listValue, value)
	}
	return valueCount
}

// UInt16D5 returns count of value in list.
func UInt16D5(list [][][][][]uint16, value uint16) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt16D4(listValue, value)
	}
	return valueCount
}

// UInt32 returns count of value in list.
func UInt32(list []uint32, value uint32) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// UInt32D2 returns count of value in list.
func UInt32D2(list [][]uint32, value uint32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt32(listValue, value)
	}
	return valueCount
}

// UInt32D3 returns count of value in list.
func UInt32D3(list [][][]uint32, value uint32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt32D2(listValue, value)
	}
	return valueCount
}

// UInt32D4 returns count of value in list.
func UInt32D4(list [][][][]uint32, value uint32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt32D3(listValue, value)
	}
	return valueCount
}

// UInt32D5 returns count of value in list.
func UInt32D5(list [][][][][]uint32, value uint32) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt32D4(listValue, value)
	}
	return valueCount
}

// UInt64 returns count of value in list.
func UInt64(list []uint64, value uint64) int {
	valueCount := 0
	for _, listValue := range list {
		if listValue == value {
			valueCount++
		}
	}
	return valueCount
}

// UInt64D2 returns count of value in list.
func UInt64D2(list [][]uint64, value uint64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt64(listValue, value)
	}
	return valueCount
}

// UInt64D3 returns count of value in list.
func UInt64D3(list [][][]uint64, value uint64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt64D2(listValue, value)
	}
	return valueCount
}

// UInt64D4 returns count of value in list.
func UInt64D4(list [][][][]uint64, value uint64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt64D3(listValue, value)
	}
	return valueCount
}

// UInt64D5 returns count of value in list.
func UInt64D5(list [][][][][]uint64, value uint64) int {
	valueCount := 0
	for _, listValue := range list {
		valueCount += UInt64D4(listValue, value)
	}
	return valueCount
}
