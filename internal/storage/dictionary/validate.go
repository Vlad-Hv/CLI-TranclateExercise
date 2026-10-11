package dictionary

import (
	"errors"
)

func (d Dictionary) ValidateExist(word string) error {
	if _, ok := d[word]; ok {
		return errors.New("word already exist")
	}

	return nil
}
