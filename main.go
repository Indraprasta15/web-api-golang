package main

import (
	"fmt"
	"log"
	"web-api-golang/book"
	"web-api-golang/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "admin:@FEBruary1994@tcp(127.0.0.1:3306)/web_api_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db Connection error..")
	}

	db.AutoMigrate(&book.Book{})
	//CRUD

	book := book.Book{}
	book.Title = "Fenomena Jahiliyyah"
	book.Description = "Ini adalah buku yang inspiratif dan cocok untuk semua kalangan."
	book.Price = 80000
	book.Rating = 5
	book.Discount = 10

	err = db.Create(&book).Error
	if err != nil {
		fmt.Println("========================")
		fmt.Println("Error creating book record..")
		fmt.Println("========================")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBookshandler)

	router.Run()

}
