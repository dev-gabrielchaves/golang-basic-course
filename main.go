package main

import "fmt"

func main() {
	// Using formatted string
	var age uint8 = 25
	var name = "Gabriel"
	fmt.Printf("My name is %v and I am %v years old\n", name, age)

	// Saving formatted string
	var str = fmt.Sprintf("My name it is not %v and I am not %v years old\n", name, age)

	// Printing it
	fmt.Print(str)
}
