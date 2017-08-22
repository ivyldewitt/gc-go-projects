# Grouping Data
***********************************************************************************

Looking at structures that allow us to group the same types together. Or in the case of structs, different types.

Taking values of a specific type and putting them into one data structure. You could also call these aggregate data types.

A composite data type is a type using other primitive types or other composite/aggregate types.

### Arrays
-----------------------------------------------------------------------------------

At its core an array in Go is that it is a building block. For arrays, we specify the size at the start. Arrays in Go (and in most languages) start at a 0-based index.

In most cases, we'd actually want to use ```slices```, not ```arrays```.

[Effective Go](https://golang.org/doc/effective_go.html#arrays)

>"Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation, but primarily they are a building block for slices, the subject of the next section. To lay the foundation for that topic, here are a few words about arrays."

Language Spec

Array Types:

>"An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length and is never negative. The length is part of the array's type; it must evaluate to a non-negative constant representable by a value of type int. The length of array a can be discovered using the built-in function len. The elements can be addressed by integer indices 0 through len(a)-1. Array types are always one-dimensional but may be composed to form multi-dimensional types."

```go

package main

import (
    "fmt"
)

func main() {
    var x [5]int
    var y [7]int
    fmt.Printf("%T\t", x)
    fmt.Prinft("%T\t", y)

    x[0] = 65
    x[3] = 23
    x[1] = 88
    fmt.Println(x)
}

//Output: 
// [5]int
// [7]int
// [65 88 0 23 0]
```

### Slices - Composite Literals
-----------------------------------------------------------------------------------

At its core, a slice holds values of the same type. In this video, we'll use the composite literal to create a slice. Composite literals are created by having the TYPE followed by CURLY BRACES and then putting the appropriate values in the curly brace area.

[Effective Go - Slices](https://golang.org/doc/effective_go.html#slices)

>"Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimension such as transformation matrices, most array programming in Go is done with slices rather than simple arrays. Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array."

```go

package main

import (
    "fmt"
)

func main() {
   //  x := type{values} Syntax
   x := []int{3, 56, 66, 43, 6, 86}
   fmt.Println(x)


}

//A slice allows you to group VALUES of the same TYPE.

```

### Slices - For Range
-----------------------------------------------------------------------------------

We can loop over a slice in the range clause, and access them via the index position.

```go

 package main

import (
    "fmt"
)

func main() {
    //  x := type{values} Syntax
    x := []int{3, 56, 66, 43, 6, 86}
    fmt.Println(len(x)) //6
    fmt.Println(x[0])   //3

    for i, v := range x {
        fmt.Println(i, ":", v)
    }

}
//A slice allows you to group VALUES of the same TYPE.

```
[Playground](https://play.golang.org/p/De0jNGYADF)

### Slices - Slicing a Slice
-----------------------------------------------------------------------------------

```go
package main

import (
    "fmt"
)

func main() {
    w := []int{7, 45, 91, 56, 84, 65, 14}
    fmt.Println(w[1])
    fmt.Println(w)
    fmt.Println(w[1:])
    fmt.Println(w[1:3])

    for i := 0; i < len(w); i++ {
        fmt.Println(i, w[i]) //index, value
    }
}
```
For ```fmt.Println(w[1:3])``` it's important to note that we go up to, but not including that specific position in the slice. We print out 45, 91, but **not** 56.

Note that slicing a slice is also how we delete a slice. If we wanted to take out a number, we could slice different sections and combine it into a new slice.

[Playground](https://play.golang.org/p/p-THBlesrg)

### Slices - Appending a Slice
-----------------------------------------------------------------------------------

Now we have the missing piece we needed to explain the design of the append built-in function. The signature of append is different from our custom Append function above. Schematically, it's like this:

```go
func append(slice []T, elements ...T) []T
```

**where T is a placeholder for any given type.** You can't actually write a function in Go where the type T is determined by the caller. That's why append is built in: it needs support from the compiler.

What append does is append the elements to the end of the slice and return the result. The result needs to be returned because, as with our hand-written Append, the underlying array may change. 

[Append - Effective Go](https://golang.org/doc/effective_go.html#append)

Append returns another slice of the same time.

```go
package main

import (
    "fmt"
)

func main() {
    w := []int{7, 45, 91, 56, 84, 65, 14}
    fmt.Println(w)
    w = append(w, 6, 23, 86, 99, 923)
    fmt.Println(w)
    
    e := []int{62, 44, 73, 332}
    w = append(w, e...)
    fmt.Println(w)

}
```

[Playground](https://play.golang.org/p/loNSZmedIX)

```go
slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...)
```

[GoDoc - Append](https://godoc.org/builtin#append)


### Slices - Deleting a Slice
-----------------------------------------------------------------------------------

We can delete from a slice using both append and slicing. This is a wonderful and elegant example of why Go is great and how Go provides ease of programming.

```go
package main

import (
    "fmt"
)

func main() {
    w := []int{7, 45, 91, 56, 84, 65, 14}
    fmt.Println(w)
    w = append(w, 6, 23, 86, 99, 923)
    fmt.Println(w)
    
    e := []int{62, 44, 73, 332}
    w = append(w, e...)
    fmt.Println(w)
    
    w = append(w[:2], w[5:]...) //... our way of 'unfurling' th data
    fmt.Println(w)

}


```
[Playground](https://play.golang.org/p/TxjVHlvQEk)


### Slices - Make
-----------------------------------------------------------------------------------

Slices are built on top of arrays. A slice is dynamic in that it will grow in size. The underlying array, however, does not grow in size. When we create a slice, we can use the built in function make to specify how large our slice should be and also how large the underlying array should be. This can enhance performance a little bit.

We can use the ```make``` function which takes a slice, length, and capacity. We can use make to create the underlying array that's big enough to use the value in question. The syntax looks like the following:

```go
make([]int, 10, 100)
```

>"Back to allocation. The built-in function make(T, args) serves a purpose different from new(T). It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T). The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use. A slice, for example, is a three-item descriptor containing a pointer to the data (inside an array), the length, and the capacity, and until those items are initialized, the slice is nil. For slices, maps, and channels, make initializes the internal data structure and prepares the value for use. For instance,"

[Effective Go on Make](https://golang.org/doc/effective_go.html#allocation_make)

```go

package main

import (
    "fmt"
)

func main() {
    a := make([]int, 10, 100)
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))

    a[0] = 22

    a[8] = 433
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))

    a = append(a, 88)
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))

   //Note: We can't do something like a[10] = 333 
}
```
Every time we append to a slice, we increase its length. But only the capacity if it's larger than the underlying array. Typically we try to create the largest capacity earlier if we can - in order to prevent us from using up processing power by constantly having to recreate the length and increase the capacity.


### Slices - Multi-Dimensional Slice
-----------------------------------------------------------------------------------

A multi-dimensional slice is like a spreadsheet. **You can have a slice of a slice of some type.**

```go
package main

import (
    "fmt"
)

func main() {
    favLanguages := []string{"golang", "python", "javascript"}
    fmt.Println(favLanguages)
    langsToLearn := []string{"elixir", "scala", "r"}
    fmt.Println(langsToLearn)
    
    combinedLangs := [][]string{favLanguages, langsToLearn}
    fmt.Println(combinedLangs)
    
}
```
[Playground](https://play.golang.org/p/hXvXBI1abe)

### Slices - Underlying Array
-----------------------------------------------------------------------------------

Underlying every slice is an array. A slice is a data structure that which contains three parts:

1. A pointer to an array
2. Len function
3. Cap function

**Ex: New Underlying Array~`

```go
package main

import (
    "fmt"
)

func main() {
    x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
    fmt.Println(x)

    y := append(x, 43, 43, 43, 43, 43, 43, 44) // new underlying array allocated

    fmt.Println(x)
    fmt.Println(y)
}
```

**Ex: Same Underlying Array**

```go
package main

import (
    "fmt"
)

func main() {
    x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
    fmt.Println(x)

    y := append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice

    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(len(y))
    fmt.Println(len(x))
    fmt.Println(cap(x))
    fmt.Println(cap(x))
}

//We're keeping the original, underlying array.
```

Both x & y are pointing to the same underlying array. X is now changed due to this change in the underlying array.

```
Output
[42 43 44 45 46 47 48 49 50 51]
[42 43 47 48 49 50 51 49 50 51]
[42 43 47 48 49 50 51]
7
10
10
10
```

[Playground](https://play.golang.org/p/EB5BAVmY3u)

### Map - Intro
-----------------------------------------------------------------------------------

Maps are a key-value store. This allows us to store values based on a key, and allow very fast and efficient lookup. For example, a map could be a phone book. The key is the name, and the phone number is the value.

Additionally, maps are unordered lists.

[Maps - Language Spec](https://mail.google.com/mail/u/0/#search/starr/15de76569ec316b1)

```go
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
```

>"Maps are a convenient and powerful built-in data structure that associate values of one type (the key) with values of another type (the element or value) The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. Slices cannot be used as map keys, because equality is not defined on them. Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller."

```go

package main

import (
    "fmt"
)

func main() {
    //composite literal
    m := map[string]int{
        "Toby":  7,
        "Macho": 9,
        "Cato":  2,
    }
    fmt.Printf("%T\t", m)
    fmt.Println(m)
    fmt.Println(m["Macho"])
    fmt.Println(m["Stormy"]) //returns 0 as it doesn't exist in the map

    v, ok := m["Stormy"]
    fmt.Println(v)
    fmt.Println(ok)

    if v, ok := m["Toby"]; ok {
       fmt.Println("This is a value in the map: ", v)
    }
}

```

The above is called the comma, OK idiom.

[Playground](https://play.golang.org/p/Z81jgvPUqt)

### Map - Adding Elements and Ranged
-----------------------------------------------------------------------------------

Add an element to a map, use the range loop to print out a map's keys & values.

```go
package main

import (
    "fmt"
)

func main() {
    //composite literal
    m := map[string]int{
        "Toby":  7,
        "Macho": 9,
        "Cato":  2,
    }
    fmt.Printf("%T\t", m)
    fmt.Println(m)
    fmt.Println(m["Macho"])
    fmt.Println(m["Stormy"]) //returns 0 as it doesn't exist in the map

    v, ok := m["Stormy"]
    fmt.Println(v)
    fmt.Println(ok)

    if v, ok := m["Toby"]; ok {
        fmt.Println("This is a value in the map: ", v)
    }

    //adding an element to a map

    m["Misty"] = 6

    for i, v := range m {
        fmt.Println(i, ":", v)
    }

}
```

### Map - Delete
-----------------------------------------------------------------------------------

You delete an entry from a map using ```delete(<map name>, “key”)```. No error is thrown if you use a key which does not exist. To confirm you delete a key/value, verify that the key/value exists with the comma ok idiom.

```go
package main

import (
    "fmt"
)

func main() {
    //composite literal
    m := map[string]int{
        "Toby":  7,
        "Macho": 9,
        "Cato":  2,
    }

    for i, v := range m {
        fmt.Println(i, ":", v)
    }
    fmt.Println("---------")
    delete(m, "Cato")

    for i, v := range m {
        fmt.Println(i, ":", v)
    }

}

```