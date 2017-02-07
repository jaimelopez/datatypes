package string

import "strings"

const Default string = ""

func IsDefault(str string) bool {
	return str == Default
}

func IsEmpty(str string) bool {
	return IsDefault(strings.Trim(str, " "))
}
