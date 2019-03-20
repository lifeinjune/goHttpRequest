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
	fmt.Println(resp) //print get response
}
