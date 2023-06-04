package main

import "fmt"

func endApp2() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
}

func runApp2(error bool) {
	defer endApp2()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp2(true)
	fmt.Println("Continue")
}
