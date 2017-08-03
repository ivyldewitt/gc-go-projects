package main

import (
	"fmt"
)

type rihanna int

var rr rihanna

func main() {
	fmt.Println(rr)
	fmt.Printf("%T\n", rr)
	rr = 42
	fmt.Printf("%v", rr)
}
