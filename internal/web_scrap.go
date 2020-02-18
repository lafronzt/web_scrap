package internal

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetURL parse the response of the Website
func GetURL(url string) {

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
