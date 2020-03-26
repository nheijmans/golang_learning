package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
//	dow := rand.Intn(6) + 1
//	fmt.Println("Day", dow)

        result := ""
        switch dow := rand.Intn(6) +1; dow {
            case 1:
                result = "It's Sunday!"
            case 7:
                result = "Saturday"
            default:
                result = "It's a weekday!"
        }
        fmt.Println(result)

        x := -42
        switch {
            case x < 0:
                result = "Less than zero"
                //fallthrough // Can be used to let other cases also be matched
            case x == 0:
                result = "Equals to zero"
            default:
                result = "Greater than zero"
         }
         fmt.Println(result)
}
