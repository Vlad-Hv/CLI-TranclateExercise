package ui

import (
	"fmt"
)

func NewPair() (string, string, error) {
	var word string
	var translate string

	askWord()
	fmt.Scanln(&word)
	if err := validateImport(word); err != nil {
		return "", "", fmt.Errorf("Cannot get word with translate: %w", err)
	}

	askTranslate()
	fmt.Scanln(&translate)
	if err := validateImport(translate); err != nil {
		return "", "", fmt.Errorf("Cannot get word with translate: %w", err)
	}

	return word, translate, nil
}

func DeleteWord() (string, error) {
	var word string
	askWord()
	fmt.Scanln(&word)

	if err := validateImport(word); err != nil {
		return "", fmt.Errorf("Cannot delete word: %w", err)
	}
	return word, nil
}

func MenuOption() (int, error) {
	var option int
	menu()
	_, err := fmt.Scanln(&option)

	if err = validateMenu(option, err); err != nil {
		return 0, err
	}

	return option, err
}

func DictionaryMenu() (int, error) {
	var option int
	dictionaryMenu()
	_, err := fmt.Scanln(&option)

	if err = validateMenu(option, err); err != nil {
		return 0, err
	}

	return option, err
}
