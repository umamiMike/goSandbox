package main

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string
	Code string
	Cost uint
}
