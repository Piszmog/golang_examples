package main

import "fmt"

func main() {
    var a = 10
    fmt.Printf("Address of a is %x\n", &a)
    // Pointer types
    var b = 20
    var c *int
    c = &b
    fmt.Printf("Address of b is %x\n", &b)
    fmt.Printf("Value of b is %d\n", b)
    fmt.Printf("Address of c is %x\n", c)
    fmt.Printf("Value of c is %d\n", *c)
    //Nil pointer
    var nilPointer *int
    fmt.Printf("Address of nilPointer is %x\n", nilPointer)
    if nilPointer == nil {
        fmt.Println("nilPointer is nil")
    }
}
