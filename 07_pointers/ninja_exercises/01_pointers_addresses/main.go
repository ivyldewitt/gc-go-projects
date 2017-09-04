// Create a value and assign it to a variable.
// Print the address of that value.

package main

import (
	"fmt"
)

func main() {
	bloodbourne := "The Old Hunters"

	fmt.Println(bloodbourne)
	fmt.Println(&bloodbourne)
}
