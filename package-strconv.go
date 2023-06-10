package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(boolean)
	}

	integer, err := strconv.ParseInt("1234232", 10, 32)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(integer)
	}

	float, err := strconv.ParseFloat("2121.23", 32)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(float)
	}

	strBool := strconv.FormatBool(true)
	fmt.Println(strBool)
}
