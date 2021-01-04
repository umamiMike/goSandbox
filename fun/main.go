package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var f = fmt.Println
var si = strconv.Atoi
var is = strconv.Itoa

type rect struct {
	width, height int
}

func main() {
	tmer := time.NewTimer(time.Second)
	go func() {
		<-tmer.C
		f("timer expired")
	}()
	stop := tmer.Stop()
	if stop {

		f("timer stop")
	}

	ioutil.ReadFile("file")

	n_int := 1234
	f("the input int is:", n_int)
	f("the sum is: ", sum(n_int))
	f("the product is: ", product(n_int))
}
func convertDigitsToIntSlice(input int) []int {

	n_str := strconv.Itoa(input)                   // convert to a string
	string_slice := strings.Split(n_str, "")       // split it up into a string slice
	int_sl := convStrSliceToIntSlice(string_slice) //convert that slice to an int slice
	return int_sl
}

func convStrSliceToIntSlice(s []string) []int {
	var intSlice []int
	for k := range s {
		cv, _ := si(s[k])
		intSlice = append(intSlice, cv)
	}
	return intSlice
}

func product(input int) int {
	int_sl := convertDigitsToIntSlice(input)

	prod := 1
	for i := range int_sl {
		prod = prod * int_sl[i]
	}
	return prod
}

func sum(input int) int {

	int_sl := convertDigitsToIntSlice(input)

	sum := 0
	for i := range int_sl {
		sum += int_sl[i]
	}
	return sum
}
