package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var url string

func main() {

	flag.StringVar(&url, "url", "", "Please provide the url! (required)")

	flag.Parse()
	if url == "" {
		fmt.Println("The url must be provided")
		os.Exit(1)
	}
	// url := "http://cloud.google.com/go/pubsub"
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
}
