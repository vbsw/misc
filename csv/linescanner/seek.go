/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package linescanner

func seekLineEnd(bytes []byte, from, to int) (int, int) {
	for i := from; i < to; i++ {
		if bytes[i] == '\n' {
			return i, i + 1
		} else if bytes[i] == '\r' {
			if i+1 < to && bytes[i+1] == '\n' {
				return i, i + 2
			}
			return i, i + 1
		}
	}
	return to, to
}

func seekContent(bytes []byte, from, to int) int {
	for i := from; i < to; i++ {
		if bytes[i] < 0 || bytes[i] > 32 {
			return i
		}
	}
	return to
}

func seekContentRight(bytes []byte, from, to int) int {
	for i := to - 1; i >= from; i-- {
		if bytes[i] < 0 || bytes[i] > 32 {
			return i + 1
		}
	}
	return from
}

func seekBytes(bytes, word []byte, from, to int) int {
	if len(word) > 0 {
		limit := to - len(word) + 1
	outter:
		for i := from; i < limit; i++ {
			for j, b := range word {
				if bytes[i+j] != b {
					continue outter
				}
			}
			return i
		}
	}
	return to
}
