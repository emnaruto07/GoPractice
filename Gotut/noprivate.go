package main

import "fmt"

// func idiot(x int) int { // idiot, identity
// 	return x
// }

// paradigm:
// OOP: Java, c++
// Procedural: C, fortran
// Functional: haskell

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	test := 0

	for _, item := range arr {
		test += item * item
	}

	fmt.Println(test)

}
