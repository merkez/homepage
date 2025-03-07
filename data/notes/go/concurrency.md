Concurrent programming

* a way to structure software
* a composition of several independent computations
* concurrency != parallelism
* inherently more difficult than sequential

Why?

* the world is a complex system of interacting, independently behaving pieces
* we want to model and interact with this world
* the number of processors in a computer grows faster than their speed

Go enables two styles of concurrent programming:

1. Communicating sequential processes (CSP) - see below
2. Shared memory multithreading - more traditional model

Goroutines

* concurrently executing activities
* when a program starts, its only goroutine is the one that calls the `main` function
* you can create new goroutines with the `go` keyword

```
f()     // call f(); wait for it to return
go f()  // create a new goroutine that calls f(); don't wait
```

* a goroutine is stopped by returning from `main` or by exiting the program.

Channels - a way for gouroutines to

* communicate with each other
* synchronize their execution

Select

* `select` statement is like a switch but for channels (i.e. it's not for expressions but for communications)
* first all channels are evaluated
* blocks until one communication can proceed, which then does
* if multiple can proceed, select chooses pseudo-randomly
* the default case, if present, executes immediately if no channel is ready

Example

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        for {
            ch1<- "from 1"
            time.Sleep(time.Second * 2)
        }
    }()

    go func() {
        for {
            ch2<- "from 2"
            time.Sleep(time.Second * 3)
        }
    }()

    go func() {
        for {
            select {
            case msg1 := <-ch1:
                fmt.Println(msg1)
            case msg2 := <-ch2:
                fmt.Println(msg2)
            default:
                fmt.Println("nothing ready")
                time.Sleep(time.Second * 1)
            }
        }
    }()

    var input string
    fmt.Scanln(&input)
}
```

More

* https://github.com/jreisinger/go-concurrency-patterns
* https://github.com/jreisinger/work
* https://github.com/jreisinger/katas
* The Go Programming Language (2015)
* Go In Practice (2016)
