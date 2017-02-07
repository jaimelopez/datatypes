package collection

import (
	"errors"
	"fmt"
)

const DuplicatedElementErrorString = "Duplicated element in collection"
const ElementNotFoundErrorString = "Element not found"
const InvalidIterableElement = "Invalid iterable element"
const InvalidElementTypeErrorString = "Invalid element type: collection with %s types must to be homogeneous"

func NewDuplicatedElementError() error {
	return errors.New(DuplicatedElementErrorString)
}

func NewElementNotFoundError() error {
	return errors.New(ElementNotFoundErrorString)
}

func NewInvalidIterableElementError() error {
	return errors.New(InvalidIterableElement)
}

func NewInvalidElementTypeError(collectionType string) error {
	message := fmt.Sprintf(InvalidElementTypeErrorString, collectionType)

	return errors.New(message)
}
