package main

import "fmt"

func getFullName() (string, string, string) {
	return "Den", "Ridwan", "Saputra"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
