package main

import (
	"fmt"
)

type rihanna int

var rr rihanna

var wfl rihanna

func main() {
	fmt.Println(rr)
	fmt.Printf("%T\n", rr)
	rr = 42
	fmt.Printf("%v\n", rr)
	fmt.Println("------ seperator ------")
	wfl := int(rr)
	fmt.Println(wfl)
	fmt.Printf("%T\n", wfl)
}
