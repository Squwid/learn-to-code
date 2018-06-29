package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf) // set output changes to where the output gets saved to
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}

}
