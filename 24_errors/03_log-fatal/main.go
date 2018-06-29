package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err) // also makes a call to os.Exit(1)
	}
}
