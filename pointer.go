package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	// no pointer
	address1 := Address{"Sumedang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1) // address 1 still Sumedang
	fmt.Println(address2)

	// with pointer
	address3 := Address{"Sumedang", "Jawa Barat", "Indonesia"}
	address4 := &address3 // => var address4 *Address = &address3

	address4.City = "Bandung"

	fmt.Println(address3) // address 3 changed to "Bandung"
	fmt.Println(address4)

	address5 := &address3
	address5 = &Address{"Sumedang 123", "Jawa Barat", "Indonesia"}

	fmt.Println(address3) // address 3 will not change
	fmt.Println(address5)

	address6 := &address3
	*address6 = Address{"Sumedang 123", "Jawa Barat", "Indonesia"} // force all reference varibles to change also

	fmt.Println(address3) // address 3 will change
	fmt.Println(address6)

	// another way to create pointer
	address7 := new(Address)
	fmt.Println(address7)
}
