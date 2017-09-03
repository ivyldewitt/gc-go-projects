// Assign a func to a variable, then call that func

package main

import (
	"fmt"
)

func main() {

	darkwood := func(door int) {
		fmt.Printf("The door in Darkwood is number %v.\n", door)
	}

	darkwood(21)

	fmt.Printf("The type of darkwood is %T:", darkwood)

}
