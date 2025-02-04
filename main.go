package main

import (
	"fmt"
	"sync"
)

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello from reader %d!", id)
}

func main() {
	// Introducing go routines

	// WaitGroup is used to wait for the program to finish executing of go routines
	var wg sync.WaitGroup

	// Add 1 to the WaitGroup
	wg.Add(1)

	// As you can see, the go keyword is used to create a new go routine
	// The go routine is a lightweight thread managed by the Go runtime
	go func() {
		// Defer the Done method of the WaitGroup
		defer wg.Done()
		fmt.Println("Hello from go routine!")
	}()

	// Doing the same thing again
	wg.Add(1)

	go func() {
		// Defer the Done method of the WaitGroup
		defer wg.Done()
		fmt.Println("Hello from go routine!")
	}()

	// Call the reader function in a go routine
	wg.Add(2)
	go reader(1, &wg)
	go reader(2, &wg)

	// As you will notice, the output won't be in order!

	// Wait for the go routines to finish
	wg.Wait()
}
