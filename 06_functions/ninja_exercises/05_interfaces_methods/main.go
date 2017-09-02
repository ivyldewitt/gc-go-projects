// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
	height float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.height
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		length: 4,
		height: 8,
	}

	ci := circle{
		radius: 7,
	}

	info(sq)
	info(ci)
}
