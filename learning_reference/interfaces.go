/*
Interfaces
Substitutability
Polymorphism


*/

package main

import (
	"fmt"
	"math"
)

//create a user defined type struct
type Square struct {
	side float64
}

//attach a method to it called area()
func (z Square) area() float64 {
	return z.side * z.side
}

//new user type Shape of type Interfaces
//anything implements this interface
//anything that implements area() float64
//
type Shape interface {
	area() float64
}

//since square is a shape (by implementing funcs called for by shape interface)
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {

	return math.Pi * c.radius * c.radius
}
func main() {
	s := Square{10}
	fmt.Println(s.area()) //default straight way of running the method
	info(s)

	c := Circle{60}

	info(c)
}
