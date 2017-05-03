/*
Taken from https://youtu.be/eVlxuST7dCA
Yamil Asusta
*/
package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var (
	signingkey = "fuelRox"
	privateKey []byte
	publicKey  []byte
)

func init() {
	publicKey, err := ioutil.Readfile("demo.rsa.pub")
	privateKey, err := ioutil.Readfile("demo.rsa")
}
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err == nil && token.Valid() {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Not Authorized")
	}
}
func APIHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprinf(w, "You are Authorized")
}
func main() {

}
