package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	inputs := make([]string, 0)

	for scanner.Scan() {
		text := scanner.Text()
		input := append(inputs, text)

		if text == "" {
			break
		}
		_ = input
	}

	fmt.Println(inputs)
}

// go run private.go