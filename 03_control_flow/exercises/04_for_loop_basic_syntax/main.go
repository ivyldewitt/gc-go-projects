package main

import (
	"fmt"
)

func main() {
	fmt.Println("How long ago was I born?")
	year := 1993
	for {
		year++
		if year <= 2017 {
			fmt.Println(year)
		}
		if year > 2017 {
			break
		}
		//Less efficient solution but works.
	}
	fmt.Println("I was born in 1993 and I'm 24 years old!")
}
