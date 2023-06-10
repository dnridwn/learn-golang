package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("d([a-z])n")
	fmt.Println(regex.MatchString("den"))
	fmt.Println(regex.MatchString("dan"))
	fmt.Println(regex.MatchString("don"))
	fmt.Println(regex.MatchString("dln"))

	fmt.Println(regex.FindAllString("den dan don dln din dun dhn kln uih gdsd", -1))
}
