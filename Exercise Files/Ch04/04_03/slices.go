package main

import (
	"fmt"
        "sort"
)

func main() {
    var colors = []string{"Red","Green","Yellow"}
    fmt.Println(colors)

    colors = append(colors, "Purple")
    fmt.Println(colors)

    colors = append(colors[1:len(colors)])
    fmt.Println(colors)
	
    colors = append(colors[:len(colors)-1])
    fmt.Println(colors)

    numbers := make([]int, 5, 5)
    numbers[0] = 18
    numbers[1] = 60
    numbers[2] = 34
    numbers[3] = 77
    numbers[4] = 8

    fmt.Println(numbers)
    numbers = append(numbers, 235)
    fmt.Println(numbers)
    fmt.Println(cap(numbers))

    sort.Ints(numbers)
    fmt.Println(numbers)

}
