package dictionary

import (
	"errors"
)

var (
	ErrWordExist      = errors.New("word already exist")
	ErrWordUnexist    = errors.New("word unexist")
	ErrWordsNotEnough = errors.New("words donot enough")
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

func (d Dictionary) ValidateWordsAmount(amount int) error {
	if len(d) < amount {
		return ErrWordsNotEnough
	}

	return nil
}
