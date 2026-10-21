package service

import (
	"errors"
	"strings"
)

var (
	errWordEmpty      = errors.New("Word is empty")
	errTranslateEmpty = errors.New("Translate is empty")
	errAmountInvalid  = errors.New("invalid amount")
)

func validatePair(w, t string) error {
	if strings.TrimSpace(w) == "" {
		return errWordEmpty
	}

	if strings.TrimSpace(t) == "" {
		return errTranslateEmpty
	}

	return nil
}

func ValidateAmount(amount int) error {
	if amount <= 0 {
		return errAmountInvalid
	}

	return nil
}
