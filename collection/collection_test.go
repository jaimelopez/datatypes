package collection

import "testing"

func TestAddMethod(test *testing.T) {
	element := "first element"
	collection := NewEmptyCollection()
	collection.Add(element)

	if len(collection.elements) != 1 || collection.elements[0] != element {
		test.Error("Wrong behaviour adding a element")
	}
}

func TestAddRangeMethod(test *testing.T) {
	newElements := []string{"first element", "second element"}

	collection := NewEmptyCollection()
	error := collection.AddRange(newElements)

	if error != nil {
		test.Error("Wrong error returned adding a element range")
	}

	if len(collection.elements) != len(newElements) {
		test.Error("Wrong behaviour adding a element range")
	}

	// TODO
	//invalidRange := "simple string"
	//error = collection.AddRange(invalidRange)
	//
	//if error != InvalidIterableElement {
	//	test.Error("Method should return an InvalidIterableElement error adding an invalid range")
	//}
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

	error := collection.AddCollection(otherCollection)

	if error != nil {
		test.Error("Unexpected error adding a collection to another collection")
	}

	if len(collection.elements) != 3 {
		test.Error("Wrong elements number adding a collection to another collection")
	}
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

	error := collection.Delete(elementTwo)

	if error != nil {
		test.Error("Unexpected error delenting an element")
	}

	if len(collection.elements) != 2 {
		test.Error("Invalid number of elements after a element deletion")
	}

	if collection.elements[0] != elementOne || collection.elements[1] != elementThree {
		test.Error("Invalid expected elements after a single element deletion")
	}
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

	error := collection.DeleteRange([]string{elementOne, elementThree})

	if error != nil {
		test.Error("Unexpected error delenting a range elements")
	}

	if len(collection.elements) != 2 {
		test.Error("Wrong elements number deleting  a collection to another collection")
	}

	for _, element := range collection.elements {
		if element == elementOne || element == elementThree {
			test.Error("Elements not correctly deleted in DeleteRange method")
		}
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

	error := collection.DeleteCollection(otherCollection)

	if error != nil {
		test.Error("Unexpected error deleting a collection from another collection")
	}

	if len(collection.elements) != 1 || collection.elements[0] != elementTwo {
		test.Error("Wrong elements number deleting a collection from another collection")
	}
}

func TestContainsMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"
	inexistentElement := "inexistent element"

	collection := NewCollection([]string{elementOne, elementTwo})

	if !collection.Contains(elementOne) {
		test.Error("Contains return a false positive with existent elements")
	}

	if collection.Contains(inexistentElement) {
		test.Error("Contains return a false positive with inexistent elements")
	}
}

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

func TestFirstMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewCollection([]string{elementOne, elementTwo})

	if collection.First() != elementOne {
		test.Error("First method do not return the correct element")
	}
}

func TestLastMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewCollection([]string{
		elementOne,
		elementTwo,
	})

	if collection.Last() != elementTwo {
		test.Error("Last method do not return the correct element")
	}
}

func TestElementAtMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewCollection([]string{elementOne, elementTwo})

	if collection.ElementAt(0) != elementOne || collection.ElementAt(1) != elementTwo {
		test.Error("Wrong returned element in specific position on ElementAt method")
	}
}

func TestElementsMethod(test *testing.T) {
	elementOne := "first element"
	elementTwo := "second element"

	collection := NewEmptyCollection()

	if collection.Elements() != nil || len(collection.Elements()) != 0 {
		test.Error("Elements method should return no elements on new empty instance")
	}

	collection.elements = append(collection.elements, elementOne)
	collection.elements = append(collection.elements, elementTwo)

	if collection.Elements()[0] != elementOne || collection.Elements()[1] != elementTwo {
		test.Error("Elements method do not return the correct stored elements in the collection")
	}
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
