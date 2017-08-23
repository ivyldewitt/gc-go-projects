package main

import (
	"fmt"
)

func main() {
	x := struct {
		books        map[int]string
		favTacoTypes []string
		aRune        rune
	}{
		books: map[int]string{
			1: "Go Programming Language",
			2: "Invent Your Own Computer Games with Python",
			3: "Computer Science Demystified",
		},
		favTacoTypes: []string{"Beef Tacos", "Shrimp Tacos", "Chicken Tacos", "Duck Tacos"},
		aRune:        'x',
	}

	//Printing the values

	for i, v := range x.books {
		fmt.Println("Postion:", i, "Title:", v)
	}

	for _, v := range x.favTacoTypes {
		fmt.Println(v)
	}

	fmt.Println(x.aRune)
}
