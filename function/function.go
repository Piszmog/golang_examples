package main

import "fmt"

func main() {
    var a, b = 5, 10
    max := max(a, b)
    fmt.Printf("Max %d\n", max)
    //Multi return values
    c, d := swap("Mahesh", "Kumar")
    fmt.Println(c, d)
}

func max(num1, num2 int) int {
    var result int
    if num1 > num2 {
        result = num1
    } else {
        result = num2
    }
    return result
}

func swap(x, y string) (string, string) {
    return y, x
}
