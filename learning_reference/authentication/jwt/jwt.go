/*
Taken from https://youtu.be/eVlxuST7dCA
Yamil Asusta

function with return the publicKey
when you sign you use private
verify use public key

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
	publicKey, _ = ioutil.ReadFile("demo.rsa.pub")
	privateKey, _ = ioutil.ReadFile("demo.rsa")
}
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err == nil && token.Valid {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Not Authorized")
	}
}
func APIHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You are Authorized")
}
func main() {
	router := mux.NewRouter()
	n := negroni.Classic()

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		token := jwt.New(jwt.GetSigningMethod("RS256"))
		tokenString, _ := token.SignedString(privateKey)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
	})
	router.Handle("/api", negroni.New(negroni.HandlerFunc(AuthMiddleware), negroni.HandlerFunc(APIHandler)))

	n.UseHandler(router)
	http.ListenAndServe(":3000", n)
}
