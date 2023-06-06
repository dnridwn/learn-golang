package main

import (
	"fmt"
	"flag"
)

func main() {
	host := flag.String("host", "localhost", "Put your host")
	username := flag.String("username", "root", "Put your username")
	password := flag.String("password", "password", "Put your password")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
}