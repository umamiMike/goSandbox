/**
https://golang.org/doc/faq#methods_on_values_or_pointers
EX:
func (s *MyStruct) pointerMethod() { } // method on pointer
func (s MyStruct)  valueMethod()   { } // method on value
*/

package main

import "fmt"

type Mutatable struct {
	a int
	b int
}

func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate() {
	m.a = 5
	m.b = 7
}

func main() {
	m := Mutatable{0, 0}
	fmt.Println(m)
	m.StayTheSame()
	fmt.Println(m)
	m.Mutate()
	fmt.Println(m)
}
