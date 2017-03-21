// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package datatypes/dictionary provides an easy dictionary (key => value) homogeneous
// struct management, making the iteration of a unique-key lists more powerfull,
// simple and clean, accepting primitives types and complex user structs as well.

// This part of package contains the core behaviour

package dictionary

import "reflect"

// Represents Key in Key-Value object
type KeyElement interface{}

// Represents Value in Key-Value object
type ValueElement interface{}

// Key-Value object
type KeyValueElement struct {
	Key   KeyElement
	Value ValueElement
}

// List of Key-Value elements
type KeyValueList []KeyValueElement

// Map of Key-Vale elements
type KeyValueMap map[KeyElement]ValueElement

// Easy dictionary (key => value) struct
type Dictionary struct {
	keyDefinition   reflect.Type
	valueDefinition reflect.Type
	elements        KeyValueMap
}

// Adds a key-value element to the dictionary
// Dictionary must to be homogeneous in key and in value as well so the specified elements
// should be the same type such the other elements already stored in the dictionary.
// If the dictionary is empty and have no elements, it will take the type of
// the first element as type definition
func (dic *Dictionary) Add(key KeyElement, value ValueElement) {
	if dic.IsEmpty() {
		dic.keyDefinition = reflect.TypeOf(key)
		dic.valueDefinition = reflect.TypeOf(value)
	}

	if !dic.isHomogeneousWith(key, value) {
		NewInvalidKeyValueElementTypeError(dic.keyDefinition.Name(), dic.valueDefinition.Name())
	}

	if dic.Contains(key) {
		NewDuplicatedKeyError()
	}

	dic.elements[key] = value
}

// Adds an composed element KeyValueElement to the dictionary
func (dic *Dictionary) AddKeyValueElement(element KeyValueElement) {
	dic.Add(element.Key, element.Value)
}

// Inserts a range (slice) of KeyValueElement inside the dictionary
// If the parameter can't be converted to a iterable data type it's return an error
func (dic *Dictionary) AddRange(elements KeyValueList) {
	for _, element := range elements {
		dic.AddKeyValueElement(element)
	}
}

// Returns the specified key element in the dictionary
func (dic *Dictionary) Element(key KeyElement) KeyValueElement {
	return KeyValueElement{key, dic.elements[key]}
}

// Returns the stored elements as slice of this elements
// This is the proper way to iterate over all the elements inside de dicionary
// treating them as a normal range
func (dic *Dictionary) Elements() KeyValueMap {
	return dic.elements
}

// Returns all the keys in the dicionary as a list of KeyElement
func (dic *Dictionary) Keys() []KeyElement {
	keys := []KeyElement{}

	for current, _ := range dic.elements {
		keys = append(keys, current)
	}

	return keys
}

// Returns all the values in the dicionary as a list of ValueElement
func (dic *Dictionary) Values() []ValueElement {
	values := []ValueElement{}

	for _, current := range dic.elements {
		values = append(values, current)
	}

	return values
}

// Extract the first element and return it
// Keep in mind that this method will modify the dictionary elements substracting that element
func (dic *Dictionary) Extract() KeyValueElement {
	for key, value := range dic.elements {
		dic.Delete(key)

		return KeyValueElement{key, value}
	}

	NewEmptyDictionaryErrorString()

	return KeyValueElement{}
}

// Extract the specified key element and return it
// Keep in mind that this method will modify the dictionary elements substracting that element
func (dic *Dictionary) ExtractKey(key KeyElement) KeyValueElement {
	element := KeyValueElement{key, dic.elements[key]}
	dic.Delete(key)

	return element
}

// Sets a new value for a specified index element
func (dic *Dictionary) Set(key KeyElement, value ValueElement) {
	if !dic.isHomogeneousWith(key, value) {
		NewInvalidKeyValueElementTypeError(dic.keyDefinition.Name(), dic.valueDefinition.Name())
	}

	if !dic.Contains(key) {
		NewElementNotFoundError()
	}

	dic.elements[key] = value
}

// Removes an specified already stored element
// If it's not found the method will return an error
func (dic *Dictionary) Delete(key KeyElement) {
	if !dic.Contains(key) {
		NewElementNotFoundError()
	}

	delete(dic.elements, key)
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

// Returns the number of elements inside the dicionary
func (dic *Dictionary) Size() int {
	return len(dic.elements)
}

// Checks if the dictionary is empty or not
func (dic *Dictionary) IsEmpty() bool {
	return dic.Size() == 0
}

func (dic *Dictionary) isHomogeneousWith(key KeyElement, value ValueElement) bool {
	return dic.keyDefinition == reflect.TypeOf(key) &&
		dic.valueDefinition == reflect.TypeOf(value)
}

// Instances a new empty dictionary
func NewEmptyDictionary() *Dictionary {
	dic := new(Dictionary)
	dic.elements = make(KeyValueMap)

	return dic
}

// This method allows to instance a new Dictionary with a group of key-value elements
func NewDictionary(elements KeyValueList) *Dictionary {
	dictionary := NewEmptyDictionary()
	dictionary.AddRange(elements)

	return dictionary
}
