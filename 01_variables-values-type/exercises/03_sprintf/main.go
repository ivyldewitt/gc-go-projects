package main

import (
	"fmt"
)

var w int
var b string
var h bool

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", w, b, h)
	fmt.Println(s)
}
