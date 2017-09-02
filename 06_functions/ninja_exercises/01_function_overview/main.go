// Create a function with the identifier foo that returns an int
// Create a function with the identifier bar that returns an in and string
// Call the functions
// Return the results

package main

import "fmt"

func main() {
	fmt.Println(bar())
	fmt.Println(foo())
}

func foo() int {
	orwell := 1984
	return orwell
}

func bar() (int, string) {
	game := "Darkwood"
	door := 21
	return door, game
}
