// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// The datatypes/generic package includes some functionalities
// to treat in a simple way the 'generic' objects in Go

// This part of package contains the core behaviour

package generic

import "reflect"

// ToSlice converts a non knowed type to slice
func ToSlice(slice interface{}) ([]interface{}, error) {
	sliceValues := reflect.ValueOf(slice)

	if sliceValues.Kind() != reflect.Slice {
		return nil, ErrInvalidIterableElement
	}

	values := make([]interface{}, sliceValues.Len())

	for x := 0; x < sliceValues.Len(); x++ {
		values[x] = sliceValues.Index(x).Interface()
	}

	return values, nil
}

// AreSameType checks if two elements have the same type
func AreSameType(first interface{}, second interface{}) bool {
	return reflect.TypeOf(first) == reflect.TypeOf(second)
}
