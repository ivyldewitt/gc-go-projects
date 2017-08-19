package main

import (
	"fmt"
)

func main() {
	favSport := "Basketball"

	switch favSport {
	case "Soccer":
		fmt.Println("Soccer or Futbol?")
	case "Basketball":
		fmt.Println("Dunkin'")
	case "Baseball":
		fmt.Println("Eh, no good ideas for this one.")
	}
}
