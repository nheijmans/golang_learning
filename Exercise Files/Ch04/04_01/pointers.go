package main

import "fmt"

//A pointer references a location in memory, and obtaining the value stored at that location 
//is known as dereferencing the pointer. As an analogy, a page number in a book's index could 
//be considered a pointer to the corresponding page; dereferencing such a pointer would be 
//done by flipping to the page with the given page number and reading the text found on that page.

//Using pointers significantly improves performance for repetitive operations like traversing 
//iterable data structures, e.g. strings, lookup tables, control tables and tree structures. 
//In particular, it is often much cheaper in time and space to copy and dereference pointers 
//than it is to copy and access the data to which the pointers point.

//source: https://en.wikipedia.org/wiki/Pointer_(computer_programming)

func main() {
    //Pointers are variables to address space
    // With the asterix, a pointer is created
    var p *int

    if p!= nil {
        fmt.Println("Value of p:",p)
    } else {
        fmt.Println("p is nil")
    }

    // & symbol means "connect the pointer to this variable"
    var v int = 42
    p = &v

    // to print the value of the pointer address, the variable starts with asterix (*)
    // if you print the variable "p" only, it shows the address reserved
    if p!= nil {
        fmt.Println("Value of p:",*p)
    } else {
        fmt.Println("p is nil")
    }
}
