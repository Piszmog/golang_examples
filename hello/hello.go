package main

import (
    "fmt"
    "time"
)

var g int

func main() {
    var hello, name string
    hello = "hello"
    name = "Piszmog"
    // dynamic type declaration
    world := "world"
    //Prints Hello
    fmt.Println(hello + ", " + name)
    fmt.Printf("hello is of type %T\n", hello)
    fmt.Println(hello + " " + world)
    //Prints the time
    fmt.Println("The time is", time.Now())
    // mixed declaration
    var a, b, c = 1, 2, "foo"
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    //Global variable
    g = 10
    fmt.Printf("Global %d", g )
}
