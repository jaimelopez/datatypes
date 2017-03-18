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
