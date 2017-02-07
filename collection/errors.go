package collection

import (
	"errors"
	"fmt"
)

const DuplicatedElementErrorString = "Duplicated element in collection"
const ElementNotFoundErrorString = "Element not found"
const InvalidElementTypeErrorString = "Invalid element type: collection with %s types must to be homogeneous"

func NewDuplicatedElementError() error {
	return errors.New(DuplicatedElementErrorString)
}

func NewElementNotFoundError() error {
	return errors.New(ElementNotFoundErrorString)
}

func NewInvalidElementTypeError(collectionType string) error {
	message := fmt.Sprintf(InvalidElementTypeErrorString, collectionType)

	return errors.New(message)
}
