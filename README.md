# Go Examples
This is for learning Go and having a location to go back and understand Go.

## What is Go
It was initially developed by Google in 2007. It provides a built-in garbage
collection and supports concurrent programming.

It is directly compiled to machine language and not ran in a virtual machine (like Java).

## When is it useful to use?
Go was built with concurrency in mind. Go has 'goroutines' instead of threads
(similar to Kotlin routines). And Goroutines will only use memory when needed.

Goroutines have built-in primitives (channels) to safely communicate between each other.

A single Goroutine can run in multiple threads.

## Differences with other languages (Java)
* There are no classes
  * Everything is in packages and Go uses structs
* Does not support inheritance
* No exceptions (there is error handling tho)

## References
https://gobyexample.com/
https://www.tutorialspoint.com/go/index.htm
https://tour.golang.org/welcome/1
https://golangbot.com/golang-tutorial-part-1-introduction-and-installation/