# Functions

***********************************************************************************

Functions or methods are reusable blocks of code that we have to perform certain specific actions.

## Functions Overview

***********************************************************************************

Functions are about being modular - breaking up code into smaller chunks. Back in the day people use spaghetti code - just implementing up code line-by-line. Eventually we moved to structural or OOP where code was broken up into modules.

In Go, one way to break up code is functions - another is packages.

All functions start with func which is a keyword. Then it has a receiver which is attached to a type. Then an identifier, then parameters in the identifier, what is being returned, and the code.

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

## Variadic Parameters

***********************************************************************************

We can create a function that takes in an unlimited number of arguments using the variadic parameter. We use the lexical operator (...T) in order to pass in values of a specific type.

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

## Unfurling a Slice

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

Note that variadic means 0 or more number of arguments. We can legitimately pass in 0 and that's also fine.

[Todd's Link Go Playground](https://play.golang.org/p/Qc5sq_AK_T)

>"If f is variadic with a final parameter p of type ...T, then within f the type of p is equivalent to type []T. If f is invoked with no actual arguments for p, the value passed to p is nil. Otherwise, the value passed is a new slice of type []T with a new underlying array whose successive elements are the actual arguments, which all must be assignable to T. The length and capacity of the slice is therefore the number of arguments bound to p and may differ for each call site."

Not only is nothing there, but the underyling array did not get created.

[Language Spec on Variadic Params](https://golang.org/ref/spec#Passing_arguments_to_..._parameters)

## Defer

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

Interesting note - defer appears to go from top to bottom. Reveiw file: /Users/ivydewitt/go/src/github.com/zentechnista/gc-go-projects/06_functions/ninja_exercises/03_defer/main.go for reference - if I create 3 functions, the 'top' defer function is the last one to go.

## Methods

***********************************************************************************

A method is nothing more than a function attached to a TYPE. When you create a function of a type it is a method of that type.

>"A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type."

```plaintext

MethodDecl = "func" Receiver MethodName ( Function | Signature ) .
Receiver   = Parameters .

```

>"The receiver is specified via an extra parameter section preceding the method name. That parameter section must declare a single non-variadic parameter, the receiver. Its type must be of the form T or *T (possibly using parentheses) where T is a type name. The type denoted by T is called the receiver base type; it must not be a pointer or interface type and it must be defined in the same package as the method. The method is said to be bound to the base type and the method name is visible only within selectors for type T or *T."

[Golang Lang Spec - Method Declarations](https://golang.org/ref/spec#Method_declarations)

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

## Interfaces & Polymorphism

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

### What does an interface say?

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

### Review

***********************************************************************************

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

## Anonymous Functions

***********************************************************************************

An anonymous (unnamed) self-executing functions.

```go

package main

import (
    "fmt"
)

func main() {
    dog()

    func() {
        fmt.Println("Anon func ran!")
    }()

    func(x int) {
        fmt.Println("Anon fun with int:" x)
    }(7)
}

func dog(){
    fmt.Println("Non anon func ran!")
}

```

Note that with anon functions we immediately invoke it with the () after the function declaration.

## Function Expressions

***********************************************************************************

Assiging the value of a function to a variable. 

```go

package main

import (
    "fmt"
)

func main() {
   f := func(){
       fmt.Println("My 1st function expression!")
   }

   f()

   ob := func(x int){
       fmt.Println("The book in Observer was: ", 1984, " by George Orwell")
   }

   ob(1984)
}
```

In Go, functions are first-class citizens. This means that a function can be used just like any other variable. Aka a function is a type just like any other types.

## Returning a Function

***********************************************************************************

How we can return a function from a function.

```go

package main

import (
    "fmt"
)

func main() {
    s1 := observer()
    fmt.Println(s1)

    d := darkwood()

    fmt.Printf("%T\n", d)  // func() int
    fmt.Println(d())
}


func observer() string {
    ob := "Observer was an OK game"
    return ob
}


func darkwood() func() int {
    return func() int {
        return 21
    }
}
```

[Playground Example](https://play.golang.org/p/uso7xtlsgz)

```go

package main

import (
    "fmt"
)

func main() {

    fmt.Printf("%T\n", darkwood())  // func() int
    fmt.Println(darkwood()())
}
func darkwood() func() int {
    return func() int {
        return 21
    }
}
```

[Playground Example](https://play.golang.org/p/WaYfdmo1cU)

### Why would we want to do this?

Sometimes you'll return functions, for example we may need to return a function that returns a function and formats it in a specific way.

## Callback

***********************************************************************************

A a callback function is when we pass a function as a argument. A function can be returned to a func, assigned to variables, pass a function into another function as an argument.

This is a type of programming called functional programming. We may use this to do and accomplish specific things.

```go

package main

import (
    "fmt"
)

func main() {
    ii := []int{3, 45, 53, 1, 22, 5, 7, 9}
    q := sum(ii...)
    fmt.Println(q)
}

func sum(xi ...int) int {
    fmt.Printf("%T\n", xi)
    total := 0
    for _, val := range xi {
        total += val
    }
    return total
}

```

Adding an even function.

```go

package main

import (
    "fmt"
)

func main() {
    ii := []int{3, 45, 53, 1, 22, 5, 7, 9}
    q := sum(ii...)
    fmt.Println(q)

    v := even(sum, ii...)
    fmt.Println(v)
}

func sum(xi ...int) int {
    fmt.Printf("%T\n", xi)
    total := 0
    for _, val := range xi {
        total += val
    }
    return total
}
// Callback  func(xi ...int) int

func even(f func(xi ...int) int, iv ...int) int {
    var vy []int
    for _, v := range iv {
        if v % 2 == 0 {
            vy = append(vy, v)
        }
    }
    return f(vy...)
}

```

Odd callback function:

```go
package main

import (
    "fmt"
)

func main() {
    ii := []int{1, 2, 3, 4, 5, 6, 7,8, 9}
    q := sum(ii...)
    fmt.Println(q)

    v := odd(sum, ii...)
    fmt.Println(v)
}

func sum(xi ...int) int {
    fmt.Printf("%T\n", xi)
    total := 0
    for _, val := range xi {
        total += val
    }
    return total
}

// Callback  func(xi ...int) int

func odd(f func(xi ...int) int, iv ...int) int {
    var vy []int
    for _, v := range iv {
        if v%2 != 0 {
            vy = append(vy, v)
        }
    }
    return f(vy...)
}
```

[Go Playground](https://play.golang.org/p/_ugdD3pYjT)

Callback Functions from [Stack Overflow](https://stackoverflow.com/questions/824234/what-is-a-callback-function).

Examples on callback functions: [Go Playground](https://play.golang.org/p/WeDhi1HXkD).

## Closure

***********************************************************************************

Closure is a concept of enclosing a scope of a variable around the variable so that the scope of that variable is limited - so that variable's scope is smaller part of the overall memory access.

Jedi level is to narrow scope as much as possible.

```go

package main

import (
    "fmt"
)

var ix int

func main() {
    fmt.Println(ix)
    ix++
    fmt.Println(ix)
    darkwood() //also increments ix
    fmt.Println(ix)

}

func darkwood() {
    fmt.Println("Darwood rules!")
    ix++
}

```

Scope of x when it's narrowed down to a func main:

```go
package main

import (
    "fmt"
)


func main() {
    var ix int
    fmt.Println(ix)
    ix++
    fmt.Println(ix)
    darkwood() //also increments ix
    fmt.Println(ix)

}

func darkwood() {
    fmt.Println("Darwood rules!")
    ix++ //will cause an error if we run it! ix is now only in func main
}

```

Code block inside of another code block with iy:

```go
package main

import (
    "fmt"
)


func main() {
    var ix int
    fmt.Println(ix)
    ix++
    {
        iy := 21

        fmt.Println(iy)
    }

    fmt.Println(iy) //This will also cause an error!!!
    fmt.Println(ix)
    darkwood() //also increments ix
    fmt.Println(ix)

}

func darkwood() {
    fmt.Println("Darwood rules!")
    ix++ //will cause an error if we run it! ix is now only in func main
}

```

Using closure - keeping our scope as narrow as possible and limiting the scope of oy.

```go
package main

import (
    "fmt"
)


func main() {
    v := incrementor()
    fmt.Println(v())
    w := incrementor()
    fmt.Println(v())
    fmt.Println(v())
    fmt.Println(v())
    fmt.Println("------------")
    fmt.Println(w())
    fmt.Println(w())
    fmt.Println(w())

}

func incrementor() func() int {
    var ix int
    return func() int {
        ix++
        return ix
    }
}

```

Returning a function that is encloused around that ix. Both w and v are returning different scopes - they have different locations in memory. This is a nice way to use closure in order to limit our scope - and in this case we use it to create two separate variables that are stored in memory and increment those variables.

[Go Playground](https://play.golang.org/p/CK3_UpF5iU)

## Recursion

***********************************************************************************

Recursion is a function that calls itself. Anything we do with recursion we can do with loops - recursion is nice, but it can be unnecessarily complex and also cause us to use unnecessary amounts of memory.

Recursion is effectively when a function calls itself. We looked at closure when we're enclosing a variables scope in a function, callbacks when we have a function calling another functions, and recursion.

Classic example is a factorial.

```go

package main

import (
    "fmt"
)

func main() {
    n := factorial(7)
    fmt.Println(n)
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

// 7 * 6 (7-1) * 5 (6-1) * 4 (5-1) * 3 (4-1) * 2 (3-1) * 1 (2-1) * 0 (defaults to 1)

```

[Go Playground - Factorial](https://play.golang.org/p/hgceXL4f2b)

The function 'unwinds' similiar to the following: ```7 * 6 (7-1) * 5 (6-1) * 4 (5-1) * 3 (4-1) * 2 (3-1) * 1 (2-1) * 0 (defaults to ```.

Once it's unwound, then we exit the function and return the value.

```go

package main

import (
    "fmt"
)

func main() {
    n := factorial(7)
    fmt.Println(n)
}

func factorial(n int) int {

    total := 1
    for v := n; v > 0; v-- {
        total *= v
    }
    return total
}

```

[Factorial with Loop](https://play.golang.org/p/6FBJmet6M_)