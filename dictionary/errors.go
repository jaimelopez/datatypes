package dictionary

import (
	"errors"
	"fmt"
)

const InvalidKeyValueElementTypeErrorString = "Invalid key-value element type: dictionary must be homogeneous to %s-%s"
const DuplicatedKeyErrorString = "Duplicated key in dictionary"
const EmptyDictionaryErrorString = "Empty dictionary can not be iterable"

func NewInvalidKeyValueElementTypeError(keyType string, valueType string) error {
	message := fmt.Sprintf(InvalidKeyValueElementTypeErrorString, keyType, valueType)

	return errors.New(message)
}

func NewDuplicatedKeyError() error {
	return errors.New(DuplicatedKeyErrorString)
}

func NewEmptyDictionaryErrorString() error {
	return errors.New(EmptyDictionaryErrorString)
}
