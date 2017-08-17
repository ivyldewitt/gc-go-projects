package main

import (
	"fmt"
)

func main() {
	fmt.Println("Numbers 10-100 divided by 4.")
	for q := 10; q < 100; q++ {
		fmt.Printf("The remainder of %v divided by 4 is: %v\n", q, q%4)
	}
}
