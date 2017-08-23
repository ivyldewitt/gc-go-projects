package main

import (
	"fmt"
)

type vehicle struct {
	numOfDoors int
	color      string
}

type truck struct {
	vehicle
	fourWheels bool
}

type jaguar struct {
	vehicle
	luxury bool
}

func main() {

	j := jaguar{
		vehicle: vehicle{
			numOfDoors: 4,
			color:      "Silver",
		},
		luxury: true,
	}

	t := truck{
		vehicle: vehicle{
			numOfDoors: 4,
			color:      "Magenta",
		},
		fourWheels: true,
	}

	fmt.Println(t.numOfDoors, t.color, t.fourWheels)
	fmt.Println(j.numOfDoors, j.color, j.luxury)
}

// Create a new type: vehicle.
// The underlying type is a struct.
// The fields:
// doors
// color
// Create two new types: truck & sedan.
// The underlying type of each of these new types is a struct.
// Embed the “vehicle” type in both truck & sedan.
// Give truck the field “fourWheel” which will be set to bool.
// Give sedan the field “luxury” which will be set to bool. solution
// Using the vehicle, truck, and sedan structs:
// using a composite literal, create a value of type truck and assign values to the fields;
// using a composite literal, create a value of type sedan and assign values to the fields.
// Print out each of these values.
// Print out a single field from each of these values.
