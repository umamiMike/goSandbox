package main

import "fmt"

type Mutatable struct {
	a int
	b int
}

func (m Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

//here we have a pointer to the Mutatable struct and so can 
func (m *Mutatable) MutatePointer(x int, y int) {
	m.a += x
	m.b += y
}

func main() {
	fmt.Println("instantiating a Mutatable struct with initial values 0 and 0")
	m := Mutatable{0, 0}
	fmt.Printf("%v variable m \n running the function StayTheSame, with a non pointer receiver", m)
	m.Mutate()
	fmt.Println(m)
	m.MutatePointer(2,3)
	fmt.Println(m)
	fmt.Println("altering the values inside the method using a pointer reciever updates the state")
	m.MutatePointer(2,3)
	fmt.Println(m)
}
