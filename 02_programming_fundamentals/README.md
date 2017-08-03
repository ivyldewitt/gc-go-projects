# Programming Fundamentals
*****************************************************

### Bool Type
------------------------------------------------------------------
Exploring types in more detail - starting with the bool type. Booleans are binary variables representing true/false values.

Important to understand how the people who built this language are talking about the language.

From the [language spec](https://golang.org/ref/spec#Boolean_types):

>"A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false. The predeclared boolean type is bool."

Booleans are important in programming as they allow us to evaluate an expression down to T/F values.

```go
package main

import (
	"fmt"
)

var d bool

func main() {
	fmt.Println(d)
	d = true
	fmt.Println(d)
    a := 7
	b := 42
	fmt.Println(a == b)
}

```

Declaring d as type 'bool'. If we do not declare a value, the default/zero value of bool is false.

*Note: Go is all about ease of programming and efficient execution.*

We can use comparison operators along with our evaluation. See above for the "is equal to" operation where we evaluate an expression. There is no triple equal operator in Go vs JavaScript.

```
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

Single equal is assignment, double equal is comparison.

### How Computers Work
------------------------------------------------------------------