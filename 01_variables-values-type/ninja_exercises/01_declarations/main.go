package main

import (
	"fmt"
)

func main() {
	x := 72
	y := "Sweet Nothing by Florence Welch"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\t%T\t%T", x, y, z)
}
