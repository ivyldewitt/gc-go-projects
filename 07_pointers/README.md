# Pointers

*************************************************************************************

Go supports pointers, allowing you to pass references to values and records within your program.

>"In computer science, a pointer is a programming language object, whose value refers to (or "points to") another value stored elsewhere in the computer memory using its memory address. A pointer references a location in memory, and obtaining the value stored at that location is known as dereferencing the pointer. As an analogy, a page number in a book's index could be considered a pointer to the corresponding page; dereferencing such a pointer would be done by flipping to the page with the given page number and reading the text found on the indexed page."

[Wikipedia on Pointers](https://en.wikipedia.org/wiki/Pointer_(computer_programming))

## Pointers Overview

*************************************************************************************

A pointer is simply a pointer to a location in memory where a value is stored. Memory in a computer is like a post office which a bunch of P.O. boxes, and we can store a value in each of those P.O. boxes. And every box has an address that we reference to.

```go
package main

import (
    "fmt"
)

func main() {
    a := 42
    fmt.Println(a)
    fmt.Prinltn(&a) // ampersand references the address in memory
    fmt.Printf("%T\n",a)
    fmt.Printf("%T\n", a)

    var b int = &a //This will give an error
    var c *int = &a // This works
}

```

We can use the ```&``` operator to see the memory address. If we reference the type of the address it's a type ```*int```, aka type pointer to an int. This is a completely different type than an int. We can't mix/match these. It's a type **pointer to a location where the int is stored.**

Everything in Go is pass by value - this allows us to share addresses, but we can't assign a pointer to a type to the type.

If we wanted to see the values stored at that address, we can use the ```*``` to dereference the variable. If we see the ```*T``` notation, that's the pointer to a type. If you see it like ```*v``` you are deferencing the variable and retreieving the value ouf of it.

```go
package main

import (
    "fmt"
)

func main() {
    a := 42
    fmt.Println(a)
    fmt.Prinltn(&a) // & references the address in memory
    fmt.Printf("%T\n",a)
    fmt.Printf("%T\n", a)

    d := &a
    fmt.Println(d)
    fmt.Println(*d)
    fmt.Printf("%T\n", *d) // * gives you the value stored at that address
    fmt.Println(*&a) //42

    *d = 24 //Taking the address, then reassign the value
    fmt.Println(a) // Now a is equal to 24
}

```

We can take an address, and once we have an address, we can dereferencfe that adddress and get the value that's stored in there with the asterisk. 

Why is a equal to 43? Because we're taking a value and printing it at that address. A is the value stored at the address, b is a reference to that memory address, so changing b, changes a.

To put it in other terms, these values are pointing to the same post office box.

## When to use Pointers

*************************************************************************************

Pointers are useful when you have a large chunk of data and you don't want to pass around all of the data where that value is stored. Instead of 'carrying' that data with you, you can carry around a pointer to that data instead and reference it whenever necessary.

Another way to use pointers is to change something that's at a certain location. If you need to change a value you can use a pointer and it'll change the value at that location.

One important thing to note is that **everything in Go is pass by value**. It's the only phrase you need to know and use. *What you see is what you go.* Don't worry about the other common phrases.

### Pointers

```go
package main

import (
    "fmt"
)

func main() {
    d := 0
    oni(d)
    fmt.Println(d)
}

func oni(c int) {
    fmt.Println(c)
    c = 6
    fmt.Println(c)
}

// Output 0 6 0
```

We create a function, pass in d. In the function, we assign c to 6, then print it. Then we exit the variable and print D again.

### Pointer to an Int

```go
package main

import (
    "fmt"
)

func main() {
    d := 0
    fmt.Println("D before", &d)
    fmt.Println("D before", d)
    oni(&d)
    fmt.Println("D after", &d)
    fmt.Println("D after", d)
}

func oni(c *int) {
    fmt.Println("C before", c)
    fmt.Println("C before", *c)
    *c = 6
    fmt.Println("C after", *c)
    fmt.Println("C after", c)
}

```

[Go Playground](https://play.golang.com/p/tPyw8KJ9Rb)

The above example is how a pointer allows us to mutate a value (i.e. change the value of an address).

## Method Sets

*************************************************************************************

Method Sets - when we have a type, we can attach a method set. Those are called method sets. Depending on where we use a pointer or a non-pointer type - we have specific rules on how those method sets can work.

>"A type may have a method set associated with it. The method set of an interface type is its interface. The method set of any other type T consists of all methods declared with receiver type T. The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T (that is, it also contains the method set of T). Further rules apply to structs containing embedded fields, as described in the section on struct types. Any other type has an empty method set. In a method set, each method must have a unique non-blank method name. The method set of a type determines the interfaces that the type implements and the methods that can be called using a receiver of that type."

[Golang Spec - Method Sets](https://golang.org/ref/spec#Method_sets)

Method sets determine what methods attach to a TYPE. It is exactly what the name says: What is the set of methods for a given type? That is its method set.

### IMPORTANT: “The method set of a type determines the INTERFACES that the type implements.....”

The above “important” was left out of this video but will be discussed in the “concurrency” section in a video called “method sets revisited”.

#### a NON-POINTER RECEIVER

works with values that are POINTERS or NON-POINTERS.

#### a POINTER RECEIVER

only works with values that are POINTERS.

Receivers     |  Values

---------------------------

(t T)         |  T and *T
(t *T)        |  *T

Non pointer receiver and non-pointer type:

```go
package main

import (
    "fmt"
    "math"
)

type circle struct {
    radius float64
}

type shape interface {
    area() float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func info(s shape) {
    fmt.Println("area", s.area())
}

func main() {
    c := circle{5}
}
```

Pointer receiver and pointer type:

```go
package main

import (
    "fmt"
    "math"
)

type circle struct {
    radius float64
}

type shape interface {
    area() float64
}

func (c *circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func info(s shape) {
    fmt.Println("area", s.area())
}

func main() {
    c := circle{5}
    info(&c)
}
```

Non-Pointer receiver and a Pointer value

```go
package main

import (
    "fmt"
    "math"
)

type circle struct {
    radius float64
}

type shape interface {
    area() float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func info(s shape) {
    fmt.Println("area", s.area())
}

func main() {
    c := circle{5}
    info(&c)
}
```

### Resources

[Go by Ex: Pointers](https://gobyexample.com/pointers)
[A Tour of Go](https://tour.golang.org/moretypes/1)
[GolangBot: Pointers](https://golangbot.com/pointers/)
[Stack Overflow](https://stackoverflow.com/questions/15096329/golang-pointers)
[Going Go Do Not](https://www.goinggo.net/2014/12/using-pointers-in-go.html)
[Golang Spec: Pointeres](https://golang.org/ref/spec#Pointer_types)