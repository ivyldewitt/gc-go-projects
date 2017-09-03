// Build and use an anonymous func
package main

import (
	"fmt"
)

func main() {
	func() {
		for i := 0; i <= 5; i++ {
			fmt.Println(i)
		}
		fmt.Println("I'm an anon function - look at me go!")
	}()
}
