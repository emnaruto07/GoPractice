package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	con, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	wg := new(sync.WaitGroup)
	for i := 0; i < con; i++ {
		go worker(wg, i)
		go master(wg, i)
		wg.Add(2) // counter = 10

	}

	wg.Wait() // counter = 0, release
	// all workers are done

	fmt.Println("done")
}

func worker(wg *sync.WaitGroup, n int) {
	defer wg.Done() // counter--
	fmt.Printf("worker %d \n", n)
}

func master(wg *sync.WaitGroup, n int) {
	defer wg.Done()
	fmt.Printf("master %d \n", n)
}

func blyat() {
	// var mu sync.Mutex
	// j := 0

	// for i := 1; i <= 1000; i++ {
	// 	go func() {
	// 		mu.Lock()
	// 		j++
	// 		mu.Unlock()
	// 	}()
	// }
	// time.Sleep(time.Second * 3)
	// fmt.Printf("Value %d \n", j)
}
