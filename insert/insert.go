/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package insert provides the "insert" function for slices of basic types.
package insert

import (
	"unsafe"
)

// Bool inserts value in list.
func Bool(list []bool, index int, value bool) []bool {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// BoolD2 inserts value in list.
func BoolD2(list [][]bool, index int, value []bool) [][]bool {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// BoolD3 inserts value in list.
func BoolD3(list [][][]bool, index int, value [][]bool) [][][]bool {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// BoolD4 inserts value in list.
func BoolD4(list [][][][]bool, index int, value [][][]bool) [][][][]bool {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// BoolD5 inserts value in list.
func BoolD5(list [][][][][]bool, index int, value [][][][]bool) [][][][][]bool {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// BoolN inserts values in list.
func BoolN(list []bool, index int, values []bool) []bool {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// BoolND2 inserts values in list.
func BoolND2(list [][]bool, index int, values [][]bool) [][]bool {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// BoolND3 inserts values in list.
func BoolND3(list [][][]bool, index int, values [][][]bool) [][][]bool {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// BoolND4 inserts values in list.
func BoolND4(list [][][][]bool, index int, values [][][][]bool) [][][][]bool {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// BoolND5 inserts values in list.
func BoolND5(list [][][][][]bool, index int, values [][][][][]bool) [][][][][]bool {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Byte inserts value in list.
func Byte(list []byte, index int, value byte) []byte {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ByteD2 inserts value in list.
func ByteD2(list [][]byte, index int, value []byte) [][]byte {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ByteD3 inserts value in list.
func ByteD3(list [][][]byte, index int, value [][]byte) [][][]byte {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ByteD4 inserts value in list.
func ByteD4(list [][][][]byte, index int, value [][][]byte) [][][][]byte {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ByteD5 inserts value in list.
func ByteD5(list [][][][][]byte, index int, value [][][][]byte) [][][][][]byte {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ByteN inserts values in list.
func ByteN(list []byte, index int, values []byte) []byte {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// ByteND2 inserts values in list.
func ByteND2(list [][]byte, index int, values [][]byte) [][]byte {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// ByteND3 inserts values in list.
func ByteND3(list [][][]byte, index int, values [][][]byte) [][][]byte {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// ByteND4 inserts values in list.
func ByteND4(list [][][][]byte, index int, values [][][][]byte) [][][][]byte {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// ByteND5 inserts values in list.
func ByteND5(list [][][][][]byte, index int, values [][][][][]byte) [][][][][]byte {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex64 inserts value in list.
func Complex64(list []complex64, index int, value complex64) []complex64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex64D2 inserts value in list.
func Complex64D2(list [][]complex64, index int, value []complex64) [][]complex64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex64D3 inserts value in list.
func Complex64D3(list [][][]complex64, index int, value [][]complex64) [][][]complex64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex64D4 inserts value in list.
func Complex64D4(list [][][][]complex64, index int, value [][][]complex64) [][][][]complex64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex64D5 inserts value in list.
func Complex64D5(list [][][][][]complex64, index int, value [][][][]complex64) [][][][][]complex64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex64N inserts values in list.
func Complex64N(list []complex64, index int, values []complex64) []complex64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex64ND2 inserts values in list.
func Complex64ND2(list [][]complex64, index int, values [][]complex64) [][]complex64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex64ND3 inserts values in list.
func Complex64ND3(list [][][]complex64, index int, values [][][]complex64) [][][]complex64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex64ND4 inserts values in list.
func Complex64ND4(list [][][][]complex64, index int, values [][][][]complex64) [][][][]complex64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex64ND5 inserts values in list.
func Complex64ND5(list [][][][][]complex64, index int, values [][][][][]complex64) [][][][][]complex64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex128 inserts value in list.
func Complex128(list []complex128, index int, value complex128) []complex128 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex128D2 inserts value in list.
func Complex128D2(list [][]complex128, index int, value []complex128) [][]complex128 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex128D3 inserts value in list.
func Complex128D3(list [][][]complex128, index int, value [][]complex128) [][][]complex128 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex128D4 inserts value in list.
func Complex128D4(list [][][][]complex128, index int, value [][][]complex128) [][][][]complex128 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex128D5 inserts value in list.
func Complex128D5(list [][][][][]complex128, index int, value [][][][]complex128) [][][][][]complex128 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Complex128N inserts values in list.
func Complex128N(list []complex128, index int, values []complex128) []complex128 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex128ND2 inserts values in list.
func Complex128ND2(list [][]complex128, index int, values [][]complex128) [][]complex128 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex128ND3 inserts values in list.
func Complex128ND3(list [][][]complex128, index int, values [][][]complex128) [][][]complex128 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex128ND4 inserts values in list.
func Complex128ND4(list [][][][]complex128, index int, values [][][][]complex128) [][][][]complex128 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Complex128ND5 inserts values in list.
func Complex128ND5(list [][][][][]complex128, index int, values [][][][][]complex128) [][][][][]complex128 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Error inserts value in list.
func Error(list []error, index int, value error) []error {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ErrorD2 inserts value in list.
func ErrorD2(list [][]error, index int, value []error) [][]error {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ErrorD3 inserts value in list.
func ErrorD3(list [][][]error, index int, value [][]error) [][][]error {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ErrorD4 inserts value in list.
func ErrorD4(list [][][][]error, index int, value [][][]error) [][][][]error {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ErrorD5 inserts value in list.
func ErrorD5(list [][][][][]error, index int, value [][][][]error) [][][][][]error {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// ErrorN inserts values in list.
func ErrorN(list []error, index int, values []error) []error {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// ErrorND2 inserts values in list.
func ErrorND2(list [][]error, index int, values [][]error) [][]error {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// ErrorND3 inserts values in list.
func ErrorND3(list [][][]error, index int, values [][][]error) [][][]error {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// ErrorND4 inserts values in list.
func ErrorND4(list [][][][]error, index int, values [][][][]error) [][][][]error {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// ErrorND5 inserts values in list.
func ErrorND5(list [][][][][]error, index int, values [][][][][]error) [][][][][]error {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float32 inserts value in list.
func Float32(list []float32, index int, value float32) []float32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float32D2 inserts value in list.
func Float32D2(list [][]float32, index int, value []float32) [][]float32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float32D3 inserts value in list.
func Float32D3(list [][][]float32, index int, value [][]float32) [][][]float32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float32D4 inserts value in list.
func Float32D4(list [][][][]float32, index int, value [][][]float32) [][][][]float32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float32D5 inserts value in list.
func Float32D5(list [][][][][]float32, index int, value [][][][]float32) [][][][][]float32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float32N inserts values in list.
func Float32N(list []float32, index int, values []float32) []float32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float32ND2 inserts values in list.
func Float32ND2(list [][]float32, index int, values [][]float32) [][]float32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float32ND3 inserts values in list.
func Float32ND3(list [][][]float32, index int, values [][][]float32) [][][]float32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float32ND4 inserts values in list.
func Float32ND4(list [][][][]float32, index int, values [][][][]float32) [][][][]float32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float32ND5 inserts values in list.
func Float32ND5(list [][][][][]float32, index int, values [][][][][]float32) [][][][][]float32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float64 inserts value in list.
func Float64(list []float64, index int, value float64) []float64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float64D2 inserts value in list.
func Float64D2(list [][]float64, index int, value []float64) [][]float64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float64D3 inserts value in list.
func Float64D3(list [][][]float64, index int, value [][]float64) [][][]float64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float64D4 inserts value in list.
func Float64D4(list [][][][]float64, index int, value [][][]float64) [][][][]float64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float64D5 inserts value in list.
func Float64D5(list [][][][][]float64, index int, value [][][][]float64) [][][][][]float64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Float64N inserts values in list.
func Float64N(list []float64, index int, values []float64) []float64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float64ND2 inserts values in list.
func Float64ND2(list [][]float64, index int, values [][]float64) [][]float64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float64ND3 inserts values in list.
func Float64ND3(list [][][]float64, index int, values [][][]float64) [][][]float64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float64ND4 inserts values in list.
func Float64ND4(list [][][][]float64, index int, values [][][][]float64) [][][][]float64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Float64ND5 inserts values in list.
func Float64ND5(list [][][][][]float64, index int, values [][][][][]float64) [][][][][]float64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int inserts value in list.
func Int(list []int, index int, value int) []int {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// IntD2 inserts value in list.
func IntD2(list [][]int, index int, value []int) [][]int {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// IntD3 inserts value in list.
func IntD3(list [][][]int, index int, value [][]int) [][][]int {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// IntD4 inserts value in list.
func IntD4(list [][][][]int, index int, value [][][]int) [][][][]int {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// IntD5 inserts value in list.
func IntD5(list [][][][][]int, index int, value [][][][]int) [][][][][]int {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// IntN inserts values in list.
func IntN(list []int, index int, values []int) []int {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// IntND2 inserts values in list.
func IntND2(list [][]int, index int, values [][]int) [][]int {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// IntND3 inserts values in list.
func IntND3(list [][][]int, index int, values [][][]int) [][][]int {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// IntND4 inserts values in list.
func IntND4(list [][][][]int, index int, values [][][][]int) [][][][]int {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// IntND5 inserts values in list.
func IntND5(list [][][][][]int, index int, values [][][][][]int) [][][][][]int {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int8 inserts value in list.
func Int8(list []int8, index int, value int8) []int8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int8D2 inserts value in list.
func Int8D2(list [][]int8, index int, value []int8) [][]int8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int8D3 inserts value in list.
func Int8D3(list [][][]int8, index int, value [][]int8) [][][]int8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int8D4 inserts value in list.
func Int8D4(list [][][][]int8, index int, value [][][]int8) [][][][]int8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int8D5 inserts value in list.
func Int8D5(list [][][][][]int8, index int, value [][][][]int8) [][][][][]int8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int8N inserts values in list.
func Int8N(list []int8, index int, values []int8) []int8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int8ND2 inserts values in list.
func Int8ND2(list [][]int8, index int, values [][]int8) [][]int8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int8ND3 inserts values in list.
func Int8ND3(list [][][]int8, index int, values [][][]int8) [][][]int8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int8ND4 inserts values in list.
func Int8ND4(list [][][][]int8, index int, values [][][][]int8) [][][][]int8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int8ND5 inserts values in list.
func Int8ND5(list [][][][][]int8, index int, values [][][][][]int8) [][][][][]int8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int16 inserts value in list.
func Int16(list []int16, index int, value int16) []int16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int16D2 inserts value in list.
func Int16D2(list [][]int16, index int, value []int16) [][]int16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int16D3 inserts value in list.
func Int16D3(list [][][]int16, index int, value [][]int16) [][][]int16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int16D4 inserts value in list.
func Int16D4(list [][][][]int16, index int, value [][][]int16) [][][][]int16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int16D5 inserts value in list.
func Int16D5(list [][][][][]int16, index int, value [][][][]int16) [][][][][]int16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int16N inserts values in list.
func Int16N(list []int16, index int, values []int16) []int16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int16ND2 inserts values in list.
func Int16ND2(list [][]int16, index int, values [][]int16) [][]int16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int16ND3 inserts values in list.
func Int16ND3(list [][][]int16, index int, values [][][]int16) [][][]int16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int16ND4 inserts values in list.
func Int16ND4(list [][][][]int16, index int, values [][][][]int16) [][][][]int16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int16ND5 inserts values in list.
func Int16ND5(list [][][][][]int16, index int, values [][][][][]int16) [][][][][]int16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int32 inserts value in list.
func Int32(list []int32, index int, value int32) []int32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int32D2 inserts value in list.
func Int32D2(list [][]int32, index int, value []int32) [][]int32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int32D3 inserts value in list.
func Int32D3(list [][][]int32, index int, value [][]int32) [][][]int32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int32D4 inserts value in list.
func Int32D4(list [][][][]int32, index int, value [][][]int32) [][][][]int32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int32D5 inserts value in list.
func Int32D5(list [][][][][]int32, index int, value [][][][]int32) [][][][][]int32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int32N inserts values in list.
func Int32N(list []int32, index int, values []int32) []int32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int32ND2 inserts values in list.
func Int32ND2(list [][]int32, index int, values [][]int32) [][]int32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int32ND3 inserts values in list.
func Int32ND3(list [][][]int32, index int, values [][][]int32) [][][]int32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int32ND4 inserts values in list.
func Int32ND4(list [][][][]int32, index int, values [][][][]int32) [][][][]int32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int32ND5 inserts values in list.
func Int32ND5(list [][][][][]int32, index int, values [][][][][]int32) [][][][][]int32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int64 inserts value in list.
func Int64(list []int64, index int, value int64) []int64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int64D2 inserts value in list.
func Int64D2(list [][]int64, index int, value []int64) [][]int64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int64D3 inserts value in list.
func Int64D3(list [][][]int64, index int, value [][]int64) [][][]int64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int64D4 inserts value in list.
func Int64D4(list [][][][]int64, index int, value [][][]int64) [][][][]int64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int64D5 inserts value in list.
func Int64D5(list [][][][][]int64, index int, value [][][][]int64) [][][][][]int64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// Int64N inserts values in list.
func Int64N(list []int64, index int, values []int64) []int64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int64ND2 inserts values in list.
func Int64ND2(list [][]int64, index int, values [][]int64) [][]int64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int64ND3 inserts values in list.
func Int64ND3(list [][][]int64, index int, values [][][]int64) [][][]int64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int64ND4 inserts values in list.
func Int64ND4(list [][][][]int64, index int, values [][][][]int64) [][][][]int64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Int64ND5 inserts values in list.
func Int64ND5(list [][][][][]int64, index int, values [][][][][]int64) [][][][][]int64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Interface inserts value in list.
func Interface(list []interface{}, index int, value interface{}) []interface{} {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// InterfaceD2 inserts value in list.
func InterfaceD2(list [][]interface{}, index int, value []interface{}) [][]interface{} {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// InterfaceD3 inserts value in list.
func InterfaceD3(list [][][]interface{}, index int, value [][]interface{}) [][][]interface{} {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// InterfaceD4 inserts value in list.
func InterfaceD4(list [][][][]interface{}, index int, value [][][]interface{}) [][][][]interface{} {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// InterfaceD5 inserts value in list.
func InterfaceD5(list [][][][][]interface{}, index int, value [][][][]interface{}) [][][][][]interface{} {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// InterfaceN inserts values in list.
func InterfaceN(list []interface{}, index int, values []interface{}) []interface{} {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// InterfaceND2 inserts values in list.
func InterfaceND2(list [][]interface{}, index int, values [][]interface{}) [][]interface{} {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// InterfaceND3 inserts values in list.
func InterfaceND3(list [][][]interface{}, index int, values [][][]interface{}) [][][]interface{} {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// InterfaceND4 inserts values in list.
func InterfaceND4(list [][][][]interface{}, index int, values [][][][]interface{}) [][][][]interface{} {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// InterfaceND5 inserts values in list.
func InterfaceND5(list [][][][][]interface{}, index int, values [][][][][]interface{}) [][][][][]interface{} {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Pointer inserts value in list.
func Pointer(list []unsafe.Pointer, index int, value unsafe.Pointer) []unsafe.Pointer {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// PointerD2 inserts value in list.
func PointerD2(list [][]unsafe.Pointer, index int, value []unsafe.Pointer) [][]unsafe.Pointer {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// PointerD3 inserts value in list.
func PointerD3(list [][][]unsafe.Pointer, index int, value [][]unsafe.Pointer) [][][]unsafe.Pointer {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// PointerD4 inserts value in list.
func PointerD4(list [][][][]unsafe.Pointer, index int, value [][][]unsafe.Pointer) [][][][]unsafe.Pointer {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// PointerD5 inserts value in list.
func PointerD5(list [][][][][]unsafe.Pointer, index int, value [][][][]unsafe.Pointer) [][][][][]unsafe.Pointer {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// PointerN inserts values in list.
func PointerN(list []unsafe.Pointer, index int, values []unsafe.Pointer) []unsafe.Pointer {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// PointerND2 inserts values in list.
func PointerND2(list [][]unsafe.Pointer, index int, values [][]unsafe.Pointer) [][]unsafe.Pointer {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// PointerND3 inserts values in list.
func PointerND3(list [][][]unsafe.Pointer, index int, values [][][]unsafe.Pointer) [][][]unsafe.Pointer {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// PointerND4 inserts values in list.
func PointerND4(list [][][][]unsafe.Pointer, index int, values [][][][]unsafe.Pointer) [][][][]unsafe.Pointer {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// PointerND5 inserts values in list.
func PointerND5(list [][][][][]unsafe.Pointer, index int, values [][][][][]unsafe.Pointer) [][][][][]unsafe.Pointer {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// Rune inserts value in list.
func Rune(list []rune, index int, value rune) []rune {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// RuneD2 inserts value in list.
func RuneD2(list [][]rune, index int, value []rune) [][]rune {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// RuneD3 inserts value in list.
func RuneD3(list [][][]rune, index int, value [][]rune) [][][]rune {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// RuneD4 inserts value in list.
func RuneD4(list [][][][]rune, index int, value [][][]rune) [][][][]rune {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// RuneD5 inserts value in list.
func RuneD5(list [][][][][]rune, index int, value [][][][]rune) [][][][][]rune {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// RuneN inserts values in list.
func RuneN(list []rune, index int, values []rune) []rune {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// RuneND2 inserts values in list.
func RuneND2(list [][]rune, index int, values [][]rune) [][]rune {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// RuneND3 inserts values in list.
func RuneND3(list [][][]rune, index int, values [][][]rune) [][][]rune {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// RuneND4 inserts values in list.
func RuneND4(list [][][][]rune, index int, values [][][][]rune) [][][][]rune {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// RuneND5 inserts values in list.
func RuneND5(list [][][][][]rune, index int, values [][][][][]rune) [][][][][]rune {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// String inserts value in list.
func String(list []string, index int, value string) []string {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// StringD2 inserts value in list.
func StringD2(list [][]string, index int, value []string) [][]string {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// StringD3 inserts value in list.
func StringD3(list [][][]string, index int, value [][]string) [][][]string {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// StringD4 inserts value in list.
func StringD4(list [][][][]string, index int, value [][][]string) [][][][]string {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// StringD5 inserts value in list.
func StringD5(list [][][][][]string, index int, value [][][][]string) [][][][][]string {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// StringN inserts values in list.
func StringN(list []string, index int, values []string) []string {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// StringND2 inserts values in list.
func StringND2(list [][]string, index int, values [][]string) [][]string {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// StringND3 inserts values in list.
func StringND3(list [][][]string, index int, values [][][]string) [][][]string {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// StringND4 inserts values in list.
func StringND4(list [][][][]string, index int, values [][][][]string) [][][][]string {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// StringND5 inserts values in list.
func StringND5(list [][][][][]string, index int, values [][][][][]string) [][][][][]string {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt inserts value in list.
func UInt(list []uint, index int, value uint) []uint {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UIntD2 inserts value in list.
func UIntD2(list [][]uint, index int, value []uint) [][]uint {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UIntD3 inserts value in list.
func UIntD3(list [][][]uint, index int, value [][]uint) [][][]uint {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UIntD4 inserts value in list.
func UIntD4(list [][][][]uint, index int, value [][][]uint) [][][][]uint {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UIntD5 inserts value in list.
func UIntD5(list [][][][][]uint, index int, value [][][][]uint) [][][][][]uint {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UIntN inserts values in list.
func UIntN(list []uint, index int, values []uint) []uint {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UIntND2 inserts values in list.
func UIntND2(list [][]uint, index int, values [][]uint) [][]uint {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UIntND3 inserts values in list.
func UIntND3(list [][][]uint, index int, values [][][]uint) [][][]uint {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UIntND4 inserts values in list.
func UIntND4(list [][][][]uint, index int, values [][][][]uint) [][][][]uint {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UIntND5 inserts values in list.
func UIntND5(list [][][][][]uint, index int, values [][][][][]uint) [][][][][]uint {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt8 inserts value in list.
func UInt8(list []uint8, index int, value uint8) []uint8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt8D2 inserts value in list.
func UInt8D2(list [][]uint8, index int, value []uint8) [][]uint8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt8D3 inserts value in list.
func UInt8D3(list [][][]uint8, index int, value [][]uint8) [][][]uint8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt8D4 inserts value in list.
func UInt8D4(list [][][][]uint8, index int, value [][][]uint8) [][][][]uint8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt8D5 inserts value in list.
func UInt8D5(list [][][][][]uint8, index int, value [][][][]uint8) [][][][][]uint8 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt8N inserts values in list.
func UInt8N(list []uint8, index int, values []uint8) []uint8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt8ND2 inserts values in list.
func UInt8ND2(list [][]uint8, index int, values [][]uint8) [][]uint8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt8ND3 inserts values in list.
func UInt8ND3(list [][][]uint8, index int, values [][][]uint8) [][][]uint8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt8ND4 inserts values in list.
func UInt8ND4(list [][][][]uint8, index int, values [][][][]uint8) [][][][]uint8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt8ND5 inserts values in list.
func UInt8ND5(list [][][][][]uint8, index int, values [][][][][]uint8) [][][][][]uint8 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt16 inserts value in list.
func UInt16(list []uint16, index int, value uint16) []uint16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt16D2 inserts value in list.
func UInt16D2(list [][]uint16, index int, value []uint16) [][]uint16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt16D3 inserts value in list.
func UInt16D3(list [][][]uint16, index int, value [][]uint16) [][][]uint16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt16D4 inserts value in list.
func UInt16D4(list [][][][]uint16, index int, value [][][]uint16) [][][][]uint16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt16D5 inserts value in list.
func UInt16D5(list [][][][][]uint16, index int, value [][][][]uint16) [][][][][]uint16 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt16N inserts values in list.
func UInt16N(list []uint16, index int, values []uint16) []uint16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt16ND2 inserts values in list.
func UInt16ND2(list [][]uint16, index int, values [][]uint16) [][]uint16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt16ND3 inserts values in list.
func UInt16ND3(list [][][]uint16, index int, values [][][]uint16) [][][]uint16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt16ND4 inserts values in list.
func UInt16ND4(list [][][][]uint16, index int, values [][][][]uint16) [][][][]uint16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt16ND5 inserts values in list.
func UInt16ND5(list [][][][][]uint16, index int, values [][][][][]uint16) [][][][][]uint16 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt32 inserts value in list.
func UInt32(list []uint32, index int, value uint32) []uint32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt32D2 inserts value in list.
func UInt32D2(list [][]uint32, index int, value []uint32) [][]uint32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt32D3 inserts value in list.
func UInt32D3(list [][][]uint32, index int, value [][]uint32) [][][]uint32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt32D4 inserts value in list.
func UInt32D4(list [][][][]uint32, index int, value [][][]uint32) [][][][]uint32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt32D5 inserts value in list.
func UInt32D5(list [][][][][]uint32, index int, value [][][][]uint32) [][][][][]uint32 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt32N inserts values in list.
func UInt32N(list []uint32, index int, values []uint32) []uint32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt32ND2 inserts values in list.
func UInt32ND2(list [][]uint32, index int, values [][]uint32) [][]uint32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt32ND3 inserts values in list.
func UInt32ND3(list [][][]uint32, index int, values [][][]uint32) [][][]uint32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt32ND4 inserts values in list.
func UInt32ND4(list [][][][]uint32, index int, values [][][][]uint32) [][][][]uint32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt32ND5 inserts values in list.
func UInt32ND5(list [][][][][]uint32, index int, values [][][][][]uint32) [][][][][]uint32 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt64 inserts value in list.
func UInt64(list []uint64, index int, value uint64) []uint64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt64D2 inserts value in list.
func UInt64D2(list [][]uint64, index int, value []uint64) [][]uint64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt64D3 inserts value in list.
func UInt64D3(list [][][]uint64, index int, value [][]uint64) [][][]uint64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt64D4 inserts value in list.
func UInt64D4(list [][][][]uint64, index int, value [][][]uint64) [][][][]uint64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt64D5 inserts value in list.
func UInt64D5(list [][][][][]uint64, index int, value [][][][]uint64) [][][][][]uint64 {
	extendedList := append(list, value)
	if index < len(list) {
		copy(extendedList[index+1:], list[index:])
		extendedList[index] = value
	}
	return extendedList
}

// UInt64N inserts values in list.
func UInt64N(list []uint64, index int, values []uint64) []uint64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt64ND2 inserts values in list.
func UInt64ND2(list [][]uint64, index int, values [][]uint64) [][]uint64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt64ND3 inserts values in list.
func UInt64ND3(list [][][]uint64, index int, values [][][]uint64) [][][]uint64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt64ND4 inserts values in list.
func UInt64ND4(list [][][][]uint64, index int, values [][][][]uint64) [][][][]uint64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}

// UInt64ND5 inserts values in list.
func UInt64ND5(list [][][][][]uint64, index int, values [][][][][]uint64) [][][][][]uint64 {
	if len(values) > 0 {
		if index < len(list) {
			if len(values) <= cap(list)-len(list) {
				list = append(list[:index+len(values)], list[index:]...)
				copy(list[index:], values)
			} else {
				extendedList := append(list, values...)
				copy(extendedList[index+len(values):], list[index:])
				copy(extendedList[index:], values)
				list = extendedList
			}
		} else {
			list = append(list, values...)
		}
	}
	return list
}
