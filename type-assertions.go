package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)

	fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	switch result.(type) {
	case string:
		fmt.Println(result.(string))
	case int:
		fmt.Println(result.(int))
	case bool:
		fmt.Println(result.(bool))
	}
}
