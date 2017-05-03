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
	publicKey, err := ioutil.Readfile("public.pem")
	privateKey, err := ioutil.Readfile("private.pem")
}
func AuthMidddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token() (interface{}, error){
		return publicKey, nil
	})
	if err == nil && token.Valid(){
		next(r,w)
	}else{w.WriteHeader}
}
func main() {

}
