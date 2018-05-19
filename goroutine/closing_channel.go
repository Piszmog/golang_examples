package main

import "fmt"

func main() {
    data := make(chan int)
    go emit(data)
    for {
        value, ok := <-data
        if ok == false {
            break
        }
        fmt.Println("Data", value, ok)
    }
    fmt.Println("Another")
    // another
    anotherData := make(chan int)
    go emit(anotherData)
    for value := range anotherData {
        fmt.Println("Data", value)
    }
}

func emit(data chan int) {
    for i := 0; i < 10; i++ {
        data <- i
    }
    close(data)
}
