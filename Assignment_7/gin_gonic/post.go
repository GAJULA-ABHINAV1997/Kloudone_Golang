package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

type Post struct {
	ID     uint   `json:"id" gorm:"primary_key`
	Name   string `json:"name"`
	Auther string `json:"auther"`
}

type CreatePostInput struct {
	Name   string `json:"name" binding:"required"`
	Auther string `json:"auther" binding:"required"`
}
type UpdatePostInput struct {
	Name   string `json:"name"`
	Auther string `json:"auther"`
}

func SetupModels() *gorm.DB {

	user := "postgres"
	password := "8151"
	dbname := "golang"
	host := "127.0.0.1"
	port := "5432"

	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, dbname, password)

	fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Post{})

	return db
}

// GET /posts
// Get all posts
func FindPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var posts []Post
	db.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// POST /posts
// Create new posts
func CreatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create post
	post := Post{Name: input.Name, Auther: input.Auther}
	db.Create(&post)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// GET /posts/:id
// Find a post
func FindPost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var post Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// PATCH /post/:id
// Update a post
func UpdatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var post Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&post).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DELETE /posts/:id
// Delete a post
func DeletePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var post Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func main() {
	r := gin.Default()
	db := SetupModels() // new
	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.Use(static.Serve("/", static.LocalFile("./views", true)))
	r.GET("/posts", FindPosts)
	r.POST("/posts", CreatePost)       // create
	r.GET("/posts/:id", FindPost)      // find by id
	r.PATCH("/posts/:id", UpdatePost)  // update by id
	r.DELETE("/posts/:id", DeletePost) // delete by id
	r.Run(":7000")
}
