// create a person struct
// create a func called “changeMe” with a *person as a parameter
// change the value stored at the *person address
// important
// to dereference a struct, use (*value).field
// p1.first
// (*p1).first
// “As an exception, if the type of x is a named pointer type and (*x).f
// is a valid selector expression denoting a field (but not a method), x.f
// is shorthand for (*x).f.”

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "Jean",
		last:  "Grey",
		age:   29,
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	(*p).first = "Ororo" //Can also use p.first
	(*p).last = "Munroe"
	(*p).age = 27
}
