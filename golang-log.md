# Course Reflections
*************************************

Review & reflections from each day of learning to code.

**D1: Monday, July 31st, 2017**

**What did I learn today?**:

Today, I started v2 of the Learn How to Code course. So far, I've covered the following topics:

* Learning the basics of the Go playground and its uses for debugging/sharing code.
* Review of the basics of a Go program - including package main, func main, how imported packages work, and the distinction of visible/non-visible packages. 
* A quick recap of Go Fmt, and go's formatting rules.
* How we use the short declaration operator, and some key terms such as expression, statement, operand and operator.
* The var keyword, which we can use to declare variables at the package level, instead of within a specific function's scope.
* Assigning values with the var keyword, and how we can declare a value on the same statement as var, but we can't declare ```var x int``` and then ```x = 10```.
* Types - Go is statically typed, not dynamic like JS/Python. That means when you declare a type as an int, string, etc. that's exactly what it is and it can't be changed. I also reviewed primitive data types (ints, floats) vs composite types (structs, arrays, etc).

**How does this build on from yesterday?**:
It doesn't! Technically this is my first day. Though I did provide an explantation to someone else yesterday. I added that as a separate answer in my 01 folder. 

That was a lot of fun. I won't always know the answer, but it's nice when I do!

**Key Takeways & Points to Review**:
* Types - primitive vs composite
* Structure of a Go program - package main, func main, exported/non-exported functions, etc.
* Variable declarations - with var and :=

**Any Reflections?**:
Not anything of note just yet - it's weird at the moment to take the same course I took before so soon, but I feel like I have a much better plan on how to go about it so I'm excited for that. That enough is a great motivator.


**D2: Tuesday, August 1st, 2017**

**What did I learn today?**:

Second day of v2 of the Learn How to Code course. So far I've completed Section 2 on variables, values, and type and learned the following topics:

* How Go implements the "zero value" across its various types.
* The concepts of static typed languages - i.e. that a type can't be converted from one format to another.
* The basics of how packages work by review the "fmt" package - including documentation, exported functions and types.
* Also, how string formatting works with "%T", "%v".
* Creating types - aka how we can alias types with type declarations. Typically done wih structs as a general rule of thumb.
* Type conversions - how to convert one type to another type. Including converting a declared type to its underlying type.

**How does this build on from yesterday?**:
Yesterday I spent most of my time reviewing/learning about the basics of variables - including the short declaration operator, var, the basic structure of a Go program, and a short intro to types. All in all nothing new, but definitely a lot of useful info to know. 

**Key Takeways & Points to Review**:
* Types - declarations and conversion
* Structure of packages - fmt - Sprint familt, Print family, Fprint family.
* Zero values

**Any Reflections?**:
Mostly more review - now onto the exercises!


**D3: Wednesday, August 3rd, 2017**

**What did I learn today?**:

Shorter day today - mostly just me starting to enjoy my vacation. =)

* I completed the section 3 exercises of the 'learn how to code' course - focusing on zero values and type declarations
* I also learned the basics of using 'bool' - and how it's used to express true/false values
* The comparison operators, and how they can also be used to express true/false values

**How does this build on from yesterday?**:
Yesterday I learned about the zero value, type, the fmt package family (print, sprint, and fprint), and type declarations/creating types.


**Key Takeways & Points to Review**:
* Type: Bool - A way to evaluate true/false statements

**Any Reflections?**:
Completed exercises, now onto more lessons.

**D4: Thursday, August 3rd, 2017**

**What did I learn today?**:

Lots and lots of information about computer fundamentals and the intro to numeric types and string types.

* Coding schemes, including ASCII and Unicode systems
* The basics of circuitry/transistors - how computers run on electricty and on/off states
* Review of the intros and evolutions of computers - vacuum tubes, silicon wafers, microprocessors, etc.

* Numeric Types - ints and float64s, uints and ints, and the different 'levels' - int8, int32, int64, etc.
* Bytes - alias for uint8. Runes - alias for int32.
* Unsigned ints vs ints: unsigned means it starts at 0 and up, regular ints go from negative to positive.

* String Types - a string is effectively a slice of bytes, which are an alias for an uint8.
* Because of this, we can use a variety of string formatting methods to reveal the hexidecimal codes, unicode, and various other formats of the strings.


**How does this build on from yesterday?**:
Yesterday I mostly focused on the variable exercises, and learned about bool types - which are a way to eval true/false statements.


**Key Takeways & Points to Review**:
* Type: String - a slice of bytes. We can format strings down to their unicode representations.
* Type: int and float64 - a review of numeric types, including signed/unsigned ints.
* Computer basics - transitors/circuits, coding schemes, and Moore's Law.

**Any Reflections?**:
Reviewing the basics of computers and how they operate/work according to different coding formations and schemes. Overall, the basics are useful for knowing the underlying usefulness/points of computers.

Regarding typing, understanding how go types relate to each other is useful. Additionally, if I remember that one of the benefits of Go is that it is 1) statically typed and b) a compiled language, and the latter means we can compile it down to machine code, understanding how strings ultimately translate to unicode characters, which ultimately represent coding schemes is useful.

**D5: Sunday, August 6th, 2017**

**What did I learn today?**:

I learned specifically about numeral systems, constants, and iota.

* Iota - The representation of a number constant. Effectively, using Iota will start you at 0, and increment by one until you declare a new const.

* Numeral systems - hexadecimal, binary, decimal. We use numeral systems to keep track of numbers. The most 'common' one is the base10/decimal system, and then we obviously have the binary system for computers.

* Const - a keyword for constant values. They cannot be changed after being declared. We can have untyped (where we let the compiler determine the type) and typed constants.


**How does this build on from yesterday?**:
A few days ago I learned mostly about typing - string types, number types, etc.

**Key Takeways & Points to Review**:
* Const - a constant value, cannot be changed after being declared.
* Numeral systems - ways we determine number types, including decimal, hexidecimal, etc.

**Any Reflections?**:
Nothing to world-ending at the moment as I'm familar with Numeral Systems now. The idea of a const also is very intuitive, so for the most part isn't too difficult to get my head around.

**D6: Monday, August 7th, 2017**

**What did I learn today?**:

I finished up the section on programming fundamentals with a recap on bit shifting - aka what happens when we move a binary digit to the left or right.

I also started a few exercises for the programming fundamentals section.

**How does this build on from yesterday?**:
Yesterday, I learned about constant values, numeral systems, and iotas (numbered constants).

**Key Takeways & Points to Review**:
* Bit Shifting - moving a binary digit from left to right.

**Any Reflections?**:
I will have to review the bit shifting section again, because at the moment I don't really quite get why it's useful? I see that it's useful for assembly programming, so at a base level my assumption is that this is useful for low-level programming. It also looks like it has applications in cryptography (this doesn't surprise me though).

Some links for reference:

1. [Bit Hacking in Go](https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827)

2. [Bitwise Shift](https://msdn.microsoft.com/en-us/library/f96c63ed.aspx)

3. [Bit Shifting - Technopedia](https://www.techopedia.com/definition/26846/bit-shifting)

4. [Stack Overflow](https://stackoverflow.com/questions/141525/what-are-bitwise-shift-bit-shift-operators-and-how-do-they-work#141873)