package dictionary

import "reflect"

type KeyElement interface{}
type ValueElement interface{}

type KeyValueElement struct {
	Key   KeyElement
	Value ValueElement
}

type KeyValueList map[KeyElement]ValueElement

type Dictionary struct {
	elements KeyValueList
}

////////////////////////////////
// TODO : to implement homogeneus checkings

func (dic *Dictionary) Add(key KeyElement, value ValueElement) error {
	if dic.Contains(key) {
		return NewDuplicatedKeyError()
	}

	dic.elements[key] = value

	return nil
}

func (dic *Dictionary) AddKeyValueElement(element KeyValueElement) error {
	return dic.Add(element.Key, element.Value)
}

func (dic *Dictionary) AddRange(elements []KeyValueElement) error {
	for _, element := range elements {
		error := dic.AddKeyValueElement(element)

		if error != nil {
			return error
		}
	}

	return nil
}

func (dic *Dictionary) Contains(element KeyElement) bool {
	for key, _ := range dic.elements {
		if reflect.DeepEqual(key, element) {
			return true
		}
	}

	return false
}

func (dic *Dictionary) ContainsValue(element ValueElement) bool {
	for _, value := range dic.elements {
		if reflect.DeepEqual(value, element) {
			return true
		}
	}

	return false
}

// Extract the first element and return it
// Keep in mind that this method will modify the dictionary elements substracting that element
func (dic *Dictionary) Get() KeyValueElement {
	element := dic.First()
	// dic.elements = dic.elements[1:]

	return element
}

// Returns the first element without removing it from the collection
func (dic *Dictionary) First() KeyValueElement {
	return KeyValueElement{}
}

// Returns the last element without removing it from the collection
func (dic *Dictionary) Last() KeyValueElement {
	return KeyValueElement{}
}

func (dic *Dictionary) Elements() KeyValueList {
	return dic.elements
}

func (dic *Dictionary) GetKeys() []KeyElement {
	return nil
}

func (dic *Dictionary) GetValues() []ValueElement {
	return nil
}

func (dic *Dictionary) Count() int {
	return len(dic.elements)
}

func NewEmptyDictionary() *Dictionary {
	dic := new(Dictionary)
	dic.elements = make(KeyValueList)

	return dic
}

func NewDictionary(elements []KeyValueElement) *Dictionary {
	dictionary := NewEmptyDictionary()
	dictionary.AddRange(elements)

	return dictionary
}
