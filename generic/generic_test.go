// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// The datatypes/generic package includes some functionalities
// to treat in a simple way the 'generic' objects in Go

// This part of package contains the tests for the whole package

package generic

import (
	"testing"
)

func TestToSliceMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	genericIterableObject := []string{elementOne, elementTwo}
	slicedObject, err := ToSlice(genericIterableObject)

	if err != nil {
		test.Error("Iterable objects shouldn't return an error")
	}

	if len(genericIterableObject) != len(slicedObject) {
		test.Error("Invalid length in generic slice")
	}

	nonIterableSlicedObject, err := ToSlice("non-iterable object")

	if nonIterableSlicedObject != nil {
		test.Error("Non-iterable object should return nil")
	} else if err == nil {
		test.Error("Non-iterable object should return an error")
	}
}

func TestAreSameTypeMethod(test *testing.T) { /* @TODO */ }
