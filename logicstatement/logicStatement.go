package main

import (
    "fmt"
    "time"
)

func main() {
    // If statement
    var success = false
    if success {
        fmt.Println("Success")
    } else {
        fmt.Println("Fail")
    }
    // Switch statement
    value := 1
    switch value {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    default:
        fmt.Println("default")
    }
    // Select statement
    output1 := make(chan string)
    output2 := make(chan string)
    // Call the functions -- they will sleep
    go server1(output1) // can call the function in like -- go func() {..}
    go server2(output2)
    // Block until one of the calls completes
    select {
    case s1 := <-output1:
        fmt.Println(s1)
    case s2 := <-output2:
        fmt.Println(s2)
    }
}

// Functions for select statement
func server1(ch chan string) {
    time.Sleep(6 * time.Second)
    ch <- "from server 1"
}
func server2(ch chan string) {
    time.Sleep(3 * time.Second)
    ch <- "from server 2"
}
