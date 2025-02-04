package main

import (
	"fmt"
)

// You can think that GO routines are like lightweight threads
// They are managed by the Go runtime and not by the OS
// They are multiplexed into a small number of OS threads
// This means that if one goroutine blocks, such as waiting for I/O, other goroutines continue to run
// This is in contrast to OS threads, where if a thread blocks, the OS needs to create another thread to continue running other tasks.
// This makes goroutines very cheap to create and destroy
// You can have thousands of goroutines in a single application without any performance issues

// In this script that are two goroutines
// The first goroutine is the main goroutine that is created when the program starts
func main() {
	// Here we create a channel that will be used to communicate between the two goroutines
	dataChan := make(chan int)

	// The second goroutine is created by the go keyword and is an anonymous function that is executed concurrently with the main goroutine
	go func() {
		for i := 0; i < 10; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()

	// The main goroutine reads from the channel and prints the data
	for data := range dataChan {
		fmt.Println(data)
	}
}
