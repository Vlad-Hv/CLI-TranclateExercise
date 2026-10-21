package ui

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func readLine(in *bufio.Scanner) (string, error) {
	if !in.Scan() {
		if err := in.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return in.Text(), nil
}

func readInt(in *bufio.Scanner) (int, error) {
	text, err := readLine(in)
	if err != nil {
		return 0, err
	}

	number, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return 0, fmt.Errorf("expected as integer %w", err)
	}

	return number, nil
}
