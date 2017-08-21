package main

import (
	"fmt"
)

func main() {
	xmen := map[string][]string{
		"munroe_ororo": []string{"ororo", "munroe", "storm", "flying"},
		"logan_james":  []string{"james", "logan", "wolverine", "eating"},
		"pryde_kitty":  []string{"kitty", "pryde", "shadowcat", "sleeping"},
	}

	xmen["labeau_remny"] = []string{"remny", "labeau", "gambit", "flirting"}

	for i, v := range xmen {
		fmt.Println("Username:", i)
		for i, vv := range v {
			fmt.Println("--- Index:", i, " --- Values:", vv)
		}
	}
}
