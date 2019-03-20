package main

import (
	"fmt"
	"io"
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

	// bs := make([]byte, 99999) // make empty byte slice with empty element of 99999 to begin with by using make function
	// //becaseu reader function cannot resize the byte slice, slice initialize with many empty element to accomodate upcoming data stored inside without run out of space
	// resp.Body.Read(bs)
	// /*
	// * Reader function accept slice of byte as argument
	// * Body element in struct is type of readerCloser which mean as long as it meet the requirment of the interface anything can be in the Body as value
	// * original data this case "Body" is element of the resp struct and type of the ReaderCloser interface so can access to Reader
	// * with Body as source of Reader function, the byte slice we passed as argument will contain the Body part of the request because reader function have access to the original data because
	// slice is reference type
	// */
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body) //simpler form of print out body of http request
	/*
		io.Copy accept type Writer and Reader (writer as destination and reader as source) and copy the data from source to the destination
		Writer is exactly oposit of the Reader it read the data from the program and interface (interpret) to many different source of output
	*/
}

// interface like reader is interface for many other output with different types and change them into something common
// that will eliminate making or differnt function for all different source that doing same process
