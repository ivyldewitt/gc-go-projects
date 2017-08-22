# Structs
***********************************************************************************

A struct is a data strcuture which allows us to compose together values of different types. It's an aggregate data type.

### Structs
-----------------------------------------------------------------------------------

```go

package main

import (
    "fmt"
)

type dog struct{
    name string
    breed string
    fourLegs bool
}

func main() {
    d1 := dog{
        name: "Toby",
        breed: "Pomeranian-Sheltie Mix",
        fourLegs: true,
    }

    d2 := dog{
        name: "Macho",
        breed: "Bishon mix?",
        fourLegs: true,
    }

    fmt.Println(d1)
    fmt.Println(d2)

    fmt.Println(d1.name, d1.breed, d1.fourLegs)
    fmt.Println(d2.name, d2.breed, d2.fourLegs)
}

```

Note: We don't say we created a class/object, we say we creates a value of a type. A composite data structure where we're composing values of a different type. Also know as a complex data type or aggregate data type.

### Embedded Structs
-----------------------------------------------------------------------------------

How we can take one type and embed it in another type. Go's approach to creating different data structures and using other data structures.

The inner fields of the example below (*name, breed, fourLegs*) we don't need to reference the dog struct first aka ```cartoonDog.dog.name```. Those fields were promted to cartoonDog.

```go

package main

import (
    "fmt"
)

type dog struct{
    name string
    breed string
    fourLegs bool
}

type cartoonDog struct {
    dog
    catchphrase string
}

func main() {
    d1 := dog{
        name: "Toby",
        breed: "Pomeranian-Sheltie Mix",
        fourLegs: true,
    }

    d2 := dog{
        name: "Macho",
        breed: "Bishon mix?",
        fourLegs: true,
    }

    scooby := cartoonDog{
        dog: dog{
            name: "Scooby Dog",
            breed: "Mutt",
            fourLegs: true,
        },
        catchphrase: "Ruh-roh!"
    }

    fmt.Println(d1)
    fmt.Println(d2)
    fmt.Println(scooby)

    fmt.Println(d1.name, d1.breed, d1.fourLegs)
    fmt.Println(d2.name, d2.breed, d2.fourLegs)
    fmt.Println(scooby.name, scooby.breed, scooby.fourLegs, scooby.catchphrase)
}

```

[Playground](https://play.golang.org/p/WaTFBpLQj8)

### Reading Documentation
-----------------------------------------------------------------------------------

Golang is a language about types. 

[Golang Lang Spec - Structs](https://golang.org/ref/spec#Struct_types)

From the language spec:

>"A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (AnonymousField). Within a struct, non-blank field names must be unique."

```
StructType     = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl      = (IdentifierList Type | AnonymousField) [ Tag ] .
AnonymousField = [ "*" ] TypeName .
Tag            = string_lit .
```

```go
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
    //Note: if you have variables of the same type you can declare w/commas
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}


type dog struct{
    // All of the below are field declarations
    name string //explicit naming
    //name string - not allowed, must be unique
    breed string
    fourLegs bool
}

type cartoonDog struct {
    dog         //anonymous field
    catchphrase string
}
```

>"A field or method f of an anonymous field in a struct x is called promoted if x.f is a legal selector that denotes that field or method f.Promoted fields act like ordinary fields of a struct except that they cannot be used as field names in composite literals of the struct."

### Anon Structs
-----------------------------------------------------------------------------------

An anonymous struct is a struct which is not associated with a specific identifier. Effectively, the struct ```person``` is representing those fields. If we remove person and just create a new variable and declare the struct fields within it. 

We would do this to keep our code clean - only using our code once.

```go
package main

import (
    "fmt"
)


func main() {
    p1 := struct {
    first string
    last string
    age int
    },{
        first: "John",
        last: "Doe",
        age: 25,
    }

    fmt.Println(p1)
}

```


### Housekeeping
-----------------------------------------------------------------------------------

The Main Takeaway: We can use primitive types in code, and use the struct aggregate data type to allow us to create our own type. Go's goals: efficent compilation, efficient execution, and ease of programming.

Note that we don't don't talk about public/private - instead we use exported/unexported.

In Go, we don't create classes, we instatiate a type.

**It’s all about type**
* Is go an object oriented language? Go has OOP aspects. But it’s all about TYPE. We create TYPES in Go; user-defined TYPES. We can then have VALUES of that type. We can assign VALUES of a user-defined TYPE to VARIABLES. Anecdote: makes me think of that song, “It’s all about the bass, all about the bass” except “it’s all about the TYPE, all about the TYPE”

**Go is Object Oriented**
* Encapsulation
    * state ("fields")
    * behavior ("methods")
    * exported & unexported; viewable & not viewable
* Reusability
    * inheritance ("embedded types")
* Polymorphism
    * interfaces
* Overriding
    * "promotion"

**Traditional OOP**
* Classes
    * data structure describing a type of object
    * you can then create "instances"/"objects" from the class / blueprint
    * classes hold both:
        * state / data / fields
        * behavior / methods
    * public / private
*   Inheritance

**In Go:**
* you don't create classes, you create a TYPE
* you don't instantiate, you create a VALUE of a TYPE

**User defined types**
* We can declare a new type
* foo 
    * the underlying type of foo is int
    * int conversion
        * int(myAge) 
        * converting type foo to type int 
* IT IS A BAD PRACTICE TO ALIAS TYPES 
    * one exception: 
        * if you need to attach methods to a type 
        * see the time package for an example of this [godoc.org/time](godoc.org/time) 
        * Example: type Duration int64 
        * Duration has methods attached to it

**Named types vs anonymous types**
* Anonymous types are indeterminate. They have not been declared as a type yet. The compiler has flexibility with anonymous types. You can assign an anonymous type to a variable declared as a certain type. If the assignment can occur, the compiler will figure it out; the compiler will do an implicit conversion. You cannot assign a named type to a different named type.

**Padding & architectural alignment**
* Convention: logically organize your fields together. Readability & clarity trump performance as a design concern. Go will be performant. Go for readability first. However, if you are in a situation where you need to prioritize performance: lay the fields out from largest to smallest, eg, int 64, int64, float32, bool
