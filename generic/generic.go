package generic

import "reflect"

func ToSlice(slice interface{}) []interface{} {
	sliceValues := reflect.ValueOf(slice)

	if sliceValues.Kind() != reflect.Slice {
		panic("Non-iterable type can not be converted to slice")
	}

	values := make([]interface{}, sliceValues.Len())

	for x := 0; x < sliceValues.Len(); x++ {
		values[x] = sliceValues.Index(x).Interface()
	}

	return values
}
