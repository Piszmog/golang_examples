package main

import "fmt"

func main() {
    data := make(chan int)
    go sendData(data)
    fmt.Println(<-data)
}

// Can only receive data
func sendData(data chan<- int) {
    data <- 10
}
