package main

import (
	"fmt"
)

func main() {
	y := (3 >= 3)
	z := (6 <= 15)
	w := (7 != 7)
	m := (4 == 23)
	n := (2 > 89)
	o := (10 > 1)
	fmt.Println("Y evals to ", y)
	fmt.Println("Z evals to ", z)
	fmt.Println("W evals to ", w)
	fmt.Println("M evals to ", m)
	fmt.Println("N evals to", n)
	fmt.Println("O evals to", o)
}
