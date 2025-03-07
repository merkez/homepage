> Types are life. -- Bill Kennedy

Go is statically typed. It means that variables always have specific type and the type cannot change during the program run time.

Data types help us reason about what our program is doing and help us catch many errors.

Types classify values into groups and determine:

* what is the memory size allocated for the value
* what does the value represent (e.g. bits `11111111` represent decimal number 255 if the type is `int`)
* what are intrinsic [operations](https://tour.golang.org/methods/16) of that representation (e.g. arithmetic operations for numbers, `len` for strings)

Go's types categories:

* [basic types](https://play.golang.org/p/z5uVUJsKxBw): numbers (integers, floats, complex), strings, booleans
* aggregate types: arrays, structs (user-defined types)
* [reference types](https://play.golang.org/p/NZ0VhQ_pwYR): pointers, slices, maps, functions, channels
* interface types

See also [Basic types](https://tour.golang.org/basics/11), [Zero values](https://tour.golang.org/basics/12), [Type conversions](https://tour.golang.org/basics/13) and [Type inference](https://tour.golang.org/basics/14).
