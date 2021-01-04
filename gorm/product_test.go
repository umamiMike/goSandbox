package main

import (
	"testing"

	"github.com/jinzhu/gorm"
)


func NewStore(dbName string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	return db

}

func TestGet(t *testing.T) {
	// Create
	db.Create(&Product{Code: "L1212", Cost: 1000})
	db

	//// Read
	//var productwithcode product
	//db.first(&productwithcode, 1)                   // find product with id 1
	//db.first(&productwithcode, "code = ?", "l1212") // find product with code l1212
	//db.First(&productWithCode).Update("Cost", rand.Intn(1000))
	//ck := func(p *Product) []byte {
	//v, _ := json.MarshalIndent(p, " ", " ")
	//return v
	//}(&productWithCode)
	//fmt.Println(string(ck))
	//db.Save(&product)
	//fmt.Println(product.Get(db))
	//Delete - delete product
	//	db.Delete(&product)

	//db.Create(&Product{Code: "L1212", Cost: 1000})

}
