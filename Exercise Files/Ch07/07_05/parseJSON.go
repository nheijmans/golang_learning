package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	 "encoding/json"
	 "strings"
	 "math/big"
)

type Tour struct {
    Name, Price string

}

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

//	fmt.Println(content)
        tours := toursFromJson(content)
//        fmt.Println(tours)

        // go over each tour and print the name and price, with price converted to a float
        for _, tour := range tours {
            price, _, _ := big.ParseFloat(tour.Price, 10,2,big.ToZero)
            fmt.Printf("%v ($%.2f)\n", tour.Name, price)
        }

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
    // retrieve content from a website
	resp, err := http.Get(url)
	checkError(err)
	
	defer resp.Body.Close()
        // extract the content of the body
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func toursFromJson(content string) []Tour {
    // make a slice object, with reserved 20 spaces
    tours := make([]Tour, 0, 20)

    // decode the JSON content from the web request
    decoder := json.NewDecoder(strings.NewReader(content))
    _, err := decoder.Token()
    checkError(err)

    // create an object of the Tour structure
    var tour Tour
    for decoder.More() {
        //parse the json content and it only gets the values in the struct
        err := decoder.Decode(&tour)
        checkError(err)
        tours = append(tours, tour)
    }

    return tours
}
