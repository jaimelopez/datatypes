package generic

const InvalidIterableElement = "Non-iterable type can not be converted to slice"

func NewInvalidIterableElementError() {
	panic(InvalidIterableElement)
}
