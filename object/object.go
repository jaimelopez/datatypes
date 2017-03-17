package object

import "reflect"
import "fmt"

// @TODO
func IsDefault(object interface{}) bool {
	return false
}

func AreEqualType(first interface{}, second interface{}) bool {
	fmt.Println(reflect.TypeOf(first))

	return reflect.TypeOf(first) == reflect.TypeOf(second)
}
