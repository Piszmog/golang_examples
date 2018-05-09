package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    for i := range numbers {
        fmt.Println(i)
    }
    // create a map
    capitalMap := map[string]string{"France": "Paris", "Colorado": "Denver"}
    for capital := range capitalMap {
        fmt.Println(capitalMap[capital])
    }
    for location, capital := range capitalMap {
        fmt.Println("Capital of", location, "is", capital)
    }
    var anotherCapitalMap map[string]string
    anotherCapitalMap = make(map[string]string)
    anotherCapitalMap["France"] = "Paris"
    anotherCapitalMap["Colorado"] = "Denver"
    for location, capital := range anotherCapitalMap {
        fmt.Println("Capital of", location, "is", capital)
    }
    capital, exists := anotherCapitalMap["United States"]
    if exists {
        fmt.Println("United states capital is", capital)
    } else {
        fmt.Println("United states capital is not in map")
    }
    // delete from map
    delete(anotherCapitalMap, "France")
    franceCapital, franceCapitalExists := anotherCapitalMap["France"]
    if franceCapitalExists {
        fmt.Println("France capital is", franceCapital)
    } else {
        fmt.Println("France capital is not in map")
    }
}
