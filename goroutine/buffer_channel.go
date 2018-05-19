package main

import (
    "fmt"
    "time"
)

func main() {
    names := make(chan string, 2)
    names <- "John"
    names <- "Bob"
    fmt.Println(<-names)
    fmt.Println(<-names)
    fmt.Println("Another")
    data := make(chan int, 2)
    go write(data)
    time.Sleep(2 * time.Second)
    for value := range data {
        fmt.Println("Data", value)
        time.Sleep(2 * time.Second)
    }
}

func write(data chan int) {
    for i := 0; i < 4; i++ {
        data <- i
    }
    close(data)
}
