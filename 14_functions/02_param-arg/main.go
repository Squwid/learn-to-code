package main

import "fmt"

func main() {
	// arguments
	_ = combine("a", "b")

}

// str1 and str2 are parameters
func combine(str1, str2 string) string {
	return fmt.Sprint(str1, str2)
}
