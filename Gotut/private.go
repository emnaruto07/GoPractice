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
		inputs := append(inputs, text)

		if text == "" {
			break
		}
		// _ = input
	}

	fmt.Println(inputs)
}

// go run private.go
// take words as input and count the occurence of each word
// hi, hello, hi, bye, hello
// hi: 2, hello: 2, bye: 1
