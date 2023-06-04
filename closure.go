package main

import "fmt"

func main() {
	name := "Den"
	counter := 0
	increment := func() {
		name = "Ridwan"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
