package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}

	mm := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	v := [][]string{jb, mm}

	fmt.Println("The multi-dimensional array values are:", v)

	for i, xs := range v {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
