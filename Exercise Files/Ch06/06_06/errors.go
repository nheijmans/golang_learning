package main

import (
	"fmt"
        "os"
        "errors"
)

func main() {
	f, err := os.Open("filename.ext")

        if err == nil {
            fmt.Println(f)
        } else {
            fmt.Println(err)
        }

        myError := errors.New("My error string")
        fmt.Println(myError)

	attendance := map[string]bool{
	    "Ann": true,
	    "Mike": true}

            // the ok verifies if the M key is present, otherwise jump to else will be triggered
        attended, ok := attendance["M"]
        if ok {
            fmt.Println("Mike attended?", attended)
        } else {
            fmt.Println("No info for Mike")
        }
}
