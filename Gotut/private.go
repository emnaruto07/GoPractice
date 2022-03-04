package main

import (
	"fmt"
	"bufio"
	"os"
)

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	inputs := make([]string, 0)

// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		inputs := append(inputs, text)

// 		if text == "" {
// 			break
// 		}
// 	}

// 	fmt.Println(inputs)
// }

// func fn() {
// 	ok := isLifeOkay()

// 	if !ok {
// 		makeLifeBetter()
// 		return // bad path, ends
// 	}

// 	// happy path
// 	yes = haveMoney()
// 	if !yes {
// 		makeMoney()
// 		return // early exit
// 	}

// 	// happy path
// 	rejoice()
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	m := make(map[string]int)

	for scanner.Scan() {
		word := scanner.Text()
		if word == "" {
			break
		}

		_, ok := m[word]
		if !ok { // we are looking at this word for the first time
			m[word] = 1
			continue
		} 
		//happy path
		m[word] += 1
		
	}
	fmt.Println(m)
}

//map[:1 bye:1 hello:2 hi:2]

// 2min bhai
// input: frist, hi
// map: hi -> 1, 

// go run private.go
// take words as input and count the occurence of each word
// hi, hello, hi, bye, hello
// hi: 2, hello: 2, bye: 1
