package object

import "reflect"

// @TODO
func IsDefault(object interface{}) bool {
	return false
}

func AreEqualType(first interface{}, second interface{}) bool {
	return reflect.TypeOf(first) == reflect.TypeOf(second)
}
