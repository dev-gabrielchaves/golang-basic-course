package main

import "fmt"

func main() {
	// Arrays are statically typed, what I mean is that, once set you can not change it's length

	// First way to declare an array...
	var names [3]string = [3]string{"Gabriel", "João", "Thor"}

	// Printing the array
	fmt.Println(names)

	// Second and shorter way to declare an array...
	var numbers = [3]int{1, 2, 3}
	fmt.Println(numbers)

	// Third and even shorter...
	ages := [5]int{15, 18, 25, 31, 65}
	fmt.Println(ages)

	// That's how you change items from  it...
	ages[2] = 1000
	fmt.Println(ages)

	// Slices (Uses arrays under the hood) they don't need to have a size specified

	// Example...
	var sliceExample = []int{1, 2, 3}
	fmt.Println(sliceExample)

	// Append an item an assigning it to the older slice
	sliceExample = append(sliceExample, 6969)
	fmt.Println(sliceExample, len(sliceExample))

	// Ranges...
	range1 := names[1:2]
	range2 := names[:2]
	range3 := names[1:]

	fmt.Println(range1, range2, range3)
}
