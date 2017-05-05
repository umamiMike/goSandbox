/**
Taken from https://youtu.be/eVlxuST7dCA
Yamil Asusta

function with return the publicKey
when you sign you use private
verify use public key
test by
**/
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
		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Not Authorized")
	}
}
func main() {
	router := mux.NewRouter()
	n := negroni.Classic()
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "in the login route:\n")
		token := jwt.New(jwt.GetSigningMethod("RS256"))
		tokenString, _ := token.SignedString(privateKey)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
	})
	n.UseHandler(router)
	http.ListenAndServe(":3000", n)
}
