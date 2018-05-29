package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	a := "stored in a"
	b := "stored in b"
	fmt.Println("a- ", a)
	// _ is a Blank Identifier
	_ = b
	getHTTPPage()
}

func getHTTPPage() {
	res, _ := http.Get("http://github.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
