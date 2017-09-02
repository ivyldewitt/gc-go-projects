// create a func with the identifier foo that
// takes in a variadic parameter of type int
// pass in a value of type []int into your func (unfurl the []int)
// returns the sum of all values of type int passed in
// create a func with the identifier bar that
// takes in a parameter of type []int
// returns the sum of all values of type int passed in

package main

import (
	"fmt"
)

func main() {
	fi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(foo(fi...))
	fmt.Println(bar(bi))
}

func foo(iv ...int) int {
	sum := 0
	for _, v := range iv {
		//fmt.Println(v)
		sum += v
	}
	return sum
}

func bar(ii []int) int {
	sum := 0
	for _, v := range ii {
		//fmt.Println(v)
		sum += v
	}
	return sum
}
