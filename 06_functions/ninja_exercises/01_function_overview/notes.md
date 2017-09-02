# Review of Functions

-------------------------------------------------------------------------------------

We learned the syntax of a function:

```go

func (r receiver) identifier(p parameters) r returns {
    //Code goes here
}

```

We learned about variadic parameters and arguments. Basically how we can take in a slice of a specific type as a parameter, and then use variadic arguments to 'uncouple' the slice parameters out of that.

```go

package main

import "fmt"

func main() {
    x := []int{2, 3, 4, 5, 6, 7, 8}

    num := nums(x...)
    fmt.Println(num)
}

func nums(x ...int) {
    total := 1
    for _, v in x {
        total *= v
    }
    return total
}

```

We learned how to use defer to defer the execusion of a function to allow a function to execute on delay.

We also used functions to attach them to a type a create a method. See section on methods for greater review.

The purpose of functions is to create dry code and improve code reusability.

## Func Exxpressions

Assigning a func to a variable.

## Callbacks

Passing a func into another func as an argument.

## Closure

One scope enclosing another, closure helps us limit the scope of varaibles.
Variables declared in the outer scope are accessible in inner scopes.

## Recursion

A function that calls itself - common example is a factorial.

At the end, Go is all about type - all about ease of programming.