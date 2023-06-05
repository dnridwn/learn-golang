package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	var contohError error = errors.New("Ups")
	fmt.Println(contohError)

	result, err := Pembagian(12, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result", result)
	}
}
