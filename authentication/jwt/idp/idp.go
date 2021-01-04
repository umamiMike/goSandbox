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
	"crypto/rsa"
	"fmt"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	//	"time"
)

const (
	PORT = ":3000"
)

var (
	signKey     *rsa.PrivateKey
	verifyKey   *rsa.PublicKey
	privKeyPath string = "demo.rsa"
	pubKeyPath  string = "demo.rsa.pub"
	signBytes   []byte
	verifyBytes []byte
)

func init() {
	var err error

	signBytes, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		fmt.Println("priv key fail: ", err)
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		fmt.Println(err)
	}
	verifyBytes, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		fmt.Println("virifyBytes failure reading public key: ", err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	router := mux.NewRouter()
	n := negroni.Classic()
	router.Handle("/login", negroni.New(negroni.HandlerFunc(LoginHandler)))
	n.UseHandler(router)
	http.ListenAndServe(PORT, n)
}
func LoginHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Fprint(w, "in the login route:\n")
	token := jwt.New(jwt.SigningMethodRS256)
	fmt.Fprint(w, "the token is ")
	fmt.Fprint(w, token)
	/*
		claims := make(jwt.MapClaims)
		claims["username"] = "theusername"
		claims["iat"] = time.Now()
	*/
	//	token.Claims = claims
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, "This is your token string, use it wisely\n")
	fmt.Fprint(w, tokenString)
}
