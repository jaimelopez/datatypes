package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMethod(test *testing.T) {
	keyElement := "key"
	valueElement := "value"

	dictionary := NewEmptyDictionary()

	assert.NotPanics(test, func() {
		dictionary.Add(keyElement, valueElement)
	}, "Unexpected error adding a element")

	assert.Len(test, dictionary.elements, 1, "Wrong behaviour adding a element")
	assert.Exactly(test, dictionary.elements[keyElement], valueElement, "Wrong behaviour adding a element")

	assert.Panics(test, func() {
		dictionary.Add(keyElement, valueElement)
	}, "Duplicated keys should return an error on Add method")

	differentTypeKeyElement, differentTypeValueElement := 1, 2

	assert.Panics(test, func() {
		dictionary.Add(differentTypeKeyElement, differentTypeValueElement)
	}, "Add method should throw an exception trying to insert a non-homogeneous elements")
}

func TestAddKeyValueElementMethod(test *testing.T) {
	element := KeyValueElement{"key", "value"}

	dictionary := NewEmptyDictionary()

	assert.NotPanics(test, func() {
		dictionary.AddKeyValueElement(element)
	}, "Unexpected error adding a KeyValueElement")

	assert.Len(test, dictionary.elements, 1, "Wrong behaviour adding a KeyValueElement")
	assert.Exactly(test, dictionary.elements[element.Key], element.Value, "Wrong behaviour adding a KeyValueElement")

	assert.Panics(test, func() {
		dictionary.AddKeyValueElement(element)
	}, "Duplicated keys should return an error on AddKeyValueElement method")
}

func TestAddRangeMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}
	elementList := KeyValueList{elementOne, elementTwo}

	dictionay := NewEmptyDictionary()

	assert.NotPanics(test, func() {
		dictionay.AddRange(elementList)
	}, "Wrong error returned adding a element range")

	assert.Len(test, dictionay.elements, len(elementList), "Wrong behaviour adding a element range")
}

func TestFirstMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}

	dictionary := NewDictionary(KeyValueList{elementOne, elementTwo})

	first := dictionary.First()

	assert.Exactly(test, first.Key, elementOne.Key, "First method do not return the correct element")
	assert.Exactly(test, first.Value, elementOne.Value, "First method do not return the correct element")
}

func TestLastMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}
	elementThree := KeyValueElement{"3Key", "3Value"}

	dictionary := NewDictionary(KeyValueList{
		elementOne,
		elementTwo,
		elementThree,
	})

	last := dictionary.Last()

	assert.Exactly(test, last.Key, elementThree.Key, "Last method do not return the correct element")
	assert.Exactly(test, last.Value, elementThree.Value, "Last method do not return the correct element")
}

func TestElementMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}

	dictionary := NewDictionary(KeyValueList{elementOne, elementTwo})

	assert.Exactly(test, dictionary.Element(elementOne.Key), elementOne, "Wrong returned element in specific position on Element method")
	assert.Exactly(test, dictionary.Element(elementTwo.Key), elementTwo, "Wrong returned element in specific position on Element method")
}

func TestElementsMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}

	dictionary := NewEmptyDictionary()

	assert.Empty(test, dictionary.Elements(), "Elements method should return no elements on new empty instance")

	dictionary.AddKeyValueElement(elementOne)
	dictionary.AddKeyValueElement(elementTwo)

	assert.NotEmpty(test, dictionary.Elements(), "Elements method do not return the correct stored elements in the collection")
	assert.Equal(test, dictionary.Elements()[elementOne.Key], elementOne.Value, "Elements method do not return the correct stored elements in the collection")
	assert.Equal(test, dictionary.Elements()[elementTwo.Key], elementTwo.Value, "Elements method do not return the correct stored elements in the collection")
}

func TestKeysMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}

	dictionary := NewEmptyDictionary()
	dictionary.AddKeyValueElement(elementOne)
	dictionary.AddKeyValueElement(elementTwo)

	keys := dictionary.Keys()

	assert.Exactly(test, keys[0], elementOne.Key, "Wrong extracted elements on Keys method")
	assert.Exactly(test, keys[1], elementTwo.Key, "Wrong extracted elements on Keys method")
	assert.Len(test, keys, dictionary.Count(), "Wrong get value elements in the dictionary on Keys method")
}

func TestValuesMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}

	dictionary := NewEmptyDictionary()
	dictionary.AddKeyValueElement(elementOne)
	dictionary.AddKeyValueElement(elementTwo)

	values := dictionary.Values()

	assert.Exactly(test, values[0], elementOne.Value, "Wrong extracted elements on Value method")
	assert.Exactly(test, values[1], elementTwo.Value, "Wrong extracted elements on Value method")

	assert.Len(test, values, dictionary.Count(), "Wrong get value elements in the dictionary on Values method")
}

func TestExtractMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}

	dictionary := NewEmptyDictionary()
	dictionary.AddKeyValueElement(elementOne)
	dictionary.AddKeyValueElement(elementTwo)

	extracted := dictionary.Extract()

	assert.Exactly(test, extracted, elementOne, "Wrong extracted element on Extract method")
	assert.Len(test, dictionary.elements, 1, "Wrong remained elements in the collection on Extract method")
}

func TestExtractKeyMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}

	dictionary := NewDictionary(KeyValueList{elementOne, elementTwo})

	extracted := dictionary.ExtractKey("2Key")

	assert.Exactly(test, extracted, elementTwo, "Wrong extracted element on ExtractKey method")
	assert.Len(test, dictionary.elements, 1, "Wrong remained elements in the collection on ExtractKey method")
}

func TestSetMethod(test *testing.T) {
	elementOne := KeyValueElement{"key1", "value1"}
	elementTwo := KeyValueElement{"key2", "value2"}
	elementTwoNewValue := "newValue2"

	dictionary := NewDictionary(KeyValueList{elementOne, elementTwo})
	dictionary.Set(elementTwo.Key, elementTwoNewValue)

	newElement := dictionary.Element(elementTwo.Key)

	assert.Exactly(test, newElement.Value, elementTwoNewValue, "Set method doesn't works properly")
	assert.NotEqual(test, newElement, elementTwo, "Set method doesn't works properly")
	assert.Len(test, dictionary.elements, 2, "Set method doesn't mantains the right items")
}

func TestDeleteMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}
	elementThree := KeyValueElement{"3Key", "3Value"}

	dictionary := NewDictionary(KeyValueList{
		elementOne,
		elementTwo,
		elementThree,
	})

	assert.NotPanics(test, func() {
		dictionary.Delete(elementTwo.Key)
	}, "Unexpected error delenting an element")

	assert.Len(test, dictionary.elements, 2, "Invalid number of elements after a element deletion")
	assert.Equal(test, dictionary.Elements()[elementOne.Key], elementOne.Value, "Invalid expected elements after a single element deletion")
	assert.Equal(test, dictionary.Elements()[elementThree.Key], elementThree.Value, "Invalid expected elements after a single element deletion")
}

func TestContainsMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}
	inexistentElement := KeyValueElement{"3Key", "3Value"}

	dictionary := NewEmptyDictionary()
	dictionary.AddKeyValueElement(elementOne)
	dictionary.AddKeyValueElement(elementTwo)

	assert.True(test, dictionary.Contains(elementOne.Key), "Contains return a false positive with existent elements")
	assert.False(test, dictionary.Contains(inexistentElement.Key), "Contains return a false positive with inexistent elements")
}

func TestContainsValueMethod(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}
	inexistentElement := KeyValueElement{"3Key", "3Value"}

	dictionary := NewEmptyDictionary()
	dictionary.AddKeyValueElement(elementOne)
	dictionary.AddKeyValueElement(elementTwo)

	assert.True(test, dictionary.ContainsValue(elementOne.Value), "ContainsValue return a false positive with existent elements")
	assert.False(test, dictionary.ContainsValue(inexistentElement.Value), "ContainsValue return a false positive with inexistent elements")
}

func TestCountMethod(test *testing.T) {
	dictionary := NewEmptyDictionary()

	assert.Zero(test, dictionary.Count(), "Count method returns wrong size of collection when it's empty")

	dictionary.Add("key", "value")

	assert.NotZero(test, dictionary.Count(), "Count method returns 0 size when collection has elements")
}

func TestIsEmptyMethod(test *testing.T) {
	dictionary := NewEmptyDictionary()

	assert.True(test, dictionary.IsEmpty(), "Empty method returns true when it's really empty")

	dictionary.Add("key", "value")

	assert.False(test, dictionary.IsEmpty(), "Empty method returns false when it's not really empty")
}

func TestNewEmptyDictionary(test *testing.T) {
	emptyDictionary := NewEmptyDictionary()

	assert.Empty(test, emptyDictionary.elements, "Empty dictionary must to be instancied with no elements")
}

func TestNewDictionary(test *testing.T) {
	elementOne := KeyValueElement{"1Key", "1Value"}
	elementTwo := KeyValueElement{"2Key", "2Value"}
	elementList := KeyValueList{elementOne, elementTwo}

	dictionary := NewDictionary(elementList)

	assert.Len(test, dictionary.elements, len(elementList), "New dictionary don't store elements parameters as elements")
	assert.Equal(test, dictionary.Elements()[elementOne.Key], elementOne.Value, "New dictionary don't store elements parameters as elements")
	assert.Equal(test, dictionary.Elements()[elementTwo.Key], elementTwo.Value, "New dictionary don't store elements parameters as elements")
}
