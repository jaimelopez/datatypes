// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// datatypes/string is a simple package which encapsulates
// some basic operations and functionality over strings

// This part of package contains the core behaviour

package string

import "strings"

const Default string = ""

func IsDefault(str string) bool {
	return str == Default
}

func IsEmpty(str string) bool {
	return IsDefault(strings.Trim(str, " "))
}
