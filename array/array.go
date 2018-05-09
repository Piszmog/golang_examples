package main

import "fmt"

func main() {
    var array [2] string
    array[0] = "hello"
    array[1] = "Piszmog"
    fmt.Println(array)

    var array2 = []string{"hello","Piszmog"}
    fmt.Println(array2)
}
