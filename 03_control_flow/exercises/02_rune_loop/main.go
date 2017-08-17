package main

import (
	"fmt"
)

func main() {
	for q := 65; q < 91; q++ {
		fmt.Printf("%v\n", q)
		for i := 0; i < 3; i++ {
			fmt.Printf("\t%#U\n", q)
		}
	}
}
