package main

import (
	"fmt"
)

// type TwitterConfig struct {
// 	ApiKey string
// 	ApiSecret string
// 	Hash string
// }

var l;

func blah() {
	defer fmt.Println("blah")

	const dkasldk = "Dawdas"
}


// func generateHash(twc *TwitterConfig) {
	
// 	// twc is pointer(*) to TwitterConfig

// 	//0xidoapisapoi.hash = 0xidoapisapoi.apike
// 	*twc.Hash = *twc.ApiKey + *twc.ApiSecret
// 	(*twc).Hash
// }

// func checkStatus() error {
// 	return nil
// }

type Person struct{
	name string
	age int
}

type UserPool []Person

// ErrorCode the error code returned by the programs
type ErrorCode uint

type Error struct{
	code ErrorCode
	message string
}

type Coord struct{
	x int
	y int
}

var delhiLocation Coord = Coord{
	x: 424423,
	y: 424232,
}
var meerutLocation Coord = Coord{
	x: 424423,
	y: 424232,
}

type Cities []Coord

type Objects []Coord


func main() {
	cities := Cities{delhiLocation, meerutLocation}

	objects := Objects{
		{ 12, 23 },
		{ 34, 45 },
	}

	// len(cities) // 2
	// cap(cities) // 
	
	// slice := []int{ 0, 1, }
	slice = append(slice, 2, 3, 4) // No game No life

	slice := make([]int, 0, 10)

	fmt.Println(slice) // 0, 1	

	//


	path := Path{
		// coord,
		// coord.,
		// coord,
		// ...
	}

	arr := [...]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, }

	// [0 1 2 12 4 5 6 7 8 9]
	//    |       |

  slice := arr[1:6] // [1 2 12 4 5]

	slice[1] // 2	
	slice[2] = 12
	arr[3] == slice[2]

	

	person := Person{ "shazeb", 28 }

	person2 := Person{ "keo", 20 }

	person.
	blah()
	fmt.Println("testing")

	var k; // local
	
	// https://github.com/emnaruto07/beginner-javascript/blob/a8d7bae1c9ffb2b0bf2c38cbc03fe4d60371b09c/main.go#L11
}



/*
  primitives: int, uint, floats, string
	user-defined: `type <name> <definition>`
	[*] Variables
	var test = "hello"
	[*] Constants
	const url= "https://www.example.com"
	[*] Type Inference
	automatic type guessing
	[*] Type Conversions/casting
	[*] Short Assignment Statement
	test := "hello"
	[*] Iota
	use in const to assign number values like 0,1,2,3,4,5.
	[*] Pointers
	[*] Functions
	func test(test string, test1 int) int {
		return 0
	}
	[*] Multiple Returns
	func test(test string, test int) (int, string) {
		return 28, "shazeb"
	}
	[*] Named Returns
	func test(test string, test int) (ret int, i int) { // 89, 90
		ret = 0

		ret = 32
		ret = 89
		i = ret + 1
		return
	}

	[*] For / For range / While / Forever loops
	[*] If/else/
	[*] If with a short statement
	if err := checkStatus(); err != nil {

	}
	[*] Switch
	[*] Fallthrough
	[*] `defer` statement

	[*] Package
	[*] Public / Private indentifiers
	[*] Local / Global/ Private
	//data modelling..
	[*] `type` Keyword [to create custom type.]
	[*] Arrays [ghjjh,hgjhj]
	[*] Array Slices[]
	[*] `append` function (add)
	[-] `make` function
	[ ] Maps
	[*] Structs
	[ ] Methods
	[ ] `new` function
	[ ] Functions as values
	[ ] Function closures

	[ ] Interface
	[ ] Duck Typing
	[ ] Type Assertion
	[ ] Type Switch
	[ ] Stringer

	[*] Goroutines
	[ ] Channels
	[ ] Buffered Channels
	[ ] Range and Close
	[*] Select
	[*] Default Select
	[ ] Syncing
*/