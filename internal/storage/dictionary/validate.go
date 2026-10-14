package dictionary

import (
	"errors"
)

var (
	ErrWordExist   = errors.New("word already exist")
	ErrWordUnexist = errors.New("word unexist")
)

func (d Dictionary) ValidateExist(word string) error {
	if _, ok := d[word]; ok {
		return ErrWordExist
	}

	return nil
}

func (d Dictionary) ValidateUnexist(word string) error {
	if _, ok := d[word]; !ok {
		return ErrWordUnexist
	}

	return nil
}
