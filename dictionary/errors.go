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
