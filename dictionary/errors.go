package dictionary

import "errors"

const DuplicatedKeyErrorString = "Duplicated key in dictionary"

func NewDuplicatedKeyError() error {
	return errors.New(DuplicatedKeyErrorString)
}
