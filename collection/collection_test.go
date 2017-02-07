package collection

import (
	"testing"
)

func TestAddMethod(test *testing.T) {
	collection := NewEmptyCollection()
	collection.Add("first element")
}

func TestAddRangeMethod(test *testing.T) {
	newElements := []string{"first element", "second element"}

	collection := NewEmptyCollection()
	error := collection.AddRange(newElements)

	if len(collection.elements) != len(newElements) {
		test.Error("Wrong behaviour adding a element range")
	}

	if error != nil {
		test.Error("Wrong error returned adding a element range")
	}

	//invalidRange := "simple string"
	//error = collection.AddRange(invalidRange)
	//
	//if error != InvalidIterableElement {
	//	test.Error("Method should return an InvalidIterableElement error adding an invalid range")
	//}
}

func TestAddCollectionMethod(test *testing.T) {
	//newElements := []string { "first element", "second element" }
	//
	//collection := NewEmptyCollection()
	//collection.Add("first element")
}

func TestDeleteMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	elementThree := "third element"

	elements := []string{elementOne, elementTwo, elementThree}

	collection := NewCollection(elements)

	error := collection.Delete(elementTwo)

	if error != nil {
		test.Error("Unexpected error delenting an element")
	}

	if collection.Count() != 2 {
		test.Error("Invalid number of elements after a element deletion")
	}

	if collection.Elements()[0] != elementOne || collection.Elements()[1] != elementThree {
		test.Error("Invalid expected elements after a single element deletion")
	}
}

func TestDeleteRangeMethod(test *testing.T) {}

func TestContainsMethod(test *testing.T) {}

func TestElementAtMethod(test *testing.T) {}

func TestFirstMethod(test *testing.T) {}

func TestContainsAnyMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	elementThree := "third element"

	elements := []string{elementOne, elementTwo}
	inexistentElements := []string{elementThree}

	collection := NewCollection(elements)

	if !collection.ContainsAny(elements) {
		test.Error("ContainsAny return a false positive with existent elements")
	}

	if collection.ContainsAny(inexistentElements) {
		test.Error("ContainsAny return a false positive with inexistent elements")
	}
}

func TestElementsMethod(test *testing.T) {
	//elements := []string { "first element", "second element" }
	//collection := NewCollection(elements)
	//
	//if collection.elements != collection.Elements() {
	//	test.Error("")
	//}
	//
	//if collection.Elements() != elements {
	//	test.Error("")
	//}
}

func TestCountMethod(test *testing.T) {
	collection := NewEmptyCollection()

	if collection.Count() != 0 {
		test.Error("Count method returns wrong size of collection when it's empty")
	}

	collection.Add("first element")

	if collection.Count() == 0 {
		test.Error("Count method returns 0 size when collection has elements")
	}
}

func TestIsEmptyMethod(test *testing.T) {
	collection := NewEmptyCollection()

	if !collection.IsEmpty() {
		test.Error("Empty method returns true when it's really empty")
	}

	collection.Add("first element")

	if collection.IsEmpty() {
		test.Error("Empty method returns false when it's not really empty")
	}
}

func TestNewEmptyCollection(test *testing.T) {
	emptyCollection := NewEmptyCollection()

	if len(emptyCollection.elements) != 0 {
		test.Error("Empty collection must to be instancied with no elements")
	}
}

func TestNewCollection(test *testing.T) {
	elements := []int{1, 2, 3, 4}

	collection := NewCollection(elements)

	if len(collection.elements) != len(elements) {
		test.Error("New collection don't store elements parameters as elements")
	}
}
