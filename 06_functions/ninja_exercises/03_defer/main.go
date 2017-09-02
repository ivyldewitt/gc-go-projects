// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

package main

import (
	"fmt"
)

func main() {
	fi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	defer foo(fi...)
	defer woo()
	bar(bi)
}

func woo() string {
	s := "Nah nah nah"
	fmt.Println(s)
	return s
}

func foo(iv ...int) int {
	sum := 0
	for _, v := range iv {
		//fmt.Println(v)
		sum += v
	}
	fmt.Println("The sum of foo is:", sum)
	return sum
}

func bar(ii []int) int {
	sum := 0
	for _, v := range ii {
		//fmt.Println(v)
		sum += v
	}
	fmt.Println("The sum of bar is:", sum)
	return sum
}

//Interesting note - defer appears to go from top to bottom.

// Output
// The sum of bar is: 55
// Nah nah nah
// The sum of foo is: 55
