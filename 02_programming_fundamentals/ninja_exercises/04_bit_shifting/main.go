package main

import (
	"fmt"
)

func main() {
	q := 88
	fmt.Printf("%d\t%b\t\t%#x\n", q, q, q)
	fmt.Println("---- Bit Shifting ----")
	r := q << 1 //Note - this is a left shift
	fmt.Printf("%d\t%b\t%#x\t", r, r, r)
}
