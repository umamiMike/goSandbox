package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom"`
}

func main() {
	p1 := Person{"Mike", "McEvil", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)         //print the byteslice in ascii format
	fmt.Printf("%T \n", bs) //print the type []uint8
	fmt.Println(string(bs)) //convert it to a string

	var p2 Person
	newbs := []byte(`{"First":"James", "Last": "Bond", "wisdom":20}`)
	json.Unmarshal(newbs, &p2) //unmarshal newbs into p2 object
 
	var p3 Person.UnmarshalJSON([]byte(`{"First":"Mycroft", "Last": "Holmes","wisdom":40}`)
	fmt.Println("-------")
	fmt.Println(p2.First) //proof by accessing values
	fmt.Println(p2.Last)
	fmt.Println(p2.Age) //even unmarshalling json tag wisdom into Age
	fmt.Printf("%T \n", p2)
}
func (p *Person) UnmarshalJSON(data []byte) error {
json.Unmarshal(data, &p)
return nil
}
