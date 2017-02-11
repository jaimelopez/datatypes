package generic

import "testing"

func TestToSliceMethod(test *testing.T) {
	element := "first element"
	collection := NewEmptyCollection()
	collection.Add(element)

	if len(collection.elements) != 1 || collection.elements[0] != element {
		test.Error("Wrong behaviour adding a element")
	}
}