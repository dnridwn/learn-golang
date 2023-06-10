package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Nano Second:", now.Nanosecond())

	date := time.Date(2002, 8, 8, 0, 0, 0, 0, time.UTC)
	fmt.Println(date)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2023-10-01")
	fmt.Println(parse)
}
