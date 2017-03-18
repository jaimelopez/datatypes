package collection

import (
	"datatypes/generic"
	"reflect"
)

type Element interface{}
type CollectionElements interface{}

type Collection struct {
	definition reflect.Type
	elements   []Element
}

// Adds a single element to the collection
// The collection must to be homogeneous so the specified element
// should be the same type such the other elements already stored in the collection.
// If the collection is empty and has no elements, it will take the type of
// that element as type definition for the collection
func (col *Collection) Add(element Element) {
	if col.IsEmpty() {
		col.definition = reflect.TypeOf(element)
	}

	if !col.isHomogeneousWith(element) {
		NewInvalidElementTypeError(col.definition.Name())
	}

	if col.Contains(element) {
		NewDuplicatedElementError()
	}

	col.elements = append(col.elements, element)
}

// Inserts a range (slice) inside the collection
// If the parameter can't be converted to a iterable data type it's return an error
func (col *Collection) AddRange(elements CollectionElements) {
	for _, element := range generic.ToSlice(elements) {
		col.Add(element)
	}
}

// Adds the elements contained in the parameter collection inside the instanced collection
// If the parameter can't be converted to a iterable data type it's return an error
func (col *Collection) AddCollection(collection *Collection) {
	col.AddRange(collection.elements)
}

// Returns the first element without removing it from the collection
func (col *Collection) First() Element {
	return col.elements[0]
}

// Returns the last element without removing it from the collection
func (col *Collection) Last() Element {
	return col.elements[len(col.elements)-1]
}

// Although a collection is an unsorted data structure list and the position
// of the elements could be changed, this method allows to return an specific index position.
// Be aware that the order of elements could be changed constantly such it's described before
func (col *Collection) ElementAt(position int) Element {
	return col.elements[position]
}

// Returns the stored collection elements as slice of this elements
// This is the proper way to iterate over all the elements of the collection
// treating them as a normal range
func (col *Collection) Elements() []Element {
	return col.elements
}

// Extract the first element and return it
// Keep in mind that this method will modify the collection elements substracting that element
func (col *Collection) Extract() Element {
	element := col.First()
	col.elements = col.elements[1:]

	return element
}

// Sets a new value for a specified index element
func (col *Collection) Set(position int, element Element) {
	if !col.isHomogeneousWith(element) {
		NewInvalidElementTypeError(col.definition.Name())
	}

	col.elements[position] = element
}

// Removes an specified already stored element
// If it's not found the method will return an error
func (col *Collection) Delete(element Element) {
	if !col.isHomogeneousWith(element) {
		NewInvalidElementTypeError(col.definition.Name())
	}

	for index, current := range col.elements {
		if reflect.DeepEqual(current, element) {
			col.elements[index] = col.elements[col.Count()-1]
			col.elements = col.elements[:col.Count()-1]

			return
		}
	}

	NewElementNotFoundError()
}

// Removes all the found elements contained in the specified range (slice)
// If the parameter can't be converted to a iterable data type it's return an error
func (col *Collection) DeleteRange(elements CollectionElements) {
	for _, element := range generic.ToSlice(elements) {
		col.Delete(element)
	}
}

// Removes all the found elements contained in the specified collection from the instaced collection
// If the parameter can't be converted to a iterable data type it's return an error
func (col *Collection) DeleteCollection(collection *Collection) {
	col.DeleteRange(collection.elements)
}

// Checks if the specified element is already existing in the collection
func (col *Collection) Contains(element Element) bool {
	for _, iterator := range col.elements {
		if reflect.DeepEqual(iterator, element) {
			return true
		}
	}

	return false
}

// Checks if any of the parameter elements there are already contained in the collection
func (col *Collection) ContainsAny(elements CollectionElements) (result bool) {
	defer func() {
		if recover() != nil {
			result = false
		}
	}()

	for _, element := range generic.ToSlice(elements) {
		if col.Contains(element) {
			return true
		}
	}

	return
}

// Returns the number of elements inside the collection
func (col *Collection) Count() int {
	return len(col.elements)
}

// Checks if the collection is empty or not
func (col *Collection) IsEmpty() bool {
	return col.Count() == 0
}

func (col *Collection) isHomogeneousWith(element Element) bool {
	return col.definition == reflect.TypeOf(element)
}

// Instances a new empty collection
func NewEmptyCollection() *Collection {
	return new(Collection)
}

// This method allows to instance a new Collection with a group of elements
// It accepts an enumerable
func NewCollection(elements CollectionElements) (collection *Collection) {
	collection = new(Collection)

	defer func(collection *Collection) {
		if recover() != nil {
			collection.Add(elements)
		}
	}(collection)

	collection.AddRange(elements)

	return
}
