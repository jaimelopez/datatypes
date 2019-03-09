// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package datatypes/dictionary provides an easy dictionary (key => value) homogeneous
// struct management, making the iteration of a unique-key lists more powerful,
// simple and clean, accepting primitives types and complex user structs as well.

// This part of package contains the typped errors that the package uses

package dictionary

import "errors"

// ErrInvalidKeyValueElementType represents an error for invalid key-value type
var ErrInvalidKeyValueElementType = errors.New("Invalid key-value element type: dictionary must be homogeneous")

// ErrDuplicatedKey represents an error for duplicated entries
var ErrDuplicatedKey = errors.New("Duplicated key in dictionary")

// ErrEmptyDictionary represents an error for non-iterable dictionaries
var ErrEmptyDictionary = errors.New("Empty dictionary can not be iterable")

// ErrElementNotFound represents an error for non-found element
var ErrElementNotFound = errors.New("Element not found")
