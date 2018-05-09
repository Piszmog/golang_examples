package main

import "fmt"

func main() {
    // slice ~= Java's List
    var numbers []int //unspecified size
    printSlice(numbers)
    var numbersMake = make([]int, 5, 5) //slice of length 5, capacity 5
    printSlice(numbersMake)
    // Nil slice
    if numbers == nil {
        fmt.Println("numbers is nil")
    }
    //Sub slice
    newNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    printSlice(newNumbers)
    fmt.Println("numbers[1:4]", newNumbers[1:4])
    fmt.Println("numbers[:3]", newNumbers[:3])
    fmt.Println("numbers[4:]", newNumbers[4:])
    //append
    numbers = append(numbers, 0)
    printSlice(numbers)
    numbers = append(numbers, 1)
    printSlice(numbers)
}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
