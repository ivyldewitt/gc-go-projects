package main

import (
	"fmt"
)

func main() {
	num := 1993
	ivy := 24
	fmt.Printf("%T\n", num)
	fmt.Printf("%d\n", num)
	fmt.Printf("%b\n", num)
	fmt.Printf("%#X\n", num)
	fmt.Println("-----------")
	fmt.Printf("%T\t%d\t%b\t%#x", ivy, ivy, ivy, ivy)
}
