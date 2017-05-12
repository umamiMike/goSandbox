/**
aken from https://youtu.be/eVlxuST7dCA
Yamil Asusta

function with return the publicKey
when you sign you use private
verify use public key
test by
**/
package main

import (
	//	"crypto/rsa"
	"fmt"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var (
	PORT        = ":3001"
	signKey     *rsa.PrivateKey
	verifyKey   *rsa.PublicKey
	privKeyPath string = "demo.rsa"
	pubKeyPath  string = "demo.rsa.pub"
	signBytes   []byte
	verifyBytes []byte
)

func init() {
	verifyBytes, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		fmt.Println("verifyBytes failure reading public key: ", err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		fmt.Println(err)
	}
}
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := .ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
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
func APIHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You are Authorized")
}
func main() {
	router := mux.NewRouter()
	n := negroni.Classic()
	router.Handle("/api", negroni.New(negroni.HandlerFunc(AuthMiddleware), negroni.HandlerFunc(APIHandler)))
	n.UseHandler(router)
	http.ListenAndServe(PORT, n)
}
