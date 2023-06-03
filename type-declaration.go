package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var myNoKTP NoKTP = "999999999"
	fmt.Println(myNoKTP)

	var marriedStatus Married = true
	fmt.Println(marriedStatus)
}
