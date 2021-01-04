/**
This is a very simple implementation of an http reqeust to stripe.com
I pulled it out of postman.

*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.stripe.com/v1"
	resource := "/customers"

	req, _ := http.NewRequest("GET", url+resource, nil)

	req.Header.Add("authorization", "Basic c2tfdGVzdF9CUW9raWtKT3ZCaUkySGxXZ0g0b2xmUTI6Og==")
	req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("postman-token", "96c17845-1171-e85c-6cfc-a1aed083c82c")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
