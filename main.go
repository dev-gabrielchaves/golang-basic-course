package main

import "fmt"

func main() {
	// Simple for loop, looks more like a while loop to me but yep...
	x := 0
	for x < 3 {
		x++
		fmt.Println("The value of x is: ", x)
	}

	fmt.Println("")

	// This one looks more like a for loop hehe...
	for y := 0; y < 3; y++ {
		fmt.Println("The value of y is: ", y)
	}

	fmt.Println("")

	// Looping through a slice...
	var names []string = []string{"John", "Gabriel", "Carla", "Robert"}
	for i := 0; i < len(names); i++ {
		fmt.Printf("The %v° name is: %v\n", i+1, names[i])
	}

	//Another way of doing it
	for i := range names {
		fmt.Printf("The %v° name is: %v\n", i+1, names[i])
	}
}
