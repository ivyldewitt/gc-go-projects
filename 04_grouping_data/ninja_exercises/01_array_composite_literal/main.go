package main

import (
	"fmt"
)

func main() {
	a := [5]int{67, 23, 43, 54, 64}

	for i, v := range a {
		fmt.Println("Index:", i, "Value:", v)
	}
	fmt.Printf("Type of a: %T", a)
}
