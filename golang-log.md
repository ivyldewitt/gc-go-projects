# Course Reflections

*************************************

Review & reflections from each day of learning to code.

## D1: Monday, July 31st, 2017

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
It doesn't! Technically this is my first day. Though I did provide an explanation to someone else yesterday. I added that as a separate answer in my 01 folder. 

That was a lot of fun. I won't always know the answer, but it's nice when I do!

**Key Takeaways & Points to Review**:

* Types - primitive vs composite
* Structure of a Go program - package main, func main, exported/non-exported functions, etc.
* Variable declarations - with var and :=

**Any Reflections?**:
Not anything of note just yet - it's weird at the moment to take the same course I took before so soon, but I feel like I have a much better plan on how to go about it so I'm excited for that. That enough is a great motivator.

## D2: Tuesday, August 1st, 2017

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

**Key Takeaways & Points to Review**:

* Types - declarations and conversion
* Structure of packages - fmt - Sprint family, Print family, Fprint family.
* Zero values

**Any Reflections?**:
Mostly more review - now onto the exercises!

## D3: Wednesday, August 3rd, 2017

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

## D4: Thursday, August 3rd, 2017

**What did I learn today?**:

Lots and lots of information about computer fundamentals and the intro to numeric types and string types.

* Coding schemes, including ASCII and Unicode systems
* The basics of circuitry/transistors - how computers run on electricity and on/off states
* Review of the intros and evolutions of computers - vacuum tubes, silicon wafers, microprocessors, etc.

* Numeric Types - ints and float64s, uints and ints, and the different 'levels' - int8, int32, int64, etc.
* Bytes - alias for uint8. Runes - alias for int32.
* Unsigned ints vs ints: unsigned means it starts at 0 and up, regular ints go from negative to positive.

* String Types - a string is effectively a slice of bytes, which are an alias for an uint8.
* Because of this, we can use a variety of string formatting methods to reveal the hexadecimal codes, unicode, and various other formats of the strings.

**How does this build on from yesterday?**:
Yesterday I mostly focused on the variable exercises, and learned about bool types - which are a way to evaluate true/false statements.

**Key Takeways & Points to Review**:

* Type: String - a slice of bytes. We can format strings down to their unicode representations.
* Type: int and float64 - a review of numeric types, including signed/unsigned ints.
* Computer basics - transitors/circuits, coding schemes, and Moore's Law.

**Any Reflections?**:
Reviewing the basics of computers and how they operate/work according to different coding formations and schemes. Overall, the basics are useful for knowing the underlying usefulness/points of computers.

Regarding typing, understanding how go types relate to each other is useful. Additionally, if I remember that one of the benefits of Go is that it is 1) statically typed and b) a compiled language, and the latter means we can compile it down to machine code, understanding how strings ultimately translate to unicode characters, which ultimately represent coding schemes is useful.

## D5: Sunday, August 6th, 2017

**What did I learn today?**:

I learned specifically about numeral systems, constants, and iota.

* Iota - The representation of a number constant. Effectively, using Iota will start you at 0, and increment by one until you declare a new const.

* Numeral systems - hexadecimal, binary, decimal. We use numeral systems to keep track of numbers. The most 'common' one is the base10/decimal system, and then we obviously have the binary system for computers.

* Const - a keyword for constant values. They cannot be changed after being declared. We can have untyped (where we let the compiler determine the type) and typed constants.

**How does this build on from yesterday?**:
A few days ago I learned mostly about typing - string types, number types, etc.

**Key Takeways & Points to Review**:

* Const - a constant value, cannot be changed after being declared.
* Numeral systems - ways we determine number types, including decimal, hexadecimal, etc.

**Any Reflections?**:
Nothing to world-ending at the moment as I'm familiar with Numeral Systems now. The idea of a const also is very intuitive, so for the most part isn't too difficult to get my head around.

## D6: Monday, August 7th, 2017

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

## D7: Tuesday, August 8th 2017

**What did I learn today?**:
I focused on exercises today - woohoo!

**How does this build on from yesterday?**:
Yesterday I read up on bit shifting, which I admittedly still don't get, but will re-read the articles and let some of the concepts sink in.

**Key Takeways & Points to Review**:
Nothing of review today outside of the quiz + exercises. I did forget a few of the GB-TB-Byte conversions though, so a good thing to re-read later.

**Any Reflections?**: 

I think my log is helping me to retain my memory of day-to-day concepts, though it's obviously best if I keep at it daily instead of skipping days. Overall, its going well.

## D8: Thursday, August 10th, 2017

**What did I learn today?**:

I started the section on conditional logic and control flow, starting with the for loop - which in Go services the same functions as the for loop, while loop, and do-while loop.

I also learned about nested loops, the modulus operator (though that was a bit random) and the break and continue statements.

* Modulus - Provides the remainder of two numeric types.
* Break - The way we exit out of the program if we hit a specific condition.
* Continue - The way to skip a section of the program and continue the loop.

**How does this build on from yesterday?**:

Hmm, it builds on the general knowledge from programming fundamentals as we're now taking those building blocks - types, assigment, coding schemes, etc. - and bringing them into the larger flow of programmatic logic.

**Key Takeways & Points to Review**:

* For Loop Syntax
* Break/Continue

**Any Reflections?**:

I'm not unfamiliar w/the for loop syntax, but it's good to take a deeper second look into some of these things.

## D9: Monday, August 14th, 2017

**What did I learn today?**:
Review of more control flow basics - a greater understanding of loops, modulus, and conditionals. In particular, using conditionals with if-else.

**How does this build on from yesterday?**:
Well... it's been a few days, but it builds on from learning how to use iterations (loops) to conditionals (if-else).

**Key Takeways & Points to Review**:

* Conditionals allow us to better segment behavior based on different sequences.
* Using conditions with iterative declarations allow us to repeat a sequence of actions multiple times.

**Any Reflections?**:

This is useful in grounding me to the basics of conditionals and how those can be used to have a specific sequence of actions.

## D10: Tuesday, August, 15th 2017

**What did I learn today?**:
Completed the control flow section - review of the switch statement and a deeper dive into the documentation. Also a look into conditional operators with boolean statements.

Switch Statements come in two types:

* expression statements
* type statements

**How does this build on from yesterday?**:
Now I can put together all of the main points of control flow - if-else, switch, etc. Now I understand how the sequential processes go, and also how I can use iterative statements (loops) and conditionals (if & switch).

**Key Takeways & Points to Review**:

* Conditionals allow us to better segment behavior based on different sequences.
* Using conditions with iterative declarations allow us to repeat a sequence of actions multiple times.
* The Switch statements can be used with fallthrough, default cases, or without an expression at all.

**Any Reflections?**:

This is useful in grounding me to the basics of conditionals and how those can be used to conduct a specific sequence of actions. In conjunction with iterative statements, I can see how we can 'break' out of the standard sequence of a program by breaking it using other statements.

## D11: Wednesday, August 16th 2017

**What did I learn today?**:
Control flow is done baby! Clearly I'm listening too much to Todd lol. But I've completed exercises 1-5 out of 10.

**How does this build on from yesterday?**:
Now I'm putting those skills from the previous section to use. I just need to make sure I re-read the exercise description though - I probably get tripped up on not understanding the question more than anything.

**Key Takeways & Points to Review**:

This is from yesterday, but I think it's still good to review.

* Conditionals allow us to better segment behavior based on different sequences.
* Using conditions with iterative declarations allow us to repeat a sequence of actions multiple times.
* The Switch statements can be used with fallthrough, default cases, or without an expression at all.

**Any Reflections?**:

The exercises are fun - challenging enough to make me think and keep me engaged.

## D13: Saturday, August 19th 2017

**What did I learn today?**:
Note - technically I did do some exercises on 8/17 so I'll count that as a day even though I didn't write anything down for that. 

I made it through the slices portion of the grouping_data section which included:

* Initializing a slice with make
* Creating slice literals
* Using the append function to add onto slices
* Brief review of arrays, which we typically abandon in favor for slices
* Understanding how the underlying array concept works for slices
* Deleting slices using append & slicing
* Slicing a slice with the [:] feature
* Multi-dimensional slices

**How does this build on from yesterday?**:
I feel like I actually have a significantly better understanding of slices, even in comparison to Go 1.0. I mostly need to review a few other things, and think about the basics of how slices can be used in a variety of situations when it comes to group similar data.

**Key Takeways & Points to Review**:

* The len() and cap() functions are part of every slice, but particularly essential when creating slices with make().
* Changing and altering a slice with append changes the underlying array within the slice.

**Any Reflections?**:

Not too much at the moment, but things are well. It's steady, and I don't have a greater idea of what's going at the moment.

## D14: Sunday, August 20th, 2017

**What did I learn today?**:
Completed the section on maps, and the ninja level exercises for the section as well.

I learned about deleting map data type, creating a composite literal map data type and adding to a map data type.

**How does this build on from yesterday?**:
I feel like I actually have a significantly better understanding of slices & maps, even in comparison to Todd's 1.0 Go course. Slices allow us to create lists of a single data type, which maps allow us to create slightly more complex variations on data types - ex: I can create a map with an int as the key, but a slice of bytes as the value.

Why would I do that? Eh? But it's possible with maps.

**Key Takeways & Points to Review**:

* You can use a range for loop *inside* another range for loop in order to loop over the inner elements of a map if you're using another data type as the value - ex: a slice.
* The basics of maps are relatively straightforward, and I can see why you'd used maps (ex: phone book) vs slices (ex: shopping list) vs arrays (use slices).

**Any Reflections?**:

The end of the video kind of drug on a bit, but I feel a lot more comfortable using maps along with a quick review on the exercises to strengthen my overall knowledge.

## D15: Monday, August 21st 2017

**What did I learn today?**:

Completed the entire section on structs! Woohoo, surprised I actually banged it out in one sitting. Overall, it was a really useful section with a great review period afterwards. I've learned the following about structs so far:

* Structs are our way of creating aggregate data types - i.e. types that can hold a mix of other types.
* We can create structs using the ```struct{}``` denotation, and add the fields and list them by type.
* We can create anon structs, which are useful for establishing clean code.

**How does this build on from yesterday?**:

I completed the exercises on grouping data with maps, arrays, and slices - so now I'm moving on from creating composite data types of a single type, to composite data types of multiple types.

**Key Takeways & Points to Review**:

* Structs allow us to combine multiple types into a single data type.
* They are *not* objects/classes, but similar in nature.

**Any Reflections?**:

Nothing too much, overall this is useful for better understanding the 'why' of using various data types, including the composite/aggregate ones such as maps, structs, slices, and arrays.

The review/housekeeping section at the end also brought up quite a few points in Go - it's not OOP, structural alignment, and how Go is all about types.


## D16: Tuesday, August 22nd 2017

**What did I learn today?**:

Completed the struct exercises - and learned quite a bit from the challenges! Was having difficulties of my own trying to implement the various technicalities and nuances of the other complex/aggregate data types - slices, maps, with structs. It took me a while to fully grasp how all of those types would correlate together, but this was a great mental exercise.

**How does this build on from yesterday?**:

I did breeze through the structs section, but I don't feel like that information was lost on me - I was able to successfully implement and understand the basics of structs, but also how various data types could be integrated into it. I don't recall, but previously I had only really be working w/primitive data types in structs.

**Key Takeways & Points to Review**:

* Looping/iterating over structs.
* Using composite data types inside of structs.
* Using anon structs - with bonus work using composite data types!

**Any Reflections?**:

Considering I'm working on a blog post on types I'm more than sure this will be useful for a variety of occasions.

## D17: Wednesday, August 23rd 2017

**What did I learn today?**:

Technically not 100% part of this course, but I did finish part 1 of my post on an intro to types in Golang. Part 1 covers less Go syntax and more on general dynamics of types, but still useful to know overall.

I also began the section on functions in Go - basically this provides a quick overview of functions - including function syntax, variadic functions, and 'unfurling' (aka passing over the values in a slice within a function) variadic functions.

**How does this build on from yesterday?**:

Truthfully, writing my blog post slowed me down a bit, but that was also extremely useful in better understanding the differences between static vs dynamic typing, type checking, and strong/weak typing.

From what I learned today specifically? Understanding the key points of functions in Go is useful because Go is **all about type**. And from my experience, you can end up with a lot of typing errors in functions if you're not careful.

**Key Takeways & Points to Review**:

* Functions have the following parts: func, receiver, identifier, params, returns, and code
* We can implement variadic parameters with ... (such as ```x ...int```)
* We pass in variadic arguments with the following syntax: ```xi...```

**Any Reflections?**:

I think in my earlier knowledge of Go I didn't quite understand why variadic functions were useful - I knew the gist of it, but wasn't really putting it together. Now with my understanding of the versatility and usefulness of slices, I understand that it's a quick, useful, and essential shorthand for passing in a range of values of a specific slice type.

Link: [An Introduction to Types in Go Part 1](https://zentechnista.github.io/2017/08/an-intro-to-types-in-go-part-1/)

## D18: Friday, August 25th 2017

**What did I learn today?**:

Deep dive into the defer keyword, understand how methods work with types, and a *really* intricate look into interfaces and polymorphism.

**How does this build on from yesterday?**:

Yesterday (well Wednesday), I focused on the basics of functions, including the syntax and application. Now I'm seeing how functions are used on an a more practical level, especially interfaces and polymorphism.

**Key Takeways & Points to Review**:

* Polymorphism: In Go, this refers to the fact that a method can be attached to many types, and it can take on different forms based on that type.
* Interfaces: Allow us to implement methods that can be used by a variety of types. We can see this with some interfaces such as the io.Writer interface, which allows us to implement methods from many packages including the net/http package.
* Empty Interfaces: All types have at least *no* methods. This means that all types are also type empty interface.
* Methods are functions attached to types.
* The defer statement allows us to defer a function to the end of the execution of the main program.

**Any Reflections?**:

It's a lot to take in, but I have confidence that I can do it! Wish me luck.

## D19: Tuesday, August 29, 2017

**What did I learn today?**:

A deeper review into functions including anonymous functions and function expressions, which is good since I'm still in the 'what does all of the this mean???' section of functions and interfaces. I think I'm getting better at being more intuitive on some of the major parts though.

**How does this build on from yesterday?**:

I took a few days off of this project - for good reason though, I was learning JavaScript instead. But I now am coming back to this with a more rounded and understanding perspective in my opinion - I think that I can better comprehend what's going on here and what I should be focusing on.

**Key Takeways & Points to Review**:

* Function Expressions: Assiging a function to a variable.
* Returning functions in functions: Functions are types, just like all other types in Go and we can return a function in another function. My hypothesis to why this might be useful is in terms of trying to manage our scope.
* Anon Functions: Exactly what it says on the tin - it's an anon function.

**Any Reflections?**:

Still reviewing, but I think I'm working my way up, and luckily I'm doing better at pulling in my knowledge from the previous days and better understanding what's going on.