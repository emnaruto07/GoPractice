package main

import "fmt"

func accumulator() func(int) int {
	var sum int = 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func cEl(tag string) (func(string) Element) {
	return func(id string) Element {
		return createElement(tag, id)
	}
}

// const cEL = (tag) => (id) => ...

func main() {
	acc := accumulator()

	var cDiv := cEl("div")

	k := acc(3) // 3
	acc(2)      // 5
	acc(8)      // 13 
	fmt.Println(acc(0))
}
