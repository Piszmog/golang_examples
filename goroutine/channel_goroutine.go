package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan bool)
    go channelHello(done)
    <-done
    startingNumber := 20
    doubleResult := make(chan int)
    halfResult := make(chan int)
    go calcDouble(startingNumber, doubleResult)
    go calcHalf(startingNumber, halfResult)
    doubles, halfs := <-doubleResult, <-halfResult
    fmt.Println("Final Number", doubles+halfs)
    fmt.Println("Ending")
}

func channelHello(done chan bool) {
    time.Sleep(2 * time.Second)
    fmt.Println("Hello")
    done <- true
}

func calcDouble(number int, doubleResult chan int) {
    doubleResult <- number * 2
}

func calcHalf(number int, halfResult chan int) {
    halfResult <- number % 2
}
