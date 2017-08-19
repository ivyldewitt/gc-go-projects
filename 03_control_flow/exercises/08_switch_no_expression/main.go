package main

import (
	"fmt"
)

func main() {
	switch {
	case (3 != 3):
		fmt.Println("Blue")
	case (4 == 4):
		fmt.Println("Green")
	case (5 > 7):
		fmt.Println("Blue")
	case (4 < 2):
		fmt.Println("Purple")
	}
}
