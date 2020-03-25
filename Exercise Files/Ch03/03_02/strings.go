package main

import (
	"fmt"
	"strings"
)

func main() {
    str1 := "An implicit string"
    fmt.Printf("str1: %v:%T\n", str1, str1)

    var str2 string = "An explicitstring"
    fmt.Printf("str1: %v:%T\n", str2, str2)

    fmt.Println(strings.ToUpper(str1))
    fmt.Println(strings.Title(str1))

    lValue := "hello"
    uValue := "HELLO"

    fmt.Println("equal?",(lValue == uValue))
    // equalfold makes strings uppercase values
    fmt.Println("equal case sens?",strings.EqualFold(lValue, uValue))

    fmt.Println("Contains exp?",strings.Contains(str1,"impl"))
    fmt.Println("Contains exp?",strings.Contains(str2,"impl"))
}
