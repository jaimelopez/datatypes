// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// The datatypes/collection package provides new structures and
// behaviours to the iteration of non-sorted unique element and homogeneous
// lists accepting primitives types and complex user structs as well.

// This part of package contains the typped errors that the package uses

package collection

import "fmt"

const InvalidElementTypeErrorString = "Invalid element type: collection with %s types must to be homogeneous"
const DuplicatedElementErrorString = "Duplicated element in collection"
const ElementNotFoundErrorString = "Element not found"

func NewInvalidElementTypeError(collectionType string) {
	message := fmt.Sprintf(InvalidElementTypeErrorString, collectionType)
	panic(message)
}

func NewDuplicatedElementError() {
	panic(DuplicatedElementErrorString)
}

func NewElementNotFoundError() {
	panic(ElementNotFoundErrorString)
}
