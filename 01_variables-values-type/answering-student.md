# Typing in Go

**Summary:**

Q: I have the following program [here](https://play.golang.org/p/Cerg8gIEz4) which is returning a special character instead of the string version of the number here on line 23. Why is that?

A:
For starters, I included my solution here: https://play.golang.org/p/o2rYsEZ8S7.

It looks like the issue you're running in is due to two things:
1. How typing works in Go, and how type declarations work
2. Using the string() functgion

The second one is actually pretty simple, so I'll tackle it first. Using string() is a bit odd, as instead of doing what you would expect and converting an int to a string, it actually converts the signed or unsigned int to the UTF-8 representation of the integer. So for example:
```go
fmt.Println(string(97))
Returns 'a'. In contrast:
fmt.Println(string(892))
```
Returns *'U+037C'*, which is the symbol you're seeing (and fun fact - is a greek symbol from what I'm seeing). 

You can read more about string conversion and the string() from the language spec [here](https://golang.org/ref/spec#Conversions). I also think this post [here](https://www.calhoun.io/5-tips-for-using-strings-in-go-2/) is pretty helpful on the best ways to convert strings. Point 3 in that post also details how it's recommended to use the strconv package to convert data types into strings. Strconv allows us to convert strings to int, floats, runes, etc. and vice verse. Typically, to convert an int to a string we'd use strconv.Itoa() like in the code below:
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    i := 123
    t := strconv.Itoa(i)
    fmt.Println(t)
}
```

But as you can see, strconv allows us to convert strings to existing data types, but not type declarations such as score. If we try to do something like the following:

```go
package main

import (
	"fmt"
	"strconv"
)

var a string

type score int

var b score

func main() {
	a = "John Wick"
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 892
	fmt.Println(b)
	fmt.Printf("%T\n", b) 

	a = strconv.Itoa(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
```

We get an error:
tmp/sandbox317287414/main.go:23: cannot use b (type score) as type int in argument to strconv.Itoa
Ultimately, this is because of how type declarations and conversion work. We can convert a type to its underlying type - for example, if you check out the commented out sections in my code, I did a quick comparison of converting score vs int here:

```go
	var x int = 10
	var z score = score(x)
	fmt.Println(z)
	y := int(z)
	fmt.Println(y)
```

While these two can be converted to each other as score's underlying type is an int, we can't convert score to a string due to it being its own type. The type of b is still main.score, so in order to convert the int to a string we need to do two things: use strconv.Itoa and convert b back to an int in order to use Itoa. You eventually end up with the result below:

```go
package main

import (
	"fmt"
	"strconv"
)

var a string

type score int

var b score

func main() {
	a = "John Wick"
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 892
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	c := int(b)

	a = strconv.Itoa(c)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
//Output:
//John Wick
//string
//892
//main.score
//892
// string
```

In summary, the error above occurs because of how the strings() function works - it returns the Unicode character for a number, not the number itself. We should use the strconv package instead. Additionally, we should alias a primitive type such as string, int, etc. Structs, yes. Primitive types, typically no.