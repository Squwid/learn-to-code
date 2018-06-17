package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p *person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"Ben", "Whitelaw", 20}
	fmt.Println(p1.fullName())

	p2 := &person{"Tom", "Currie", 20}
	fmt.Println(p2)       // &{Tom Currie 20}
	fmt.Println(*p2)      // {Tom Currie 20}
	fmt.Println(p2.first) // Tom
}
