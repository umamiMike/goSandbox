package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

type Product struct {
	gorm.Model
	Code  string `json:"code"`
	Price int    `json:"price"`
}
