package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "Den Ridwan Saputra"
	customer.Address = "Indonesia"
	customer.Age = 21

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	customer2 := Customer{
		Name:    "Customer 1",
		Address: "Indonesia",
		Age:     10,
	}
	fmt.Println(customer2)

	customer3 := Customer{"Customer 2", "Indonesia", 11}
	fmt.Println(customer3)
}
