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
	keyDefinition   reflect.Type
	valueDefinition reflect.Type
	elements        KeyValueList
	lasKeytElement  *KeyElement
}

// Adds a key-value element to the dictionary
// Dictionary must to be homogeneous in key and in value as well so the specified elements
// should be the same type such the other elements already stored in the dictionary.
// If the dictionary is empty and have no elements, it will take the type of
// the first element as type definition
func (dic *Dictionary) Add(key KeyElement, value ValueElement) error {
	if dic.IsEmpty() {
		dic.keyDefinition = reflect.TypeOf(key)
		dic.valueDefinition = reflect.TypeOf(value)
	}

	if !dic.isHomogeneousWith(key, value) {
		return NewInvalidKeyValueElementTypeError(dic.keyDefinition.Name(), dic.valueDefinition.Name())
	}

	if dic.Contains(key) {
		return NewDuplicatedKeyError()
	}

	dic.lasKeytElement = &key
	dic.elements[key] = value

	return nil
}

// Adds an composed element KeyValueElement to the dictionary
func (dic *Dictionary) AddKeyValueElement(element KeyValueElement) error {
	return dic.Add(element.Key, element.Value)
}

// Inserts a range (slice) of KeyValueElement inside the dictionary
// If the parameter can't be converted to a iterable data type it's return an error
func (dic *Dictionary) AddRange(elements []KeyValueElement) error {
	for _, element := range elements {
		error := dic.AddKeyValueElement(element)

		if error != nil {
			return error
		}
	}

	return nil
}

// Checks if the specified key element is already existing in the dictionary
func (dic *Dictionary) Contains(element KeyElement) bool {
	_, exists := dic.elements[element]

	return exists
}

// Checks if the specified value element exists in the dictionary
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
	delete(dic.elements, element.Key)

	return element
}

// Returns the first element without removing it from the collection
func (dic *Dictionary) First() KeyValueElement {
	for key, value := range dic.elements {
		return KeyValueElement{key, value}
	}

	return KeyValueElement{}
}

// Returns the last element without removing it from the dictionary
func (dic *Dictionary) Last() KeyValueElement {
	return KeyValueElement{dic.lasKeytElement, dic.elements[dic.lasKeytElement]}
}

// Returns the specified key element in the dictionary
func (dic *Dictionary) Element(key KeyElement) KeyValueElement {
	return KeyValueElement{key, dic.elements[key]}
}

// Returns the stored elements as slice of this elements
// This is the proper way to iterate over all the elements inside de dicionary
// treating them as a normal range
func (dic *Dictionary) Elements() KeyValueList {
	return dic.elements
}

// Returns all the keys in the dicionary as a list of KeyElement
func (dic *Dictionary) GetKeys() []KeyElement {
	keys := []KeyElement{}

	for current, _ := range dic.elements {
		keys = append(keys, current)
	}

	return keys
}

// Returns all the values in the dicionary as a list of ValueElement
func (dic *Dictionary) GetValues() []ValueElement {
	values := []ValueElement{}

	for _, current := range dic.elements {
		values = append(values, current)
	}

	return values
}

// Returns the number of elements inside the dicionary
func (dic *Dictionary) Count() int {
	return len(dic.elements)
}

// Checks if the dictionary is empty or not
func (dic *Dictionary) IsEmpty() bool {
	return dic.Count() == 0
}

func (dic *Dictionary) isHomogeneousWith(key KeyElement, value ValueElement) bool {
	return dic.keyDefinition == reflect.TypeOf(key) &&
		dic.valueDefinition == reflect.TypeOf(value)
}

// Instances a new empty dictionary
func NewEmptyDictionary() *Dictionary {
	dic := new(Dictionary)
	dic.elements = make(KeyValueList)

	return dic
}

// This method allows to instance a new Dictionary with a group of key-value elements
func NewDictionary(elements []KeyValueElement) *Dictionary {
	dictionary := NewEmptyDictionary()
	dictionary.AddRange(elements)

	return dictionary
}
