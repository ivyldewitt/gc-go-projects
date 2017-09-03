// Closure is when we have “enclosed” the scope of a variable in some code block.
// For this hands-on exercise, create a func which “encloses” the scope of a variable:

package main

import (
	"fmt"
)

func main() {
	oxygenNotIncluded()

	{
		vy := 24
		fmt.Println(vy)
	}

	// fmt.Println(vy) error
	// fmt.Println(colony) error
}

func oxygenNotIncluded() {
	colony := 3
	for colony <= 12 {
		fmt.Println("Our colony now has:", colony, "dupes!")
		colony++
	}
	fmt.Println("The final total of the colony is:", colony)
}
