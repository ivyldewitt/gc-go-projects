# Control Flow 
*****************************************************

### Control Flow Overview
------------------------------------------------------------------

>"When computers read code in a program, they read it from top to bottom. This is known as reading it in SEQUENCE. This is also known as SEQUENTIAL control flow. In addition to sequential control flow, there are two other statements which can affect how a computer reads code. A computer might hit a LOOP control flow. If it hits one of these, it will loop over the code in a specified manner. Loop control flow is also known as ITERATIVE control flow. Finally there is also CONDITIONAL control flow. When the computer hits a CONDITION, like an “if statement” or a “switch statement” the computer will take some course of action depending upon some condition. So the three ways a computer reads code are: SEQUENTIAL, LOOP, CONDITIONAL"

Computers read code in 3 specific ways:

* Sequential (sequence/order)
* Loop (iterative)
* Conditional (do something based on a specific condition)

Overviewe of different types of loops/sequences:

* Sequence
* Loop
    * For loop
        * Init, cond, post
    * Bool (while)
        * Infinite
    * Do-While
        * Break
    * Continue
    * Nested
* Conditionals
    * Switch/case/default statement
    * If Statements

### Loop - Init, Condition, Post
------------------------------------------------------------------

The For Loop consists of 3 parts: the initialization, the condition, and the post/increment.

An interesting facet of Go is that there isn't a while loop. We can initialize via the following syntax:

```go
package main

import (
    "fmt"
)

func main() {
    // for init; condition; post {}
    for i := 0; i <= 100; i++ {
        fmt.Println(i)

    }
}

```

We start the for loop, initialize, create the stop condition, and then the 'driver' or post. In Go, all we have is the for loop - no while, do while, etc.

### Nested Loops
------------------------------------------------------------------

We can create a loop inside a loop. Neat thing to know for some programming puzzles.

```go

package main

import (
    "fmt"
)

func main() {
    for i := 0; i <= 10; i++ {
        for z := 0; z < 3; z++ {
            fmt.Printf("The outer loop: %d\t The inner loop: %d\n", i, z)
        }
    }
}
```

### For Statement Documentation
------------------------------------------------------------------

[Effective Go on For Loops](https://golang.org/doc/effective_go.html#for)

>The Go for loop is similar to—but not the same as—C's. It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```

Looping over arrays/slices/maps with For:

```go
for key, value := range oldMap {
    newMap[key] = value
}

// Only 1st item
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}

// Only 2nd item
sum := 0
for _, value := range array {
    sum += value
}
```

```go
package main

import (
    "fmt"
)

func main() {

    // For loop that is like a while loop
    z := 20
    for z > 5 {
        fmt.Println(z)
        z--
    }
    fmt.Println("Done")

    //For alone

    q := 3
    for {
        if q > 13 {

            break
        }
        fmt.Println(q)
        q++
    }
    fmt.Println("Bam")
}
```

[Review](https://play.golang.org/p/1FIbTy8Ijk)

### For - Break & Continue
------------------------------------------------------------------

Break & Continue are reserved words within Go, similiar to other languages.

Break - Breaks out of the loop.

Continue - Continue to the next iteration.

**Remainder in programming w/Modulus:**

```go
package main

import (
    "fmt"
)

func main() {
    i := 43 % 40
    fmt.Println(i)

    //remainder - 3
}

```

**Break and Continue Overview**

```go

package main

import (
    "fmt"
)

func main() {
    x := 1
    for {
        x++
        if x > 100 {
            break
        }

        if x%2 != 0 {
            continue
        }

        fmt.Println(x)
    }
    fmt.Println("done.")
}
```

