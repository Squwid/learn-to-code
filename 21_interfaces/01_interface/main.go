package main

import (
	"fmt"
	"math"
)

// Square Interface Example

// Square is a user defined type, underlying type is 'struct'
type Square struct {
	side float64
}

// Circle is a user defined type with the underlying type of struct
type Circle struct {
	radius float64
}

// Triangle is a user defined type with the underlying type of struct
type Triangle struct {
	side float64
}

// attach a method to the type square
func (z Square) area() float64 {
	return z.side * z.side
}

func (z Circle) area() float64 {
	return 3.14159 * math.Pow(z.radius, 2)
}

func (t *Triangle) area() float64 {
	return t.side * 1 / 2
}

// Shape is a user defined type, underlying type is 'interface'
// anything that has ()area()method signiture implements the Shape interface
// the types need everything inside of here to inherit the type
type Shape interface {
	area() float64
}

// Can pass Square into shape because it is declared a shape earlier
func info(z Shape) {
	fmt.Println(z)
	fmt.Print(z.area(), "\n\n")
}

func main() {
	s := Square{10}
	c := Circle{15}

	info(s) // concrete type is square
	info(c) // concrete type is circle

	c1 := &Circle{7}                // you are able to do this even if c1's type is *main.Circle
	fmt.Printf("c1-type: %T\n", c1) // c1-type: *main.Circle
	info(c1)

	c2 := &Triangle{3}
	fmt.Printf("c2-type: %T\n", c2)
	info(c2)
	/*
		CANT DO THIS - Different types
		c2 := Triangle{3}
		info(c2)
	*/
}
