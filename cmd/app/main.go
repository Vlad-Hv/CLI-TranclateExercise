package main

import (
	"bufio"
	"errors"
	"exerciser/internal/cli"
	"exerciser/internal/service"
	"exerciser/internal/storage/dictionary"
	result "exerciser/internal/storage/results"
	"fmt"
	"io"
	"os"
)

func main() {
	dictionary := dictionary.Create()
	results := result.NewStorage()
	service := service.NewWordStorage(dictionary, &results)
	in := bufio.NewScanner(os.Stdin)

	for {
		if isEnded, err := cli.Menu(*service, in); err != nil {
			fmt.Println(err)
			if errors.Is(err, io.EOF) || in.Err() != nil {
				break
			}
		} else {
			if isEnded {
				break
			}
		}
	}
	fmt.Println("You left the program")
}
