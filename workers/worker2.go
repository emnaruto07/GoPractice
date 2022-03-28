package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	go worker(c, 1)
	go worker(c, 2)
	master(c)
	time.Sleep(time.Second * 10)
}

//go worker c <-chan int(receving)

func worker(c <-chan int, n int) {
	for val := range c {
		fmt.Printf("worker %d got %d\n", n, val)
		time.Sleep(time.Second * 2)
	}

	fmt.Println("worker done")
}

//go master c  (sending)

func master(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	fmt.Println("master done")
}

wg + chan
interface
