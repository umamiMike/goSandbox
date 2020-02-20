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

//new user type Shape of type Interfaces
//anything implements this interface
//anything that implements area() float64
//
type Shape interface {
	area() float64
}

//create a user defined type struct
type Square struct {
	side float64
}

//attach a method to it called area()
func (z Square) area() float64 {
	return z.side * z.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

func main() {
	s := Square{10}
	fmt.Println("square area: ", s.area()) //default straight way of running the method

	c := Circle{10}
	fmt.Println("circle area: ", c.area())
}
