package main

import (
	"fmt"
)

// Creating slices different ways
// slices maps and channels are reference types

func main() {
	var s []string                // only way of creating that makes the value of s nil
	student := make([]string, 35) // length == 35; capacity == 35  Correct way to do it
	ss := []string{}
	sss := new([15]string)[:10] // length == 10; capacity == 15

	fmt.Println(student == nil) // false
	fmt.Println(s == nil)       // true
	fmt.Println(ss == nil)      // false

	fmt.Printf("\n----------\n")

	s = append(s, "test1")
	fmt.Println(s == nil) // false
	fmt.Printf("\n----------\n")

	fmt.Printf("Student:: Length %v\tCapactiy %v\n", len(student), cap(student))
	fmt.Printf("sss::: Length %v\tCapactiy %v\n", len(sss), cap(sss))
	fmt.Printf("ss::: Length %v\tCapactiy %v\n", len(ss), cap(ss))
}
