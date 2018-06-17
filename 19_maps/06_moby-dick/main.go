package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// get the book of moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		panic(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// create slice to hold counts
	buckets := make([]int, 200)

	// loop over words
	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}
	// How many words are in each bucket
	fmt.Println(buckets[65:122])
}

func hashBucket(word string) int {
	return int(word[0])
}

// hasBucketEven takes a word and a bucket and spreads it out just about evenly
func hasBucketEven(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
