package main

import (
	"exerciser/internal/cli"
	"exerciser/internal/service"
	"exerciser/internal/storage/dictionary"
	result "exerciser/internal/storage/results"
	"fmt"
)

func main() {
	dictionary := dictionary.Create()
	results := result.NewStorage()
	service := service.NewWordStorage(dictionary, &results)

	for {
		if isEnded, err := cli.Menu(*service); err != nil {
			fmt.Println(err)
		} else {
			if isEnded {
				break
			}
		}
	}
	fmt.Println("You left the program")
}
