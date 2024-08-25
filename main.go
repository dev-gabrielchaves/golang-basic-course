package main

import "fmt"

func main() {

	whatever := []string{"Jump me", "Jump me", "Hey!", "Jump me", "Hello beautiful world!", "Break here", "I should not come up!"}

	for i := range whatever {
		if whatever[i] == "Jump me" {
			continue
		} else if whatever[i] == "Break here" {
			fmt.Println()
			break
		} else {
			fmt.Print(whatever[i] + " ")
		}
	}
}
