// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument

package main

import (
	"fmt"
)

func main() {

	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	xd := double(xs...)
	fmt.Printf("%T\n", xd)
	fmt.Println(xd)

	xc := summation(double, xs...)
	fmt.Printf("%T\n", xc)
	fmt.Println(xc)
}

func double(xv ...int) []int {
	var x []int
	for _, v := range xv {
		v *= 2
		x = append(x, v)
	}
	return x

}

func summation(fc func(xv ...int) []int, xq ...int) int {
	total := 0
	for _, v := range xq {
		total += v
	}

	return total
}

// https://play.golang.org/p/WeDhi1HXkD
