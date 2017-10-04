// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package datatypes/dictionary provides an easy dictionary (key => value) homogeneous
// struct management, making the iteration of a unique-key lists more powerful,
// simple and clean, accepting primitives types and complex user structs as well.

// This part of package contains the typped errors that the package uses

package dictionary

import "fmt"

// InvalidKeyValueElementTypeErrorString represents an error for invalid key-value type
const InvalidKeyValueElementTypeErrorString = "Invalid key-value element type: dictionary must be homogeneous to %s-%s"

// DuplicatedKeyErrorString represents an error for duplicated entries
const DuplicatedKeyErrorString = "Duplicated key in dictionary"

// EmptyDictionaryErrorString represents an error for non-iterable dictionaries
const EmptyDictionaryErrorString = "Empty dictionary can not be iterable"

// ElementNotFoundErrorString represents an error for non-found element
const ElementNotFoundErrorString = "Element not found"

// NewInvalidKeyValueElementTypeError error instance
func NewInvalidKeyValueElementTypeError(keyType string, valueType string) {
	message := fmt.Sprintf(InvalidKeyValueElementTypeErrorString, keyType, valueType)
	panic(message)
}

// NewDuplicatedKeyError error instance
func NewDuplicatedKeyError() {
	panic(DuplicatedKeyErrorString)
}

// NewEmptyDictionaryErrorString error instance
func NewEmptyDictionaryErrorString() {
	panic(EmptyDictionaryErrorString)
}

// NewElementNotFoundError error instance
func NewElementNotFoundError() {
	panic(ElementNotFoundErrorString)
}
