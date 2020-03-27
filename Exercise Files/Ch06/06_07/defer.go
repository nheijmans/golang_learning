package main

import (
	"fmt"
)

func main() {
    defer fmt.Println("Close the file")
    fmt.Println("Open the file")

    defer fmt.Println("statement 1")
    defer fmt.Println("statement 2")
    defer fmt.Println("statement 3")
    defer fmt.Println("statement 4")
    fmt.Println("Undefined")

    myFunc()

    x := 1000
    defer fmt.Println("Value of x:",x)
    x++
    fmt.Println("Value of x:",x)
}

func myFunc() {
    defer fmt.Println("Deferred in the function")
    fmt.Println("Not deferred in the function")

}
