package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type dbStore struct {
	db *gorm.DB
}

func NewDB(dbName string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	return db

}
func main() {

	dbStore := dbStore{db: NewDB("test.db")}
	db := dbStore.db
	defer db.Close()
	// Migrate the schema
	if !db.HasTable(&Product{}) {
		db.CreateTable(&Product{})
	}
	db.AutoMigrate(&Product{})
}
