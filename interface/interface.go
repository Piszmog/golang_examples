package main

import "fmt"

func main() {
    var fido = dog{}
    fmt.Println("Dog says", fido.roar())
}

type animal interface {
    roar() string
}

type dog struct {
    sound string
}

func (dog dog) roar() string {
    return "bark"
}
