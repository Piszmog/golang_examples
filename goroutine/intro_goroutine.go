package main

import (
    "fmt"
    "time"
)

func main() {
    go hello()
    time.Sleep(1 * time.Second)
    go numbers()
    go letters()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("\nEnding")
}

func hello() {
    fmt.Println("Hello")
}

func numbers() {
    for i := 0; i < 10; i++{
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d", i)
    }
}

func letters() {
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c", i)
    }
}
