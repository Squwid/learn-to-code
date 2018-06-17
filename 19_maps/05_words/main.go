package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

// example of getting a website that has a list of every word in the english language
// word hash table hashtable
func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		panic(err)
	}

	words := make(map[string]string)

	/*
		Other way to do pretty much the same thing but without a buffer:

		bs, _ := ioutil.ReadAll(res.Body)
		str := string(bs)
		fmt.Println(str)

	*/

	// bufio is an input output buffer
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords) // split by words

	// scans while we have words
	for sc.Scan() {
		words[sc.Text()] = "" // each token gets added to the map
	}

	// check for errors
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	// print the words
	for k := range words {
		fmt.Println(k)
	}
}
