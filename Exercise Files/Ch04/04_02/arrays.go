package main

import (
	"fmt"
)

func main() {
    var colors [3]string
    colors[0] = "Red"
    colors[1] = "Blue"
    colors[2] = "Yellow"

    fmt.Println(colors)
    fmt.Println(colors[1])

    var numbers = [5]int{5,3,1,2,4}
    fmt.Println(numbers)

    fmt.Println("no of colors:",len(colors))
    fmt.Println("no of numbers:",len(numbers))
}
