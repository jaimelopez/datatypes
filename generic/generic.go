package generic

import "reflect"

func ToSlice(slice interface{}) []interface{} {
	sliceValues := reflect.ValueOf(slice)

	if sliceValues.Kind() != reflect.Slice {
		NewInvalidIterableElementError()
	}

	values := make([]interface{}, sliceValues.Len())

	for x := 0; x < sliceValues.Len(); x++ {
		values[x] = sliceValues.Index(x).Interface()
	}

	return values
}

func AreSameType(first interface{}, second interface{}) bool {
	return reflect.TypeOf(first) == reflect.TypeOf(second)
}
