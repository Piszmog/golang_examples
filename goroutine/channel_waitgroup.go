package main

import (
    "sync"
    "time"
    "fmt"
)

func main() {
    numbers := 3
    var wg sync.WaitGroup
    for i := 0; i < numbers; i++ {
        wg.Add(1)
        go processData(i, &wg)
    }
    wg.Wait()
    fmt.Println("Done")
}

func processData(i int, wg *sync.WaitGroup) {
    fmt.Println("Starting", i)
    time.Sleep(2 * time.Second)
    fmt.Println("Ending", i)
    wg.Done()
}
