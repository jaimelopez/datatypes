package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMethod(test *testing.T) {
	element := "first element"
	collection := NewEmptyCollection()

	assert.NotPanics(test, func() {
		collection.Add(element)
	}, "Wrong behaviour adding a element: bad error")

	if len(collection.elements) != 1 || collection.elements[0] != element {
		test.Error("Wrong behaviour adding a element")
	}

	assert.Panics(test, func() {
		collection.Add(element)
	}, "Duplicated keys should return an error on Add method")
}

func TestAddRangeMethod(test *testing.T) {
	newElements := []string{"first element", "second element"}

	collection := NewEmptyCollection()

	assert.NotPanics(test, func() {
		collection.AddRange(newElements)
	}, "Wrong error returned adding a element range")

	assert.Equal(test,
		len(collection.elements),
		len(newElements),
		"Wrong behaviour adding a element range")

	invalidRange := "simple string"

	assert.Panics(test, func() {
		collection.AddRange(invalidRange)
	}, "Method should return an InvalidIterableElement error adding an invalid range")
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

	assert.NotPanics(test, func() {
		collection.AddCollection(otherCollection)
	}, "Unexpected error adding a collection to another collection")

	assert.Len(test, collection.elements, 3, "Wrong elements number adding a collection to another collection")
}

func TestFirstMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewCollection([]string{elementOne, elementTwo})

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
	/* @TODO */
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

	assert.NotPanics(test, func() {
		collection.Delete(elementTwo)
	}, "Unexpected error delenting an element")

	assert.Len(test, collection.elements, 2, "Invalid number of elements after a element deletion")

	assert.Exactly(test, collection.elements[0], elementOne, "Invalid expected elements after a single element deletion")
	assert.Exactly(test, collection.elements[1], elementThree, "Invalid expected elements after a single element deletion")
}

func TestDeleteRangeMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	elementThree := "third element"
	elementFour := "fourth element"

	collection := NewCollection([]string{
		elementOne,
		elementTwo,
		elementThree,
		elementFour,
	})

	assert.NotPanics(test, func() {
		collection.DeleteRange([]string{elementOne, elementThree})
	}, "Unexpected error delenting a range elements")

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

	otherCollection := NewEmptyCollection()
	otherCollection.Add(elementOne)
	otherCollection.Add(elementThree)

	assert.NotPanics(test, func() {
		collection.DeleteCollection(otherCollection)
	}, "Unexpected error deleting a collection from another collection")

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

func TestCountMethod(test *testing.T) {
	collection := NewEmptyCollection()

	assert.Empty(test, collection.Count(), "Count method returns wrong size of collection when it's empty")

	collection.Add("first element")

	assert.NotZero(test, collection.Count(), "Count method returns 0 size when collection has elements")
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
