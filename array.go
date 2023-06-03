package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Den"
	names[1] = "Ridwan"
	names[2] = "Saputra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))
}
