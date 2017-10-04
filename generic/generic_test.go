// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// The datatypes/generic package includes some functionalities
// to treat in a simple way the 'generic' objects in Go

// This part of package contains the tests for the whole package

package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSliceMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	genericIterableObject := []string{elementOne, elementTwo}

	var slicedObject []interface{}

	assert.NotPanics(test, func() {
		slicedObject = ToSlice(genericIterableObject)
	}, "Iterable objects shouldn't return a panic")

	if len(genericIterableObject) != len(slicedObject) {
		test.Error("Invalid length in generic slice")
	}

	assert.Panics(test, func() {
		ToSlice("non-iterable object")
	}, "Non-iterable object does not return a panic")
}

func TestAreSameTypeMethod(test *testing.T) { /* @TODO */ }
