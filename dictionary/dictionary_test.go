package dictionary

import "testing"

func TestAddMethod(test *testing.T) {
	keyElement := "key"
	valueElement := "value"

	dictionary := NewEmptyDictionary()
	error := dictionary.Add(keyElement, valueElement)

	if error != nil {
		test.Error("Unexpected error adding a element")
	}

	if len(dictionary.elements) != 1 || dictionary.elements[keyElement] != valueElement {
		test.Error("Wrong behaviour adding a element")
	}

	error = dictionary.Add(keyElement, valueElement)

	if error == nil {
		test.Error("Duplicated keys should return an error on Add method")
	}

	differentTypeKeyElement, differentTypeValueElement := 1, 2

	if dictionary.Add(differentTypeKeyElement, differentTypeValueElement) == nil {
		test.Error("Add method should throw an exception trying to insert a non-homogeneous elements")
	}
}

func TestAddKeyValueElementMethod(test *testing.T) {
	element := KeyValueElement{"key", "value"}

	dictionary := NewEmptyDictionary()
	error := dictionary.AddKeyValueElement(element)

	if error != nil {
		test.Error("Unexpected error adding a KeyValueElement")
	}

	if len(dictionary.elements) != 1 || dictionary.elements[element.Key] != element.Value {
		test.Error("Wrong behaviour adding a KeyValueElement")
	}

	error = dictionary.AddKeyValueElement(element)

	if error == nil {
		test.Error("Duplicated keys should return an error on AddKeyValueElement method")
	}
}

func TestAddRangeMethod(test *testing.T) { /* @TODO */ }

func TestContainsMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}
	inexistentElement := KeyValueElement{"3Key", "3Value"}

	dictionary := NewEmptyDictionary()
	dictionary.AddKeyValueElement(elementOne)
	dictionary.AddKeyValueElement(elementTwo)

	if !dictionary.Contains(elementOne.Key) {
		test.Error("Contains return a false positive with existent elements")
	}

	if dictionary.Contains(inexistentElement.Key) {
		test.Error("Contains return a false positive with inexistent elements")
	}
}

func TestContainsValueMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}
	inexistentElement := KeyValueElement{"3Key", "3Value"}

	dictionary := NewEmptyDictionary()
	dictionary.AddKeyValueElement(elementOne)
	dictionary.AddKeyValueElement(elementTwo)

	if !dictionary.ContainsValue(elementOne.Value) {
		test.Error("ContainsValue return a false positive with existent elements")
	}

	if dictionary.ContainsValue(inexistentElement.Value) {
		test.Error("ContainsValue return a false positive with inexistent elements")
	}
}

func TestGetMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}

	dictionary := NewEmptyDictionary()
	dictionary.AddKeyValueElement(elementOne)
	dictionary.AddKeyValueElement(elementTwo)

	firstElement := dictionary.Get()

	if firstElement != elementOne {
		test.Error("Wrong extracted element on Get method")
	}

	if len(dictionary.elements) != 1 {
		test.Error("Wrong remained elements in the collection on Get method")
	}
}

func TestFirstMethod(test *testing.T) { /* @TODO */ }

func TestLastMethod(test *testing.T) { /* @TODO */ }

func TestElementsMethod(test *testing.T) { /* @TODO */ }

func TestGetKeysMethod(test *testing.T) { /* @TODO */ }

func TestGetValuesMethod(test *testing.T) { /* @TODO */ }

func TestCountMethod(test *testing.T) {
	dictionary := NewEmptyDictionary()

	if dictionary.Count() != 0 {
		test.Error("Count method returns wrong size of collection when it's empty")
	}

	dictionary.Add("key", "value")

	if dictionary.Count() == 0 {
		test.Error("Count method returns 0 size when collection has elements")
	}
}

func TestEmptyMethod(test *testing.T) {
	dictionary := NewEmptyDictionary()

	if !dictionary.IsEmpty() {
		test.Error("Empty method returns true when it's really empty")
	}

	dictionary.Add("key", "value")

	if dictionary.IsEmpty() {
		test.Error("Empty method returns false when it's not really empty")
	}
}

func TestNewEmptyDictionary(test *testing.T) {
	emptyDictionary := NewEmptyDictionary()

	if len(emptyDictionary.elements) != 0 {
		test.Error("Empty dictionary must to be instancied with no elements")
	}
}

func TestNewDictionary(test *testing.T) { /* TODO */ }
