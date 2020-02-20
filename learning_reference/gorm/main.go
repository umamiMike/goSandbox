package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Name string
	Code string
	Cost uint
}

type Service struct {
	gorm.Model
	Name string
	Code string
	Cost uint
}

func main() {

	rand.Seed(time.Now().UnixNano())

	//creates db if doesnt exist
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	if !db.HasTable(&Product{}) {
		db.CreateTable(&Product{})
	}
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Cost: 1000})

	// Read
	var productWithCode Product
	db.First(&productWithCode, 1) // find product with id 1
	fmt.Println(productWithCode.Code)
	db.First(&productWithCode, "code = ?", "L1212") // find product with code l1212
	fmt.Println(productWithCode.Cost)

	//// Update - update product's price to 2000
	//pr := Product{}
	db.First(&productWithCode).Update("Cost", rand.Intn(1000))
	fuck := func(p *Product) []byte {
		v, _ := json.MarshalIndent(p, " ", " ")
		return v
	}(&productWithCode)

	fmt.Println(string(fuck))
	//db.Save(&product)

	//fmt.Println(product.Get(db))

	db.Create(&Product{Code: "L1212", Cost: 1000})
	// Delete - delete product
	//	db.Delete(&product)
}

func (p *Product) Get(db *gorm.DB) (interface{}, bool) {
	return db.Where(&p).Get("Price")

}
