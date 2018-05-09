package main

import "fmt"

func main() {
    var john, bob person
    john = person{"John", 20}
    bob.name = "Bob"
    bob.age = 56
    fmt.Println(john)
    fmt.Println(bob)

    sean := person{"Sean", 87}
    seanPointer := &sean
    fmt.Println(seanPointer.age)

    seanPointer.age = 88
    fmt.Println(seanPointer.age)
}

type person struct {
    name string
    age  int
}
