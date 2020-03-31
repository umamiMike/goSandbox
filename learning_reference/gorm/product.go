package main

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string
	Code string
	Cost uint
}

func (p *Product) Get(db *gorm.DB) (interface{}, bool) {
	return db.Where(&p).Get("Price")
}
