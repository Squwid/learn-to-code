package main

import (
	"fmt"
)

// Person is
type Person struct {
	First string
	Last  string
	Age   int
}

// DoubleZero is
type DoubleZero struct {
	Person        // anon type because its only the type
	First         string
	LicenceToKill bool
}

/*
	INHERITANCE (Reusability)

	Inner type (Person), gets promoted to the outer type (DoubleZero)
	Outer type now has access to inner type
*/
 
func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenceToKill: true,
	}
	fmt.Println(p1)
	fmt.Println(p1.Person.Age)
}
