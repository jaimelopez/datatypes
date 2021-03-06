// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// The datatypes/collection package provides new structures and
// behaviours to the iteration of non-sorted unique element and homogeneous
// lists accepting primitives types and complex user structs as well.

// This part of package contains the tests for the whole package

package collection

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMethod(test *testing.T) {
	element := "first element"
	collection := NewEmptyCollection()

	assert.Nil(test, collection.Add(element), "Wrong behaviour adding a element: bad error")
	assert.EqualValues(test, len(collection.elements), 1, "Wrong behaviour adding a element")
	assert.Equal(test, collection.elements[0], element, "Wrong behaviour adding a element")
	assert.Error(test, collection.Add(element), ErrDuplicatedElement, "Duplicated keys should return an error on Add method")
}

func TestAddRangeMethod(test *testing.T) {
	newElements := []string{"first element", "second element"}
	invalidRange := "simple string"

	collection := NewEmptyCollection()

	assert.Nil(test, collection.AddRange(newElements), "Wrong error returned adding a element range")
	assert.Equal(test, len(collection.elements), len(newElements), "Wrong behaviour adding a element range")
	assert.Error(test, collection.AddRange(invalidRange), ErrInvalidElementType, "Method should return an InvalidIterableElement error adding an invalid range")
}

func TestAddCollectionMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	elementThree := "third element"

	collection := NewEmptyCollection()
	collection.Add(elementOne)

	otherCollection := NewEmptyCollection()
	otherCollection.Add(elementTwo)
	otherCollection.Add(elementThree)

	assert.Nil(test, collection.AddCollection(otherCollection), "Unexpected error adding a collection to another collection")
	assert.Len(test, collection.elements, 3, "Wrong elements number adding a collection to another collection")
}

func TestFirstMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewEmptyCollection()
	collection.Add(elementOne)
	collection.Add(elementTwo)

	assert.Exactly(test, collection.First(), elementOne, "First method do not return the correct element")
}

func TestLastMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewCollection([]string{
		elementOne,
		elementTwo,
	})

	assert.Exactly(test, collection.Last(), elementTwo, "Last method do not return the correct element")
}

func TestElementAtMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewCollection([]string{elementOne, elementTwo})

	assert.Equal(test, collection.ElementAt(0), elementOne, "Wrong returned element in specific position on ElementAt method")
	assert.Equal(test, collection.ElementAt(1), elementTwo, "Wrong returned element in specific position on ElementAt method")
}

func TestElementsMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewEmptyCollection()

	assert.Nil(test, collection.Elements(), "Elements method should return no elements on new empty instance")
	assert.Empty(test, collection.Elements(), "Elements method should return no elements on new empty instance")

	collection.elements = append(collection.elements, elementOne)
	collection.elements = append(collection.elements, elementTwo)

	assert.Exactly(test, collection.Elements()[0], elementOne, "Elements method do not return the correct stored elements in the collection")
	assert.Exactly(test, collection.Elements()[1], elementTwo, "Elements method do not return the correct stored elements in the collection")
}

func TestExtractMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewCollection([]Element{elementOne, elementTwo})

	firstElement := collection.Extract()

	assert.Exactly(test, firstElement, elementOne, "Wrong extracted element on Extract method")
	assert.Len(test, collection.elements, 1, "Wrong remained elements in the collection on Extract method")

	newCollection := NewCollection([]Element{elementOne, elementTwo})
	counter := 0

	for !newCollection.IsEmpty() {
		_ = newCollection.Extract()

		counter++
	}

	assert.Exactly(test, counter, 2, "Wrong behaviour iterating over the collection with Extract method")
}

func TestSetMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "two element"

	collection := NewEmptyCollection()
	collection.Add(elementOne)
	collection.Set(0, elementTwo)

	assert.Exactly(test, collection.elements[0], elementTwo, "Set method doesn't works properly")
	assert.Len(test, collection.elements, 1, "Set method doesn't mantains the right items")
}

func TestDeleteMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	elementThree := "third element"

	collection := NewCollection([]string{
		elementOne,
		elementTwo,
		elementThree,
	})

	assert.Nil(test, collection.Delete(elementTwo), "Unexpected error delenting an element")
	assert.Len(test, collection.elements, 2, "Invalid number of elements after a element deletion")
	assert.Exactly(test, collection.elements[0], elementOne, "Invalid expected elements after a single element deletion")
	assert.Exactly(test, collection.elements[1], elementThree, "Invalid expected elements after a single element deletion")
}

func TestDeleteRangeMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	elementThree := "third element"
	elementFour := "fourth element"

	collection := NewCollection([]Element{
		elementOne,
		elementTwo,
		elementThree,
		elementFour,
	})

	assert.Nil(test, collection.DeleteRange([]string{elementOne, elementThree}), "Unexpected error delenting a range elements")
	assert.Len(test, collection.elements, 2, "Wrong elements number deleting  a collection to another collection")

	for _, element := range collection.elements {
		assert.NotEqual(test, element, elementOne, "Elements not correctly deleted in DeleteRange method")
		assert.NotEqual(test, element, elementThree, "Elements not correctly deleted in DeleteRange method")
	}
}

func TestDeleteCollectionMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	elementThree := "third element"

	collection := NewCollection([]string{
		elementOne,
		elementTwo,
		elementThree,
	})

	otherCollection := NewCollection([]Element{
		elementOne,
		elementThree,
	})

	assert.Nil(test, collection.DeleteCollection(otherCollection), "Unexpected error deleting a collection from another collection")
	assert.Len(test, collection.elements, 1, "Wrong elements number deleting a collection from another collection")
	assert.Exactly(test, collection.elements[0], elementTwo, "Wrong elements number deleting a collection from another collection")
}

func TestContainsMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	inexistentElement := "inexistent element"

	collection := NewCollection([]string{elementOne, elementTwo})

	assert.True(test, collection.Contains(elementOne), "Contains return a false positive with existent elements")
	assert.False(test, collection.Contains(inexistentElement), "Contains return a false positive with inexistent elements")
}

func TestContainsAnyMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	elementThree := "third element"

	elements := []string{elementOne, elementTwo}
	inexistentElements := []string{elementThree}

	collection := NewCollection(elements)

	assert.True(test, collection.ContainsAny(elements), "ContainsAny return a false positive with existent elements")
	assert.False(test, collection.ContainsAny(inexistentElements), "ContainsAny return a false positive with inexistent elements")
}

func TestFilterhMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	elementThree := "third element"

	collection := NewCollection([]Element{
		elementOne,
		elementTwo,
		elementThree,
	})

	matches := collection.Filter(func(elem Element) bool {
		return strings.Contains(elem.(string), "second")
	})

	assert.Exactly(test, matches[0], elementTwo, "Wrong filtered elements!")
	assert.NotContains(test, matches, elementOne, "Inapropiated element included in filtered results!")
	assert.NotContains(test, matches, elementThree, "Inapropiated element included in filtered results!")
}

func TestSizeMethod(test *testing.T) {
	collection := NewEmptyCollection()

	assert.Empty(test, collection.Size(), "Size method returns wrong size of collection when it's empty")

	collection.Add("first element")

	assert.NotZero(test, collection.Size(), "Size method returns 0 size when collection has elements")
}

func TestIsEmptyMethod(test *testing.T) {
	collection := NewEmptyCollection()

	assert.True(test, collection.IsEmpty(), "Empty method returns true when it's really empty")

	collection.Add("first element")

	assert.False(test, collection.IsEmpty(), "Empty method returns false when it's not really empty")
}

func TestNewEmptyCollection(test *testing.T) {
	emptyCollection := NewEmptyCollection()

	assert.Empty(test, emptyCollection.elements, "Empty collection must to be instancied with no elements")
}

func TestNewCollection(test *testing.T) {
	elements := []int{1, 2, 3, 4}

	collection := NewCollection(elements)

	assert.Len(test, collection.elements, len(elements), "New collection don't store elements parameters as elements")

	singleElement := "element as string"
	singleElementCollection := NewCollection(singleElement)

	assert.Len(test, singleElementCollection.elements, 1, "New collection with a single element don't instance the right value")
	assert.Exactly(test, singleElementCollection.elements[0], singleElement, "New collection with a single element don't instance the right value")
}
