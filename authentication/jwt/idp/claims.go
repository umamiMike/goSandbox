package main

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type myClaimOne struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
