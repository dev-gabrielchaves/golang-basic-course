package main

import "fmt"

func greeting(name string) {
	fmt.Printf("Hello %v!\n", name)
}

func sayBye(name string) {
	fmt.Printf("Bye %v!\n", name)
}

func loopThrough(names []string, function func(string)) {
	for i := range names {
		function(names[i])
	}
}

func main() {
	names := []string{"John", "Gabriel", "Robert", "Carla"}

	loopThrough(names, greeting)

	fmt.Println()

	loopThrough(names, sayBye)
}
