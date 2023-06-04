package main

import "fmt"

type Customer2 struct {
	Name, Address string
	Age           int
}

func (customer Customer2) sayHi() {
	fmt.Println("Hi, my name is", customer.Name)
}

func main() {
	var customer Customer2
	customer.Name = "Den Ridwan Saputra"
	customer.Address = "Indonesia"
	customer.Age = 21

	customer.sayHi()
}
