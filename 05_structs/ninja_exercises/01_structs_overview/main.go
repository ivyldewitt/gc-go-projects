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

	fmt.Println("Storm:", storm.firstName, storm.lastName)

	for _, v := range storm.favoriteIceCream {
		fmt.Println(v)
	}

	fmt.Println("Wolverine: ", wolverine.firstName, wolverine.lastName)

	for _, v := range wolverine.favoriteIceCream {
		fmt.Println(v)
	}
}
