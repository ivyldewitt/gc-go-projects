package main

import (
	"fmt"
)

func main() {
	fmt.Println("How many years have I been alive?")
	year := 1993
	for year <= 2017 {
		fmt.Println(year)
		year++
	}
	fmt.Println("I'm 24 years old!")
}
