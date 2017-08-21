package main

import (
	"fmt"
)

func main() {
	m := 34

	if m > 50 {
		fmt.Printf("The value of m: %v, is not equal to 34.", m)
	} else if m == 34 {
		fmt.Printf("The value of m: %v is equal to 34.", m)
	} else {
		fmt.Printf("The value of m: %v is not less than 34", m)
	}
}
