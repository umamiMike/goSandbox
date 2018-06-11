package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/products/")
	{
		v1.POST("/", createProduct)
		v1.GET("/", getAllProducts)
		v1.GET("/:id", getSingleProduct)
		v1.PUT("/:id", updateSingleProduct)
		v1.DELETE("/:id", deleteSingleProduct)
	}
	router.Run()
	db, err := gorm.Open("sqlite3", "terst.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)
	// Delete - delete product
	//	db.Delete(&product)
}
func createProduct(c *gin.Context) {

}
func getAllProducts(c *gin.Context) {

}
func getSingleProduct(c *gin.Context) {

}
func updateSingleProduct(c *gin.Context) {

}
func deleteSingleProduct(c *gin.Context) {

}
