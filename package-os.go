package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments:", args[1:])

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Hostname:", hostname)
	}

	gopath := os.Getenv("GOPATH") // get from env system (.zshrc/bashrc/etc in mac)
	fmt.Println("Environment:", gopath)
}