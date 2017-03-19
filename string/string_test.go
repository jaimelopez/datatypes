// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// datatypes/string is a simple package which encapsulates
// some basic operations and functionality over strings

// This part of package contains the tests for the whole package

package string

import "testing"

func TestEmptyMethod(test *testing.T) {
	defaultString := ""
	notDefaultString := "  "

	if !IsDefault(defaultString) {
		test.Error("Wrong result in IsDefault method: it should return true with a default string")
	}

	if IsDefault(notDefaultString) {
		test.Error("Wrong result in IsDefault method: it should return false with a non-default string")
	}
}

func TestDefaultMethod(test *testing.T) {
	emptyString := "          "
	notEmptyString := "  string"

	if !IsEmpty(emptyString) {
		test.Error("Wrong result in IsEmpty method: it should return true with an empty string")
	}

	if IsEmpty(notEmptyString) {
		test.Error("Wrong result in IsEmpty method: it should return false with an non-empty string")
	}
}
