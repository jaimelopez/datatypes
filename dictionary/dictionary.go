// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package datatypes/dictionary provides an easy dictionary (key => value) homogeneous
// struct management, making the iteration of a unique-key lists more powerful,
// simple and clean, accepting primitives types and complex user structs as well.

// This part of package contains the core behaviour

package dictionary

import "reflect"

// KeyElement represents Key in Key-Value object
type KeyElement interface{}

// ValueElement represents Value in Key-Value object
type ValueElement interface{}

// KeyValueElement represents a Key-Value object
type KeyValueElement struct {
	Key   KeyElement
	Value ValueElement
}

// KeyValueMap represents a map of Key-Vale elements
type KeyValueMap map[KeyElement]ValueElement

// Dictionary represents a simple dictionary (key => value) struct
type Dictionary struct {
	keyDefinition   reflect.Type
	valueDefinition reflect.Type
	elements        KeyValueMap
}

// Add a key-value element to the dictionary
// Dictionary must to be homogeneous in key and in value as well so the specified elements
// should be the same type such the other elements already stored in the dictionary.
// If the dictionary is empty and have no elements, it will take the type of
// the first element as type definition
func (dic *Dictionary) Add(key KeyElement, value ValueElement) error {
	if dic.IsEmpty() {
		dic.keyDefinition = reflect.TypeOf(key)
		dic.valueDefinition = reflect.TypeOf(value)
	} else if !dic.isHomogeneousWith(key, value) {
		return ErrInvalidKeyValueElementType
	}

	if dic.Contains(key) {
		return ErrDuplicatedKey
	}

	dic.elements[key] = value

	return nil
}

// AddKeyValueElement adds an composed element KeyValueElement to the dictionary
func (dic *Dictionary) AddKeyValueElement(element KeyValueElement) error {
	return dic.Add(element.Key, element.Value)
}

// AddRange inserts a range (slice) of KeyValueElement inside the dictionary
// If the parameter can't be converted to a iterable data type it's return an error
func (dic *Dictionary) AddRange(elements []KeyValueElement) error {
	for _, element := range elements {
		if err := dic.AddKeyValueElement(element); err != nil {
			return err
		}
	}

	return nil
}

// Element returns the specified key element in the dictionary
func (dic *Dictionary) Element(key KeyElement) (*ValueElement, error) {
	element, exists := dic.elements[key]

	if !exists {
		return nil, ErrElementNotFound
	}

	return &element, nil
}

// Elements returns the stored elements as slice of this elements
// This is the proper way to iterate over all the elements inside de dicionary
// treating them as a normal range
func (dic *Dictionary) Elements() *KeyValueMap {
	return &dic.elements
}

// Keys returns all the keys in the dicionary as a list of KeyElement
func (dic *Dictionary) Keys() []KeyElement {
	keys := []KeyElement{}

	for current := range dic.elements {
		keys = append(keys, current)
	}

	return keys
}

// Values returns all the values in the dicionary as a list of ValueElement
func (dic *Dictionary) Values() []ValueElement {
	values := []ValueElement{}

	for _, current := range dic.elements {
		values = append(values, current)
	}

	return values
}

// Extract the first element and return it
// Keep in mind that this method will modify the dictionary elements subtracting that element
func (dic *Dictionary) Extract() *KeyValueElement {
	for key, value := range dic.elements {
		dic.Delete(key)

		return &KeyValueElement{key, value}
	}

	return nil
}

// ExtractKey extracts the specified key element and return it
// Keep in mind that this method will modify the dictionary elements subtracting that element
func (dic *Dictionary) ExtractKey(key KeyElement) (*KeyValueElement, error) {
	if dic.Contains(key) {
		element := &KeyValueElement{key, dic.elements[key]}
		dic.Delete(key)

		return element, nil
	}

	return nil, ErrElementNotFound
}

// Set a new value for a specified index element
func (dic *Dictionary) Set(key KeyElement, value ValueElement) error {
	if !dic.isHomogeneousWith(key, value) {
		return ErrInvalidKeyValueElementType
	}

	if !dic.Contains(key) {
		return ErrElementNotFound
	}

	dic.elements[key] = value

	return nil
}

// Delete an specified already stored element
// If it's not found the method will return an error
func (dic *Dictionary) Delete(key KeyElement) error {
	if !dic.Contains(key) {
		return ErrElementNotFound
	}

	delete(dic.elements, key)

	return nil
}

// Contains checks if the specified key element is already existing in the dictionary
func (dic *Dictionary) Contains(element KeyElement) bool {
	_, exists := dic.elements[element]

	return exists
}

// ContainsValue checks if the specified value element exists in the dictionary
func (dic *Dictionary) ContainsValue(element ValueElement) bool {
	for _, value := range dic.elements {
		if reflect.DeepEqual(value, element) {
			return true
		}
	}

	return false
}

// Filter returns a element colecction filtering the elements with a function
// If the functions return true the element will be filtered
func (dic *Dictionary) Filter(f func(KeyValueElement) bool) *KeyValueMap {
	results := make(KeyValueMap)

	for key, value := range *dic.Elements() {
		elem := KeyValueElement{key, value}

		if !f(elem) {
			continue
		}

		results[key] = value
	}

	return &results
}

// Size returns the number of elements inside the dicionary
func (dic *Dictionary) Size() int {
	return len(dic.elements)
}

// IsEmpty checks if the dictionary is empty or not
func (dic *Dictionary) IsEmpty() bool {
	return dic.Size() == 0
}

func (dic *Dictionary) isHomogeneousWith(key KeyElement, value ValueElement) bool {
	return dic.keyDefinition == reflect.TypeOf(key) &&
		dic.valueDefinition == reflect.TypeOf(value)
}

// NewEmptyDictionary instances a new empty dictionary
func NewEmptyDictionary() *Dictionary {
	dic := new(Dictionary)
	dic.elements = make(KeyValueMap)

	return dic
}

// NewDictionary allows to instance a new Dictionary with a group of key-value elements
func NewDictionary(elements []KeyValueElement) (*Dictionary, error) {
	dictionary := NewEmptyDictionary()
	err := dictionary.AddRange(elements)

	return dictionary, err
}
