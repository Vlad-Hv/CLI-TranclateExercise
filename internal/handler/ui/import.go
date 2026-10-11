package ui

import (
	"fmt"
)

func Pair() (string, string, error) {
	var word string
	var translate string

	askWords()
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

func askWords() {
	fmt.Print("Enter please Word: ")
}

func askTranslate() {
	fmt.Print("Enter please Translate: ")
}
