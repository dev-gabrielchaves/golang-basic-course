package main

import "fmt"

func justSomeRandomFunction(x int) (int, error) {
	someMath := x * x

	return someMath, nil
}

func main() {
	num, err := justSomeRandomFunction(10)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("No error occurred")
	}

	fmt.Printf("%v\n", num)
}
