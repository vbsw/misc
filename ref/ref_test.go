/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package ref

import (
	"math/rand"
	"testing"
)

const wordSize = 1024
const chunksCount = 20

func TestBytes(t *testing.T) {
	str := "Hello, world!"
	bytes := []byte(str)
	bytesRef := Bytes(str)
	checkBytes(t, bytes, bytesRef)

	str = "Привет, мир!"
	bytes = []byte(str)
	bytesRef = Bytes(str)
	checkBytes(t, bytes, bytesRef)
}

func TestString(t *testing.T) {
	str := "Hello, world!"
	bytes := []byte(str)
	strRef := String(bytes)

	if str != strRef {
		t.Error(strRef)
	}
}

func BenchmarkBytes(b *testing.B) {
	strings := generateStrings(chunksCount, wordSize)
	bytesGen := make([][]byte, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes := Bytes(strings[i%chunksCount])
		if len(bytesGen[i]) == 0 {
			bytesGen[i] = bytes
		}
	}
}

func BenchmarkBytesStd(b *testing.B) {
	strings := generateStrings(chunksCount, wordSize)
	bytesGen := make([][]byte, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes := []byte(strings[i%chunksCount])
		if len(bytesGen[i]) == 0 {
			bytesGen[i] = bytes
		}
	}
}

func BenchmarkString(b *testing.B) {
	bytes := generateBytes(chunksCount, wordSize)
	stringsGen := make([]string, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		strings := String(bytes[i%chunksCount])
		if len(stringsGen[i]) == 0 {
			stringsGen[i] = strings
		}
	}
}

func BenchmarkStringStd(b *testing.B) {
	bytes := generateBytes(chunksCount, wordSize)
	stringsGen := make([]string, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		strings := string(bytes[i%chunksCount])
		if len(stringsGen[i]) == 0 {
			stringsGen[i] = strings
		}
	}
}

func checkBytes(t *testing.T, bytesA, bytesB []byte) {
	if len(bytesA) != len(bytesB) {
		t.Error(len(bytesA), len(bytesB))
	} else {
		for i, b := range bytesA {
			if b != bytesB[i] {
				t.Error(i, b, bytesB[i])
				break
			}
		}
	}
}

func generateStrings(wordsCount, wordSize int) []string {
	strings := make([]string, wordsCount)
	for i := range strings {
		bytes := make([]byte, wordSize)
		fillBytesASCII(bytes, 33, 126)
		strings[i] = string(bytes)
	}
	return strings
}

func generateBytes(wordsCount, wordSize int) [][]byte {
	bytes := make([][]byte, wordsCount)
	for i := range bytes {
		fillBytesASCII(bytes[i], 33, 126)
	}
	return bytes
}

func fillBytesASCII(bytes []byte, from, to int) {
	delta := to - from
	for i := range bytes {
		asciiByte := byte(rand.Int()%delta + from)
		bytes[i] = asciiByte
	}
}
