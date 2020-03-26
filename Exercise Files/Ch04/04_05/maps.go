package main

import (
	"fmt"
        "sort"
)

func main() {
    states := make(map[string]string)
    fmt.Println(states)

    states["WA"] = "Washington"
    states["OR"] = "Oregon"
    states["CA"] = "California"
    fmt.Println(states)

    california := states["CA"]
    fmt.Println(california)

    delete(states, "OR")
    fmt.Println(states)

    states["NY"] = "New York"

    for k, v := range states {
        fmt.Printf("%v: %v\n", k, v)
    }

    // Because Maps are unsorted, we need to do that magic
    // ourselves by using a slice.
    keys := make([]string, len(states))
    i := 0
    for k := range states {
        keys[i] = k
        i++
    }

    fmt.Println(keys)

    // Then we sort the keys and print the states based on the sorted slice
    sort.Strings(keys)
    fmt.Println(keys)

    // Print the sorted map
    fmt.Println("\nSorted")
    for i := range keys {
        fmt.Println(states[keys[i]])
    }


}
