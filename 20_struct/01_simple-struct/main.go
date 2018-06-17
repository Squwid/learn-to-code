package main

import (
	"fmt"
)

type foo int // type foo with underlying type of int

// type person with following fields
type person struct {
	first string
	last  string
	age   int
}

func main() {
	var myAge foo
	myAge = 50
	fmt.Printf("%T %v\n", myAge, myAge) // main.foo 50

	var yourAge int
	yourAge = 20
	fmt.Printf("%T %v\n", yourAge, yourAge)

	// This DOESNT work
	// fmt.Println(myAge + yourAge)

	// This DOES work
	fmt.Println(int(myAge) + yourAge)

	p1 := person{"Ben", "Whitelaw", 20} // struct literal
	p2 := person{"Tom", "Currie", 20}
	fmt.Println(p1, p2)
}
