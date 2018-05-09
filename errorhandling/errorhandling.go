package main

import (
    "errors"
    "math"
    "fmt"
)

func main() {
    // function error
    result, err := Sqrt(-1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
    result, err = Sqrt(9)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
    // struct error
    for _, i := range []int{7, 42} { //_ is a 'blank identifier' -- value is discarded
        if r, e := structFunc(i); e != nil {
            fmt.Println("structFunc failed:", e)
        } else {
            fmt.Println("structFunc worked:", r)
        }
    }
    _, e := structFunc(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}

// Function error
func Sqrt(value float64) (float64, error) {
    if value < 0 {
        return 0, errors.New("math: negative number passed to sqrt")
    }
    return math.Sqrt(value), nil
}

// Struct Error
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func structFunc(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}
