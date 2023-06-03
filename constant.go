package main

import "fmt"

func main() {
	const firstName = "Den Ridwan"
	const lastName = "Saputra"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		country = "Indonesia"
		city    = "Jakarta"
	)

	fmt.Println(country)
	fmt.Println(city)
}
