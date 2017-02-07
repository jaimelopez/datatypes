package collection

import "testing"

func TestAddMethod(test *testing.T) {}

func TestAddRangeMethod(test *testing.T) {}

func TestAddCollectionMethod(test *testing.T) {}

func TestDeleteMethod(test *testing.T) {}

func TestDeleteRangeMethod(test *testing.T) {}

func TestContainsMethod(test *testing.T) {}

func TestContainsAnyMethod(test *testing.T) {}

func TestElementsMethod(test *testing.T) {}

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
