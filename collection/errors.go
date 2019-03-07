// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// The datatypes/collection package provides new structures and
// behaviours to the iteration of non-sorted unique element and homogeneous
// lists accepting primitives types and complex user structs as well.

// This part of package contains the typped errors that the package uses

package collection

import "errors"

// ErrInvalidElementType represents an error for invalid element type
var ErrInvalidElementType = errors.New("Invalid element type: collection must to be homogeneous")

// ErrDuplicatedElement represents an error for duplicated elements
var ErrDuplicatedElement = errors.New("Duplicated element in collection")

// ErrElementNotFound represents an error for not found elements
var ErrElementNotFound = errors.New("Element not found")
