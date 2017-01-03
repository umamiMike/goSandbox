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

func (p Person) Greeting() {
	fmt.Println("I am only a normal bloke or blokette")
}

//illustrating method overriding by changin receivers...you can use the same method name but since
//it is a different receiver it maps correctly
func (d DoubleZero) Greeting() {
	fmt.Println("WHy I am the doubleZed")
}

//methods
/*
func (reciever {the type of who can use me})
it is not idiomatic to use this or self

*/
func (p Person) fullName() string {
	return p.First + " " + p.Last
}

const (
	lb = "-------------------------"
)

func main() {
	//adding some person objects
	p1 := Person{"james", "bond", 24}
	p2 := Person{"miss", "moneypenny", 18}

	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Println(p2.First, p2.Last, p2.Age)
	dbloh1 := DoubleZero{
		Person:        p1,
		First:         "Double Aught Seven",
		LicenseToKill: true,
	}

	fmt.Println("manual assembly of struct properties", lb)
	fmt.Println(dbloh1.First, dbloh1.Person.Last, "...", dbloh1.Person.First, dbloh1.Person.Last)
	//this illustrates the use of the fullName Method use on one of the objects
	fmt.Println("Using the method FullName to print the full name of p1", lb)
	fmt.Println(p1.fullName())

	//example of an embedded type
	//person type is IN DoubleZero struct
	p3 := DoubleZero{
		Person: Person{
			First: "The",
			Last:  "Dude",
			Age:   50,
		},
		LicenseToKill: true,
	}
	//illustrates basic embedded type
	fmt.Println(p3.Person.Last)

	//illustrates method override by showing different greetings based on which object is the receiver of the call\
	p1.Greeting()        //should print from the person object
	p3.Greeting()        // whereas this greeting is for the DoubleZero struct
	p3.Person.Greeting() //and this shows how to get to the Person greeting INSIDE the embedded type
}
