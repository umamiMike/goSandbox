package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	rand.Seed(time.Now().UnixNano())
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
	demoProduct(db)
}
func demoProduct(db *gorm.DB) {

	// Create
	db.Create(&Product{Code: "L1212", Cost: 1000})

	// Read
	var productWithCode Product
	db.First(&productWithCode, 1)                   // find product with id 1
	db.First(&productWithCode, "code = ?", "L1212") // find product with code l1212
	db.First(&productWithCode).Update("Cost", rand.Intn(1000))
	ck := func(p *Product) []byte {
		v, _ := json.MarshalIndent(p, " ", " ")
		return v
	}(&productWithCode)
	fmt.Println(string(ck))
	//db.Save(&product)
	//fmt.Println(product.Get(db))
	// Delete - delete product
	//	db.Delete(&product)

	db.Create(&Product{Code: "L1212", Cost: 1000})

}
