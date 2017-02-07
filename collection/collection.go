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

func (col *Collection) Add(element Element) error {
	if col.Contains(element) {
		return NewDuplicatedElementError()
	}

	if col.IsEmpty() {
		col.definition = reflect.TypeOf(element)
	}

	if col.definition != reflect.TypeOf(element) {
		return NewInvalidElementTypeError(col.definition.Name())
	}

	col.elements = append(col.elements, element)

	return nil
}

func (col *Collection) AddRange(elements CollectionElements) error {
	//defer func() {
	//	if recoverer := recover(); recoverer != nil {
	//		return NewInvalidIterableElementError()
	//	}
	//}()

	for _, element := range generic.ToSlice(elements) {
		error := col.Add(element)

		if error != nil {
			return error
		}
	}

	return nil
}

func (col *Collection) AddCollection(collection Collection) error {
	for _, element := range collection.elements {
		error := col.Add(element)

		if error != nil {
			return error
		}
	}

	return nil
}

func (col *Collection) Delete(element Element) error {
	for index, current := range col.elements {
		if reflect.DeepEqual(current, element) {
			col.elements[index] = col.elements[col.Count()-1]
			col.elements = col.elements[:col.Count()-1]

			return nil
		}
	}

	return NewElementNotFoundError()
}

func (col *Collection) DeleteRange(elements CollectionElements) error {
	for _, element := range generic.ToSlice(elements) {
		error := col.Delete(element)

		if error != nil {
			return error
		}
	}

	return nil
}

func (col *Collection) Contains(element Element) bool {
	for _, iterator := range col.elements {
		if reflect.DeepEqual(iterator, element) {
			return true
		}
	}

	return false
}

func (col *Collection) ContainsAny(elements CollectionElements) bool {
	for _, element := range generic.ToSlice(elements) {
		if col.Contains(element) {
			return true
		}
	}

	return false
}

func (col *Collection) First(position int) Element {
	return col.elements[0]
}

func (col *Collection) Last(position int) Element {
	return col.elements[len(col.elements)-1]
}

func (col *Collection) ElementAt(position int) Element {
	return col.elements[position]
}

func (col *Collection) Elements() []Element {
	return col.elements
}

func (col *Collection) Count() int {
	return len(col.elements)
}

func (col *Collection) IsEmpty() bool {
	return col.Count() == 0
}

func NewEmptyCollection() *Collection {
	return new(Collection)
}

func NewCollection(elements CollectionElements) *Collection {
	collection := new(Collection)

	error := collection.AddRange(generic.ToSlice(elements))

	if error != nil {
		panic(error)
	}

	return collection
}
