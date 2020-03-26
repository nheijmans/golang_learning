package main

import (
	"fmt"
)


type Dog struct {
    Breed string
    Weight int
}

func main() {

    // structs contain data and classes, fields for data and methods

    poodle := Dog{"Poodle",34}
    fmt.Println(poodle)
    fmt.Printf("%+v\n",poodle)
    fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)


}
