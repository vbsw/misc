/*
 *       Copyright 2016, 2019 Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package semver

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	verStr := String(1, 0, 0, "", "")
	if strings.Compare(verStr, "1.0.0") != 0 {
		t.Error(verStr)
	}

	verStr = String(1, 0, 10, "", "")
	if strings.Compare(verStr, "1.0.10") != 0 {
		t.Error(verStr)
	}

	verStr = String(1, 0, 10001, "", "")
	if strings.Compare(verStr, "1.0.10001") != 0 {
		t.Error(verStr)
	}

	verStr = String(143121, 0, 10001, "", "")
	if strings.Compare(verStr, "143121.0.10001") != 0 {
		t.Error(verStr)
	}

	verStr = String(143121, 0, 10001, "", "")
	if strings.Compare(verStr, "143121.0.10001") != 0 {
		t.Error(verStr)
	}
}

func TestStringExt(t *testing.T) {
	verStr := String(1, 0, 0, "alpha", "")
	if strings.Compare(verStr, "1.0.0-alpha") != 0 {
		t.Error(verStr)
	}

	verStr = String(1, 0, 0, "z1.1", "1234")
	if strings.Compare(verStr, "1.0.0-z1.1+1234") != 0 {
		t.Error(verStr)
	}
}
