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

### Defer

***********************************************************************************

We're starting to de-couple our code, aka break it up into modular chunks that we can use whenever we need to. It makes our code easier to read and use.

The ```defer``` keyword will defer the execution of a function until wherever its being called comes to an end.

>"A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking."

[Golang Lang Spec on Defer](https://golang.org/ref/spec#Defer_statements)

```go

package main

import (
    "fmt"
)


func main() {
    defer elixir()
    golang()
}

func golang() {
    fmt.Println("Golang Rocks!")
}

func elixir() {
    fmt.Println("Haven't learned Elixir yet but I'm excited to!");
}
```

Wherever the file exits, defer is always going to run whenever the containing function exits.

### Methods

***********************************************************************************

A method is nothing more than a function attached to a TYPE. When you create a function of a type it is a method of that type.

>"A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type."

```plaintext

MethodDecl = "func" Receiver MethodName ( Function | Signature ) .
Receiver   = Parameters .

```

>"The receiver is specified via an extra parameter section preceding the method name. That parameter section must declare a single non-variadic parameter, the receiver. Its type must be of the form T or *T (possibly using parentheses) where T is a type name. The type denoted by T is called the receiver base type; it must not be a pointer or interface type and it must be defined in the same package as the method. The method is said to be bound to the base type and the method name is visible only within selectors for type T or *T."

[Golang Lang Spec - Method Declariations](https://golang.org/ref/spec#Method_declarations)

When you have a receiver you attach a function to that type. So any value of that type has access to that method.

```go

package main

import (
    "fmt"
)

type dog struct {
    name     string
    breed    string
    fourLegs bool
}

type cartoonDog struct {
    dog
    catchphrase string
}

func (c cartoonDog) bark() {
    fmt.Println("Ruh-Ruh! We're in trouble!")
}

//func (r receiver) identifier(params) return(s) {}

func main() {
    d1 := dog{
        name:     "Toby",
        breed:    "Pomeranian-Sheltie Mix",
        fourLegs: true,
    }

    d2 := dog{
        name:     "Macho",
        breed:    "Bishon mix?",
        fourLegs: true,
    }

    scooby := cartoonDog{
        dog: dog{
            name:     "Scooby Dog",
            breed:    "Mutt",
            fourLegs: true,
        },
        catchphrase: "Ruh-roh!",
    }

    fmt.Println(d1)
    fmt.Println(d2)
    fmt.Println(scooby)
    scooby.bark()
}

```

[Go Playground](https://play.golang.org/p/WXKF8rOUx8)

### Interfaces & Polymorphism

***********************************************************************************

Moving into more intense material. Settle back & enjoy the ride.

Learning interfaces and polymorphism. Interfaces allow you to define behavior and implement polymorphism.

The structure of an interface (and other variables) is: ```keyword identifer type```.

Any type that has the method speak() is also of type human. A value can be associated with more than one type. Example, the value is of type cartoonDog but also of type animal.

I can also attach the method speak to cartoonDog and dog. Adding those means that they are both also of type animal.

Additionally we can create a new function of animalSpeak that takes in animal, and pass in our variables for cartoonDog and Dog.

This is **polymorphism** - this function can take in many different types . Interfaces allow a value to be of more than one type. And value of Dog or cartoonDog are also of type animal.

Looking at documentation. For example, if we go to package io - the type Writer is an interface.

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

[type Writer](https://golang.org/pkg/io/#Writer)

For example, in net/http, many functions take in a writer - so interfaces allow us to share behavior across various examples:

```go
func (r *Request) Write(w io.Writer) error
```

[func Request Ex](https://golang.org/pkg/net/http/#Request.Write)

#### What does an interface say?

If you have these methods, then you're my type.

Review on conversion:

```go
    // conversion
    var x hotdog = 42
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    var y int
    y = int(x)
    fmt.Println(y)
    fmt.Printf("%T\n", y)
```

We can also implement a special switch statement to switch on type person via ```switch x.(type)```. This allows us to do type assertion. If we want to get back to the underlying type under that the interface was part of, asserting is much strong and allows us to confirm that underlying type (in this example, it's a struct).

```go
package main

import (
	"fmt"
)

type animal interface {
	speak()
}

type dog struct {
	name     string
	breed    string
	fourLegs bool
}

type cartoonDog struct {
	dog
	catchphrase string
}

func (c cartoonDog) speak() {
	fmt.Println("Ruh-Roh! We're in trouble!")
}

func (d dog) speak() {

	fmt.Println(d.name, "says woof!")
}

func animalSpeak(a animal) {
	switch a.(type) {
	case dog:
		fmt.Println("I am on top of the animal kingdom!", a.(dog).name)
	case cartoonDog:
		fmt.Println("I am on top of the animal kingdom!", a.(cartoonDog).name)
	}
}

//func (r receiver) identifier(params) return(s) {}

func main() {
	d1 := dog{
		name:     "Toby",
		breed:    "Pomeranian-Sheltie Mix",
		fourLegs: true,
	}

	d2 := dog{
		name:     "Macho",
		breed:    "Bishon mix?",
		fourLegs: true,
	}

	scooby := cartoonDog{
		dog: dog{
			name:     "Scooby Dog",
			breed:    "Mutt",
			fourLegs: true,
		},
		catchphrase: "Ruh-roh!",
	}

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(scooby)
	scooby.speak()
	d2.speak()
	d1.speak()

	animalSpeak(scooby)
	animalSpeak(d1)
}
```

[Go Playground](https://play.golang.org/p/eqXaa1R194)

#### Review

Remember what interfaces do: they allow us to interface in a different way with types. A value can be of more than one type. An interface says, if you have this method speak, you're also of type human.

If we ask for a value of type animal, we can pass in dog and cartoonDog. If we have an empty interface, then every other type has at least **no** methods. Then all of those types belong to that empty interface.

Example: Fmt Package.

All these functions take in an empty interface, which means all values can be passed into it. For example, even type int also is type ...interface{}.

```go

func Print(a ...interface{}) (n int, err error)

func Println(a ...interface{}) (n int, err error)

func Sprintf(format string, a ...interface{}) string
```

*Takeaway: Every type has no methods.*

[Fmt Package](https://golang.org/pkg/fmt/#Print)

This provides a lot of flexibility and power in creating programing.