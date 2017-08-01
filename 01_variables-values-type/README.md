# Variables, Values, and Type
*****************************************************

### Playground

- Go Playground: A way to play around with Go w/o sending it to anyone, and a way to send it to each other without hassle. Useful for having others debug and share/send answers!
- Packages: In Go, code is organized into packages, which we *import* into the program. We import using the following syntax:

```go
import (
    "package name"
)
```

- We import functions and types from packages, typically in the following format:

```go

func main() {
    fmt.Println("Hello go code!")
}

```

Using formatting allow us to write idomatic Go code, which we all can use.

### Barebones of a Go Program
------------------------------------------------------------------

Short intro to how a Go program is set up. 

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
}
```

Use ```package main``` which inclues func main() and is the entry/exit point of your program. You can see that package is one of the reserved keywords per the la language spec [here](https://golang.org/ref/spec#Keywords).

We can see how packages work from using godoc.org. For example, if we check out the [fmt package](http://godoc.org/fmt), we can see that there's a number of functions with capital letters. That means they're exportable. 

In Go, we call things exported/non-exported outside the package. For example.

```go 
func main() {
    n, err := fmt.Println("blah")
    _, _ = fmt.Println(n, err) //Note: _ is a 'throwaway"
}
```

Is exported. It returns an int and error, but we don't usually do anything with that. Per the language spec, there's the description for fmt.Println.

>"Println formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered."

Go also prevents/asks us to not use polluted code. We can use "_" to throwaway the code instead of declaring an unsued variable.

*Note: Practically ever function returns an empty interface! Don't worry about this quite just yet*

Key Takeaway: Every program in go starts in  ```package main``` and it uses ```func main()``` as the entry point to your program. It's also the exit point of the program. Any function that is exportable from the program is capitalized.

Short Declaration Operator
------------------------------------------------------------------

The short declaration operation is used to assign values in Go.

```go
package main

import (
    "fmt"
)

func main() {
    x := 15 + 3 //Delcaring a variable
    fmt.Println(x)
    x = 4      //Reassignment
    fmt.Prinltn(x)
}
```

**Terminology**

*Keywords:* Reserved words in Go - review the Golang language spec for the special keywords. These can only be used in specific situation. 
*Operator:* "+, :=, etc." These are characters that represent actions such as additional, assignment, etc.
*Operand:* In "2 + 2" the numbers/values are the operands.
*Statement:* A small, standalone element that expresses an action to be carried out. Statements can have multiple expressions. ```x := 15 + 2```
*Expression:* A combination of values, constants, variables, etc. that evaluates to something. ```15 + 2```

We need to reassign variables with =, not declare them with := after we've first declared/created them. 

The Var Keyword
------------------------------------------------------------------

Let's say we want to have a variable that's accessible within a wider scope. Scope is the area where a variable lives, exists, and can be accessed. In the following program:

```go

package main

var s = 's'

import (
    "fmt"
)
func main() {
    //Scope of x & y is only within func main()
    x := 26
    y := "Beyonce"
    fmt.Println(y)
    fmt.Println(s)
}

func solange() {
    fmt.Println(y) //This won't work!
    fmt.Println(s)
}
```
We can create scope which is larger than a code block to create a scope a the package level.

Still, best practice is to keep scope as *narrow* as possible. Use short declaration operator to declare code within your code blocks.

Exploring Types in Go
------------------------------------------------------------------

>"Type is life" - Bill Kennedy

Part of this is due to the fact that Go is a statically typed language, it is not dynamic like JavaScript or Python. When in Go you *declare* a variable to hold a *value* of a certain *type*, then that variable can only hold values of that type.

We can declare variables with the var keyword or with the short declaration operator. We can also directly declare the value.

```go
package main

import (
	"fmt"
)

// DECLARE that the variable with the IDENTIFIER "z" is of TYPE int
var z int = 21

func main() {
	fmt.Println(z)
}
```

We can declare the type itself without having the compiler determine it instead.

Important note: We can't assign a value to z outside of a function body using the "=" so we can do the above ^ but not the following: 
```go
package main

import (
	"fmt"
)

var z int
z = 21 // Warning! Won't work

func main() {
	fmt.Println(z)
}
```
#### Primitive vs Composite Data Types:

[Wiki Link](https://en.wikipedia.org/wiki/Primitive_data_type)
In CompSci, a primitive data type is either a basic type that is a building block or a built-in type in which the programming language provides built-in support. Many languages allow composite types to be built from built in types.

In CompSci, a compositive data type allows us to compose/aggregate other values of a certain type.

Note: I think structs could be an example of this? Along with arrays and lists. 

The act of constructing a composite data type is the art of composition. 