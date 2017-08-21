package main

import (
	"fmt"
)

func main() {
	slicing := []int{2, 54, 62, 66, 73, 81, 90, 4, 27, 10}

	for i, v := range slicing {
		fmt.Println("Index:", i, "Value:", v)
	}

	fmt.Printf("Type of slicing: %T", slicing)
}

// Using a COMPOSITE LITERAL:
// create a SLICE of TYPE int
// assign 10 VALUES
// Range over the slice and print the values out.
// Using format printing
// print out the TYPE of the slice
