package generic

import "testing"

func TestToSliceMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	genericIterableObject := []string{elementOne, elementTwo}

	slicedObject := ToSlice(genericIterableObject)

	if len(genericIterableObject) != len(slicedObject) {
		test.Error("ushdushduhs")
	}

	defer func() {
		if recover() == nil {
			test.Error("Non-iterable object does not return a panic")
		}
	}()

	ToSlice("non-iterable object")
}
