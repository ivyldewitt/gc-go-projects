# Application

---------------

In this section, now that we've learned all of the basic methods we're moving into understanding the innerworkings of the standard library. These examples show how we use the standard written library to implement those features into our own application.

## JSON Documentation

---------------

JSON is a way that we can take our data and put it in a certain data structure which is commonly used for transferring data between our application and back. It also shows how the 'case' of that determines how something is visible/non-visible.

[Godoc JSON](https://godoc.org/encoding/json)
[Golang.org JSON](https://golang.org/pkg/encoding/json/)

>"Package json implements encoding and decoding of JSON as defined in RFC 4627. The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions."

We can use the files section to view the 'behind the scenes' code. The examples section gives an overview of what available examples there are.

Example of Marshalling: [Godoc](https://godoc.org/encoding/json#example-Marshal)

In Go, we check the error right in the place where it should occuring.

```go
type ColorGroup struct {
    ID     int
    Name   string
    Colors []string
}
group := ColorGroup{
    ID:     1,
    Name:   "Reds",
    Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}
b, err := json.Marshal(group)
if err != nil {
    fmt.Println("error:", err)
}
os.Stdout.Write(b)
```

Example of Unmarshall:

```go
var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
type Animal struct {
    Name  string
    Order string
}
var animals []Animal
err := json.Unmarshal(jsonBlob, &animals)
if err != nil {
    fmt.Println("error:", err)
}
fmt.Printf("%+v", animals)
```

A slice of bytes is very close to a string - so we'll go from a slice of bytes to a string and vice-versa.

## JSON Marshal

---------------

Marshalling is when we covert our go code into JSON. Also important, this video shows how the case of an identifier - lowercase or uppercase, determines whether or not the data can be exported.

```go
package main

import (
    "fmt"
    "encoding/json"
)

type person struct {
    First string
    Last string
    Age int
}

func main() {
    storm := person{
        First: "Ororo",
        Last: "Munroe",
        Age: 27,
    }

    wolverine := person{
        First: "James",
        Last: "Logan",
        Age: 32,
    }

    people := []person{
        storm,
        wolverine,
    }
    fmt.Println(people)

    bs, err := json.Marshal(people)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(bs))
}

```

[Marshal - Go Playground](https://play.golang.org/p/zt-XiuTu9o)

Interesting Note: If you make the fields uppercase even if you don't make the struct uppercase, the Marshal function can still pull that data out.

## JSON Unmarshal

---------------

We take JSON and transform it back into regular Go code via json.Unmarshal.

JavaScript Object Notation (JSON) is a text format for the serialization of structured data.

It is derived from the object literals of JavaScript.

JSON can represent four primitive types (strings, numbers, booleans, and null) and two structured types (objects and arrays)

JSON objects are returned in name-value pairs.

Example JSON:

```json
{
    "Image": {
        "Width":  800,
        "Height": 600,
        "Title":  "View from 15th Floor",
        "Thumbnail": {
            "Url":    "http://www.example.com/image/481989943",
            "Height": 125,
            "Width":  100
        },
        "Animated" : false,
        "IDs": [116, 943, 234, 38793]
    }
}
```

Notes: [Rawgit - Understanding JSON](https://cdn.rawgit.com/GoesToEleven/golang-web-dev/17e3852d/040_json/README.html)
[JSON to Go](https://mholt.github.io/json-to-go/)

Unmarshal - takes a slice of bytes and returns an error. Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.

```go
func Unmarshal(data []byte, v interface{}) error
```

```go
package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    First string `json:"First"`
    Last string `json:"Last"`
    Age int `json:"Age"`
}

func main() {

    s := `[{"First":"Ororo","Last":"Munroe","Age":27},{"First":"James","Last":"Logan","Age":32}]`

    bs := []byte(s)

    fmt.Println(string(bs))

    //people := []person{}
    var people []Person

    err := json.Unmarshal(bs, &people)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", people)

    for i, v := range people {
        fmt.Println("\nPerson #"i)
        fmt.Println(v.First, v.Last, v.Age)
    }
}
```

[Unmarshal - Go Playground](https://play.golang.org/p/TK3REull0h)

## Writer Interface

---------------

## Sort

---------------

## Sort Custom

---------------

## Bcrypt

---------------