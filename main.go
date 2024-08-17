package main

// fmt stands for the Format package. This package allows to format basic strings, values, or anything and print them or...
// collect user input from the console, or write into a file using a writer or even print customized fancy error messages.
import "fmt"

func main() {
	// That's a way to declare a string variable, first you type "var", then the name of the variable and it's type...
	// Every variable should be used in golang, otherwise you will get an error! :(
	var myName string = "Gabriel"

	// It is also possible to concatenate strings! :)
	fmt.Println("My name is: " + (myName))

	// In this way, the variable type will be infered
	var myFriendsName = "João"

	fmt.Println("My friend's name is: " + (myFriendsName))

	// You can also define a variable and use it later on...
	// But the type should always be a string, otherwise it won't work
	var myFathersName string

	fmt.Println("My father's name is: " + (myFathersName))

	// That's a shortcut, but it can only be used inside of a function scope
	anyName := "Yoshi"

	fmt.Println((anyName))

	// int types, they are defined in the same way
	var numOne int = 1
	var numTwo = 2
	numThree := 3

	fmt.Println(numOne, numTwo, numThree)

	// int types can have the number of bits defined, ranges can be sourced...
	var numTyped int8 = 1

	fmt.Println(numTyped)

	// same then int for floats
	var numFLoat float32 = 123.25

	fmt.Println(numFLoat)
}
