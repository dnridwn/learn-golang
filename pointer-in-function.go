package main

import "fmt"

type Address3 struct {
	city     string
	province string
	country  string
}

func ChangeAddressToIndonesia1(address *Address3) {
	address.country = "Indonesia 1"
}

func main() {
	address := Address3{
		city:     "Jakarta",
		province: "DKI Jakarta",
		country:  "Indonesia",
	}

	ChangeAddressToIndonesia1(&address)
	fmt.Println(address)
}
