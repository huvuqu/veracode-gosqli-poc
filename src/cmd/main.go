package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	// defer db.Close()
	r := gin.Default()
	r.GET("/product", func(c *gin.Context) {
		var product Product
		code := c.Query("code")
		query := fmt.Sprintf(`SELECT p.code, p.price from product as p where p.code = '%v'`, code)
		db.Raw(query).Scan(&product)
		c.JSON(200, gin.H{
			"price": product.Price,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
