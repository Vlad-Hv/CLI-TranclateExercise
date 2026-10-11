package main

import (
	"exerciser/internal/storage/dictionary"
	"fmt"
)

func main() {
	dictionary := dictionary.Create()

	fmt.Println(dictionary)
}
