package main

import "fmt"

//creates new struct
//which has 3 fields
//with their own type (string or int in this case)
// structs can have anonymous fields
// make sure you Capitalize the first names of the fields
type Person struct {
	First string
	Last  string
	Age   int
}
type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := Person{"james", "bond", 24}
	p2 := Person{"miss", "moneypenny", 18}

	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Println(p2.First, p2.Last, p2.Age)

	dbloh1 := DoubleZero{
		Person:        p1,
		First:         "Double Aught Seven",
		LicenseToKill: true,
	}

	fmt.Println(dbloh1.First, dbloh1.Person.Last, "...", dbloh1.Person.First, dbloh1.Person.Last)

}
