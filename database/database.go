package database

import "fmt"

var connection string

func init() { // auto kepanggil ketika package di import
	fmt.Println("Init di panggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}