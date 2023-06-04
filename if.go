package main

import "fmt"

func main() {
	name := "Den Ridwan"

	if name == "Den" {
		fmt.Println("Hello, Den")
	} else if name == "Ridwan" {
		fmt.Println("Hello, Ridwan")
	} else {
		fmt.Println("Hello, Boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
