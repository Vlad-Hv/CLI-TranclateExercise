package ui

import (
	"fmt"
)

func menu() {
	fmt.Println("===Menu===")
	fmt.Println("1. Dictionary\n2. Training session\n3. Check results\n4. Exit")
}

func dictionaryMenu() {
	fmt.Println("===Dictionary Menu===")
	fmt.Println("1. Add word pair\n2. Delete word pair\n3. Check word list \n4. Exit")
}

func AddSuccess() {
	fmt.Println("Added Succesfully")
}

func DeleteSuccess() {
	fmt.Println("Deleted succesfully")
}

func PrintWordList(list map[string]string) {
	wordList()
	for word, translate := range list {
		fmt.Println(word, ":", translate)
	}
}

func wordList() {
	fmt.Println("===Word List===")
}

func askWord() {
	fmt.Print("Enter please Word: ")
}

func askTranslate() {
	fmt.Print("Enter please Translate: ")
}

func LeftMenu() {
	fmt.Println("You left the room")
}

func UnexistOption() {
	fmt.Println("Unexist option")
}
