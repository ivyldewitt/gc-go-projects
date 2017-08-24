# Functions

***********************************************************************************

Functions or methods are reusable blocks of code that we have to perform certain specific actions.

### Functions Overview

***********************************************************************************

Functions are about being modular - breaking up code into smaller chunks. Back in the day people use spaghetti code - just implementing up code line-by-line. Eventually we moved to structural or OOP where code was broken up into modules.

In Go, one way to break up code is functions - another is packages.

All functions start with func which is a keyword. Then it has a receiver which is attached to a type. Then an identifier, then params in the identifier, what is being returned, and the code.

```go
package main

import (
    "fmt"
)

func main() {
    foo()
    bar("Cato")
    b := woo("Beyonce")
    fmt.Println(b)
    x, y := movin("Ororo", "Munroe")
    fmt.Println(x)
    fmt.Println(y)
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
    fmt.Println("Electric Lady by Janelle Monae")
}

func bar(s string) {
    fmt.Println(s, "'s favorite song is Jason's Song by Ariana Grande");
}

//The scope of this s is only to this function.
func woo(s string) {
    return fmt.Sprint("Lemonade by ", s)
}
//The above code illustrates how we do a return.

func movin(fn, ln string) (string, bool) {
    a := fmt.Sprint(fn, ln, ` says Hiya!`)
    b := false
    return a, b
}

```

When we define a function we use parameters, when we pass in values those are called arguments.

Everything in Go is pass by value - what you see is what you get.

### Variadic Parameters

***********************************************************************************

We can create a function that takes in an unlimited number of arguments using hte variadic parameter. We use the lexical opearator (...T) in order to pass in vlues of a specific type.

```go
package main

import (
    "fmt"
)

func main() {
    vary(2, 43, 6, 8, 54, 26)
}

//func receiver identifier(params) return code

func vary(x ...int) int {
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    sum := 0

    for i, v := range x {
        sum += v
        fmt.Println("Index: ", i, "Value:", v)
    }

    fmt.Println("Total:", sum)
    return sum
}

//Type is a slice of int
```

[Playground Link](https://play.golang.org/p/IL35g3666s)

We call the ```...``` lexical elements. It's a lexical element operator.

### Unfurling a Slice

***********************************************************************************

When you have a slice of options of a specific type, we can place in the individual values of the slice using the ```x...``` operator.

```go

package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 43, 6, 8, 54, 26}
	//Program recognizes that we want an unlimited number of ints, and we unfurl these values into another slice.
	x := vary(xi...)
	fmt.Println(x)
	
	

}

//func receiver identifier(params) return code

func vary(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	sum := 0
	
	for i, v := range x {
		sum += v
		fmt.Println("Index: ", i, "Value:", v)
	}
	
	fmt.Println("Total:", sum)
	return sum
}

//Type is a slice of int

```

[Unfurling a Slice](https://play.golang.org/p/QuvTFBw-yc)

Note that varadic means 0 or more number of arguments. We can legtimately pass in 0 and that's also fine.

[Todd's Link Go Playground](https://play.golang.org/p/Qc5sq_AK_T)

>"If f is variadic with a final parameter p of type ...T, then within f the type of p is equivalent to type []T. If f is invoked with no actual arguments for p, the value passed to p is nil. Otherwise, the value passed is a new slice of type []T with a new underlying array whose successive elements are the actual arguments, which all must be assignable to T. The length and capacity of the slice is therefore the number of arguments bound to p and may differ for each call site."

Not only is nothing there, but the underyling array did not get created.

[Language Spec on Variadic Params](https://golang.org/ref/spec#Passing_arguments_to_..._parameters)
