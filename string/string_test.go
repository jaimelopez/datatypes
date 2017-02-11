package string

import "testing"

func TestEmptyMethod(test *testing.T) {
	defaultString := ""
	notDefaultString := "  "

	if !IsDefault(defaultString) {
		test.Error("Wrong result in IsDefault method: it should return true with a default string")
	}

	if IsDefault(notDefaultString) {
		test.Error("Wrong result in IsDefault method: it should return false with a non-default string")
	}
}

func TestDefaultMethod(test *testing.T) {
	emptyString := "          "
	notEmptyString := "  string"

	if !IsEmpty(emptyString) {
		test.Error("Wrong result in IsEmpty method: it should return true with an empty string")
	}

	if IsEmpty(notEmptyString) {
		test.Error("Wrong result in IsEmpty method: it should return false with an non-empty string")
	}
}
