package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSliceMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	genericIterableObject := []string{elementOne, elementTwo}

	var slicedObject []interface{}

	assert.NotPanics(test, func() {
		slicedObject = ToSlice(genericIterableObject)
	}, "Iterable objects shouldn't return a panic")

	if len(genericIterableObject) != len(slicedObject) {
		test.Error("Invalid lenght in generic slice")
	}

	assert.Panics(test, func() {
		ToSlice("non-iterable object")
	}, "Non-iterable object does not return a panic")
}

func TestAreSameTypeMethod(test *testing.T) { /* @TODO */ }
