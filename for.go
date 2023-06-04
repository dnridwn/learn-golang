package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke ", counter2)
		counter++
	}

	names := []string{
		"Den",
		"Ridwan",
		"Saputra",
	}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	mapNames := map[string]string{
		"satu": "Den",
		"dua":  "Ridwan",
		"tiga": "Saputra",
	}
	for index, mapNames := range mapNames {
		fmt.Println(index, mapNames)
	}
}
