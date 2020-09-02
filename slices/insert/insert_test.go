/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package insert

import (
	"testing"
)

func TestIntN1(t *testing.T) {
	insertedValues := newValues(3)
	values := newValues(2, 1, 2)

	values = IntN(values, 1, insertedValues)
	if len(values) != 2 {
		t.Error(len(values))
	} else if values[0] != 1 || values[1] != 2 {
		t.Error(values[0], values[1])
	}
}

func TestIntN2(t *testing.T) {
	insertedValues := newValues(3)
	values := newValues(2, 1, 2)

	insertedValues = append(insertedValues, 10, 11)
	values = IntN(values, 1, insertedValues)
	if len(values) != 4 {
		t.Error(len(values))
	} else if values[0] != 1 || values[1] != 10 || values[2] != 11 || values[3] != 2 {
		t.Error(values[0], values[1])
	}
}

func TestIntN3(t *testing.T) {
	insertedValues := newValues(3)
	values := newValues(3, 1, 2)

	insertedValues = append(insertedValues, 10, 11)
	values = IntN(values, 1, insertedValues)
	if len(values) != 4 {
		t.Error(len(values))
	} else if values[0] != 1 || values[1] != 10 || values[2] != 11 || values[3] != 2 {
		t.Error(values[0], values[1], values[2], values[3])
	}
}

func TestIntN4(t *testing.T) {
	insertedValues := newValues(3)
	values := newValues(4, 1, 2)

	insertedValues = append(insertedValues, 10, 11)
	values = IntN(values, 1, insertedValues)
	if len(values) != 4 || cap(values) != 4 {
		t.Error(len(values), cap(values))
	} else if values[0] != 1 || values[1] != 10 || values[2] != 11 || values[3] != 2 {
		t.Error(values[0], values[1], values[2], values[3])
	}
}

func newValues(capacity int, elements ...int) []int {
	values := make([]int, 0, capacity)
	values = append(values, elements...)
	return values
}
