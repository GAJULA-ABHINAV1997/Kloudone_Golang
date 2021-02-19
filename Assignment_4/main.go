package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
)

type Car struct {
	ID    uint   `json:"id" gorm:"primary_key`
	Name  string `json:"name"`
	Model string `json:"model"`
}

type CreateCarInput struct {
	Name  string `json:"name" binding:"required"`
	Model string `json:"model" binding:"required"`
}
type UpdateCarInput struct {
	Name  string `json:"name"`
	Model string `json:"model"`
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
	db.AutoMigrate(&Car{})

	return db
}

// GET /cars
// Get all cars
func FindCars(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cars []Car
	db.Find(&cars)
	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// POST /cars
// Create new cars
func CreateCar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input CreateCarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create Car
	car := Car{Name: input.Name, Model: input.Model}
	db.Create(&car)
	c.JSON(http.StatusOK, gin.H{"data": car})
}

// GET /cars/:id
// Find a car
func FindCar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var car Car
	if err := db.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": car})
}

// PATCH /cars/:id
// Update a car
func UpdateCar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var car Car
	if err := db.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input UpdateCarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&car).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": car})
}

// DELETE /cars/:id
// Delete a car
func DeleteCar(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var car Car
	if err := db.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&car)
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
	r.GET("/cars", FindCars)
	r.POST("/cars", CreateCar)       // create
	r.GET("/cars/:id", FindCar)      // find by id
	r.PATCH("/cars/:id", UpdateCar)  // update by id
	r.DELETE("/cars/:id", DeleteCar) // delete by id
	r.Run(":8000")
}
