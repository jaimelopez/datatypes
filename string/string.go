// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// datatypes/string is a simple package which encapsulates
// some basic operations and functionality over strings

// This part of package contains the core behaviour

package string

import "strings"

// Default string
const Default string = ""

// IsDefault check if string has a default value
func IsDefault(str string) bool {
	return str == Default
}

// IsEmpty checks if an string is empty or blank
func IsEmpty(str string) bool {
	return IsDefault(strings.Trim(str, " "))
}
