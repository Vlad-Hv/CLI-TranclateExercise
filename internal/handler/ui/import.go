package ui

import (
	"bufio"
	"fmt"
)

func NewPair(in *bufio.Scanner) (string, string, error) {
	var translate string

	askWord()
	word, err := readLine(in)
	if err != nil {
		return "", "", fmt.Errorf("Cannot get word with translate: %w", err)
	}

	if err := validateImport(word); err != nil {
		return "", "", fmt.Errorf("Cannot get word with translate: %w", err)
	}

	askTranslate()
	translate, err = readLine(in)
	if err != nil {
		return "", "", fmt.Errorf("Cannot get word with translate: %w", err)
	}
	if err := validateImport(translate); err != nil {
		return "", "", fmt.Errorf("Cannot get word with translate: %w", err)
	}

	return word, translate, nil
}

func DeleteWord(in *bufio.Scanner) (string, error) {
	askWord()
	word, err := readLine(in)
	if err != nil {
		return "", err
	}

	if err := validateImport(word); err != nil {
		return "", fmt.Errorf("Cannot delete word: %w", err)
	}
	return word, nil
}

func MenuOption(in *bufio.Scanner) (int, error) {
	menu()
	option, err := readInt(in)
	if err != nil {
		return 0, err
	}

	if err = validateMenu(option, err); err != nil {
		return 0, err
	}

	return option, err
}

func DictionaryMenu(in *bufio.Scanner) (int, error) {
	dictionaryMenu()
	option, err := readInt(in)
	if err != nil {
		return 0, err
	}
	if err = validateMenu(option, err); err != nil {
		return 0, err
	}

	return option, err
}

func GetAmoutWords(in *bufio.Scanner) (int, error) {
	amountWord()
	amount, err := readInt(in)
	if err != nil {
		return 0, err
	}
	if err := validateMenu(amount, err); err != nil {
		return 0, err
	}

	return amount, nil
}

func GetTranslate(word string, in *bufio.Scanner) (string, error) {
	translate(word)
	t, err := readLine(in)

	if err != nil {
		return "", err
	}
	return t, nil
}
