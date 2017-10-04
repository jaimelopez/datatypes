// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package datatypes/dictionary provides an easy dictionary (key => value) homogeneous
// struct management, making the iteration of a unique-key lists more powerful,
// simple and clean, accepting primitives types and complex user structs as well.

// This part of package contains the typped errors that the package uses

package dictionary

import "fmt"

const InvalidKeyValueElementTypeErrorString = "Invalid key-value element type: dictionary must be homogeneous to %s-%s"
const DuplicatedKeyErrorString = "Duplicated key in dictionary"
const EmptyDictionaryErrorString = "Empty dictionary can not be iterable"
const ElementNotFoundErrorString = "Element not found"

func NewInvalidKeyValueElementTypeError(keyType string, valueType string) {
	message := fmt.Sprintf(InvalidKeyValueElementTypeErrorString, keyType, valueType)
	panic(message)
}

func NewDuplicatedKeyError() {
	panic(DuplicatedKeyErrorString)
}

func NewEmptyDictionaryErrorString() {
	panic(EmptyDictionaryErrorString)
}

func NewElementNotFoundError() {
	panic(ElementNotFoundErrorString)
}
