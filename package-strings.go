package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "  Test  "
	fmt.Println(strings.Trim(str1, " "))

	str2 := "SHOULD LOWER"
	fmt.Println(strings.ToLower(str2))

	str3 := "should upper"
	fmt.Println(strings.ToUpper(str3))

	str4 := "1,2,3,4,5"
	fmt.Println(strings.Split(str4, ","))

	str5 := "random string findme random string"
	fmt.Println(strings.Contains(str5, "findme"))

	str6 := "t1m1ka"
	fmt.Println(strings.ReplaceAll(str6, "1", "i"))
}
