package main

import (
	"fmt"
)

func main() {
	// Creates a channel of type int that has the name dataChan
	// You must think that channel are like portals that allow you to send and receive data
	// By default channels do not have a buffer
	// This means that when you send data to a channel, the sender will be blocked until the receiver receives the data
	// And when you receive data from a channel, the receiver will be blocked until the sender sends the data
	// This is called synchronous communication
	// You can also create channels with a buffer, which allows you to send data to the channel without blocking the sender
	// The syntax for creating a buffered channel is make(chan int, 10), where 10 is the buffer size
	// This means that the channel can hold up to 10 elements before blocking the sender
	// You can also create a channel with a buffer of size 0, which is the same as creating a channel without a buffer
	dataChan := make(chan int)

	go func() {
		dataChan <- 1
	}()

	n := <-dataChan

	fmt.Printf("Got %d from channel!\n", n)

	anotherDataChan := make(chan string, 3)

	anotherDataChan <- "A"
	anotherDataChan <- "B"
	anotherDataChan <- "C"

	// Closing a channel does not remove existing data
	// It just prevents new data from being sent to the channel
	// If not closed, the channel will cause a deadlock in the range loop
	// This is because the range loop will keep waiting for new data to be sent to the channel
	close(anotherDataChan)

	// You can still receive remaining values after the channel is closed
	for n := range anotherDataChan {
		fmt.Printf("Got %s from channel!\n", n)
	}
}
