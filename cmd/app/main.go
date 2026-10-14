package main

import (
	"exerciser/internal/cli"
	"exerciser/internal/service"
	"exerciser/internal/storage/dictionary"
	"fmt"
)

func main() {
	dictionary := dictionary.Create()
	service := service.New(dictionary)

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
