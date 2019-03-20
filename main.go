package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com") // http get request by using http.Get function
	if err != nil {                            //if there is error
		fmt.Println("Error:", err) //print error statement
		os.Exit(1)                 //terminate the program
	}
	//fmt.Println(resp) //print get response

	bs := make([]byte, 99999) // make empty byte slice with empty element of 99999 to begin with by using make function
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}

// interface like reader is interface for many other output with different types and change them into something common
// that will eliminate making or differnt function for all different source that doing same process
