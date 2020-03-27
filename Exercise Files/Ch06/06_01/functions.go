package main

import "fmt"

func main() {
    doSomething()

    sum := addValues(5,10)
    fmt.Println(sum)

    sum = addAllValues(12,55,10)
    fmt.Println(sum)

}

func doSomething() {
    fmt.Println("Doing something!")

}

func addValues(value1, value2 int) int {
    return value1 + value2

}

// lowercase starting functions are private to this package
func addAllValues(values ...int) int {
    sum := 0
    for i := range values {
        sum += values[i]
    }

    fmt.Printf("%T\n",values)
    return sum

}
