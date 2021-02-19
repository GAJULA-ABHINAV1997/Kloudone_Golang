package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key`
	Name   string `json:"name"`
	Auther string `json:"auther`
	Year   string `json:"year"`
}

func connectDB() *gorm.DB {

	user := "postgres"
	password := "8151"
	dbname := "golang"
	host := "127.0.0.1"
	port := "5432"

	postgres_details := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, dbname, password)

	fmt.Println("conname is\t\t", postgres_details)

	db, err := gorm.Open("postgres", postgres_details)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})

	return db
}

func getBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func addBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := Book{Name: newBook.Name, Auther: newBook.Auther, Year: newBook.Year}
	db.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func getBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func updateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&book).Updates(newBook)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func deleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func main() {
	myRouter := gin.Default()
	db := connectDB()

	myRouter.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	myRouter.GET("/books", getBooks)
	myRouter.POST("/books", addBook)
	myRouter.GET("/books/:id", getBook)
	myRouter.PATCH("/books/:id", updateBook)
	myRouter.DELETE("/books/:id", deleteBook)
	myRouter.Run(":5000")
}
