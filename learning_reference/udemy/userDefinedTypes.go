/*
i created a new type foo
and we see some of the inherent issues with type conversion

conversion:int(myAge) converts myAge which is type foo to an int type

This is for example
it bad practice to alias types
exception, if you need to attach methods to a type
an example is the time package
	godoc.org/time
	type Duration int64

numeric types default value is 0

*/

package main

import "fmt"

type foo int

func main() {

	var myAge foo
	myAge = 41
	fmt.Printf("%T %v \n", myAge, myAge)
	//idiomatically if you user var, you must be declaring a type and setting it to an empty value
	//otherwise use the shorthand := notation
	anIntType := 1
	//this should work because I havent explicitly
	//made 1 be a type of int
	fmt.Println(myAge + 1)
	//wheras this will not work because I have typed anIntType as an int
	// invalid operation: myAge + anIntType (mismatched types foo and int)
	//fmt.Println(myAge + anIntType)
	//but I can do this
	fmt.Println(int(myAge) + anIntType)
}
