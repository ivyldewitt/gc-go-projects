package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

	y := x[0:5]

	fmt.Println(y)

	z := x[5:10]
	fmt.Println(z)
	a := x[2:7]
	fmt.Println(a)
	b := x[1:6]
	fmt.Println(b)
}
