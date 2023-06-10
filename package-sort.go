package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (v UserSlice) Len() int {
	return len(v)
}

func (v UserSlice) Less(i int, j int) bool {
	return v[i].Age < v[j].Age
}

func (v UserSlice) Swap(i int, j int) {
	v[i], v[j] = v[j], v[i]
}

func main() {
	users := UserSlice{
		{"Den", 30},
		{"Ridwan", 25},
		{"Saputra", 40},
	}

	sort.Sort(users)

	fmt.Println(users)
}
