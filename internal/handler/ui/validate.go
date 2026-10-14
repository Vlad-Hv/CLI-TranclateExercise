package ui

import (
	"errors"
)

var (
	errInputEmpty    = errors.New("empty input")
	errTypeInvalid   = errors.New("invalid option type")
	errOptionInvalid = errors.New("invalid option")
)

func validateImport(input string) error {
	if input == "" {
		return errInputEmpty
	}
	return nil
}

func validateMenu(option int, err error) error {
	if option < 1 {
		return errOptionInvalid
	}

	if err != nil {
		return errTypeInvalid
	}

	return nil
}
