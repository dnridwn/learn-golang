package main

import "fmt"

func main() {
	name := "Den"

	switch name {
	case "Den":
		fmt.Println("Hello, Den")
	case "Ridwan":
		fmt.Println("Hello, Ridwan")
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Nama terlalu panjang")
	case length2 > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
