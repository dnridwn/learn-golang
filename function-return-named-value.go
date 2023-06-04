package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Den"
	middleName = "Ridwan"
	lastName = "Saputra"
	return
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
