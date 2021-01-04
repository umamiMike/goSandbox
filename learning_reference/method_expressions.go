package main

import (
	"fmt"
	"reflect"
)

//type T struct{}

//func (T) Foo(s string) { println(s) }

//var fn func(T, string) = T.Foo

//func main() {
//v := fn
//fmt.Println(v)

//}

func main() {
	methodExpressionPlay2()
}

//Yes, it's possible. You have 2 (3) options:
//Spec: Method expressions
// you might add a receiver to a method

//The expression Foo.A yields a function equivalent to A but with an explicit receiver as its first argument; it has signature func(f Foo).
type Foo int

func (f Foo) A() {
	fmt.Println("A")
}
func (f Foo) B() {
	fmt.Println("B")
}
func (f Foo) C() {
	fmt.Println("C")
}

var f Foo

func methodExpressionPlay2() {

	//bar is a function
	// which runs the passed in function with an argument of a
	//type.a function
	bar := func(m func(f Foo)) {
		m(f)
	}

	bar(Foo.A)
	bar(Foo.B)
	bar(Foo.C)
}

//Here the method receiver is explicit. You only pass the method name (with the
//type it belongs to) to bar(), and when calling it, you have to pass the
//actual receiver: m(f).

//Output as expected (try it on the Go Playground):

//A
//B
//C

//Spec: Method values

//If f is a value of type Foo, the expression f.A yields a function value of type func() with implicit receiver value f.
func methodExpressionPlay() {

	var f Foo

	bar := func(m func()) {
		m()
	}
	bar(f.A)
	bar(f.B)
	bar(f.C)

}

//Note that here the method receiver is implicit, it is saved with the function
//value passed to bar(), and so it is called without explicitly specifying it:
//m().

//Output is the same (try it on the Go Playground).  (For completeness:
//reflection)

//Inferior to previous solutions (both in performance and in "safeness"), but
//you could pass the name of the method as a string value, and then use the
//reflect package to call the method by that name. It could look like this:
func methoExpressionPlayReflect() {
	var f Foo
	bar := func(name string) {
		reflect.ValueOf(f).MethodByName(name).Call(nil)
	}
	bar("A")
	bar("B")
	bar("C")
}

//Try this on the Go Playground.
