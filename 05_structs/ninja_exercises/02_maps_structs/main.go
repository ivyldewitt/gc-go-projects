package main

import (
	"fmt"
)

type person struct {
	firstName        string
	lastName         string
	favoriteIceCream []string
}

func main() {
	storm := person{
		firstName:        "Ororo",
		lastName:         "Munroe",
		favoriteIceCream: []string{"Rocky Road", "Neopolitan", "French Vanilla"},
	}

	wolverine := person{
		firstName:        "James",
		lastName:         "Logan",
		favoriteIceCream: []string{"Chocolate", "Mint Chocolate Chip", "Coffee"},
	}

	// I didn't know you could pass a struct type into this!!!
	xmen := map[string]person{
		storm.lastName:     storm,
		wolverine.lastName: wolverine,
	}

	for _, v := range xmen {
		fmt.Println("First Name:", v.firstName)
		fmt.Println("Last Name:", v.lastName)
		for i, v := range v.favoriteIceCream {
			fmt.Println(i, ":", v)
		}
		fmt.Println("------------------------")
	}
}
