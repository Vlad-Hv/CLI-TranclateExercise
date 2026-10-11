package ui

import (
	"errors"
)

func validateImport(input string) error {
	if input == "" {
		return errors.New("empty input")
	}
	return nil
}
