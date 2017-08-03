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

### Short Declaration Operator
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

### The Var Keyword
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

### Exploring Types in Go
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

### Zero Value
------------------------------------------------------------------

Understanding initialization and the zero value in Golang. Some review about declaring/assiging.

```go

package main

import (
    "fmt"
)

var q int 
//Declaring a variable of type int

func main() {
    q = 53 //Assiging or initialize a value to that variable
    fmt.Println(q)
}

```

We can only assign the type we declared to the variable. Ex: ```q = "Toby"``` would not work.

Initialization: In programming, this is the assignment of an initial value for a data object or variable. *Aka the first time you assign a value to a variable.*

Things in a computer are stored in memory. In a computer we have memory (like a post office with P.O. boxes). We havea all of these values stored in different P.O. boxes. 

Every type in Go has a zero value. If we haven't assigned an initial value, the compiler sets it to the zero value which is:

* False for bool
* 0 for ints
* 0.0 for integers
* "" for strings
* nil for points/funcs/interfaces, slices, channels, and maps.

We should typically use var for package scope, and zero declarations. As a general rule, the short declaration operator should be our first choice.


### The fmt Package
------------------------------------------------------------------

Looking more into how we read the standard library documentation with the fmt pacakge.

```go

package main

import (
    "fmt"
)

var b string = "Beyonce Knowles"
var c bool
var z bool = true

func main() {
    a := 23
    jm := "Electric Lady"
    v := `My favorite Solange song is "Losing You".`
    fmt.Println(s)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(z)
    fmt.Println(e)
    fmt.Println(jm)
    fmt.Println(b, "says: ", v)
    fmt.Println(v)

    fmt.Printf("%v\n", a)
    fmt.Printf("%T\n", a)
    fmt.Printf("%T\n", jm)
    fmt.Printf("%T\t%T", b, z)

    s := fmt.Sprint(a, jm)
}

```

[Playground Link](https://play.golang.org/p/vQ_vAvAEwW)

The above code shows a quick review of raw string literals ``, different variable types and assignments. 

Checking the [Godoc Fmt Index](http://godoc.org/fmt#pkg-index) for list of main functions and variables. To format strings, we'll use:

func Printf
```go
func Printf(format string, a ...interface{}) (n int, err error)
```
It takes a format of type strings, and a unlimited number of values. The format verbs are as follows:

```
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value
```

Fun fact: These are derived from C, but simpler. 

Escaped Characters: See the [rune literals](https://golang.org/ref/spec#Rune_literals) section in the Go Language Spec for reference on common escaped characters like \n or \t.

With format printing we add the format printing string first, and then add the variable to tell the complier what we're trying to format.

**Sprint Family**

Sprint is to string print. We can use this to print directly to a string.

```go
package main

import (
	"fmt"
)

var b string = "Beyonce Knowles"
var c bool
var z bool = true

func main() {
	a := 23
	jm := "Electric Lady"
	v := `My favorite Solange song is "Losing You".`
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(z)
	fmt.Println(jm)
	fmt.Println(b, "says: ", v)
	fmt.Println(v)

	fmt.Printf("%v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", jm)
	fmt.Printf("%T\t%T", b, z)

	s := fmt.Sprint(a, jm)
	fmt.Println(s)
	sk := fmt.Sprintf("%T\t%T", b, z)
	fmt.Println(sk)
}
```
>"Sprint formats using the default formats for its operands and returns the resulting string. Spaces are added between operands when neither is a string."

**Fprintln Family**

Allows us to print to the web ala io.Writer. 

Fprint formats using the default formats for its operands and writes to w. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.


### Creating Your Own Type
------------------------------------------------------------------

>"Go is all about type"

Since Go is a static programming language, i.e. once you declare a variable is of a certain type, it doesn't change.

```go

package main

import (
	"fmt"
)

var a int

type solange int

var b solange

func main() {
	a = 23
	b = 545
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
    // a = b - wont' work!
	fmt.Printf("%T\n", b)
}

//Output
// 23
// int
// 545
// main.solange
```

[Go Playground Link](https://play.golang.org/p/FuwEUkKTt9)


### Conversion, Not Casting
------------------------------------------------------------------

We have a specific way to talk about Go - we *convert* values, not *cast* them. We can convert a value from a *declared type to its underlying type*, or convert a primitive type to another.

```go

package main

import (
	"fmt"
)

var a int

type solange int

var b solange

func main() {
	a = 23
	b = 545
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
    a = int(b)
	fmt.Printf("%T\n", b)
}

//Output
// 23
// int
// 545
// main.solange
```

Extras from Section 2.

Review of types and underlying types: [golang language spec](https://golang.org/ref/spec#Types)

Link to my Quiz Answers for Section 1:[link](https://docs.google.com/forms/d/e/1FAIpQLSfyN4xMJZPoz_2bVy-BXctXfb1a64n4deYF9jj6JLnhpwA3dw/viewscore?viewscore=AE0zAgAui3w12-z7i1VLsfrFq9K_PCm3j6t8ELRNdJKx-79M-5I6edSl0U1dS8k)