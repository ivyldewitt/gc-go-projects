package main

import (
	"fmt"
	"strconv"
)

const (
	year1 = iota + 2017
	year2
	year3
	year4
)

func main() {
	fmt.Println(`The years listed are ` + strconv.Itoa(year1) + `, ` + strconv.Itoa(year2) + `, ` + strconv.Itoa(year3) + ` and ` + strconv.Itoa(year4) + `.`)
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
