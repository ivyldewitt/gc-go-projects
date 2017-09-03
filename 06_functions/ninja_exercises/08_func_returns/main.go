// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

package main

import (
	"fmt"
)

func main() {
	b := bloodbourne()

	fmt.Printf("%T\n", b)
	fmt.Println(b())
}

func bloodbourne() func() string {
	return func() string {
		return "The entrance to the Hunter's Nightmare will be found at Oeden Chapel."
	}
}
