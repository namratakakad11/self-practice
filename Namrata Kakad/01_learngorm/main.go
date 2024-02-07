package main

import (
	"example/01_LEARNGORM/models"
	"example/01_LEARNGORM/storage"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// type Book struct {
// 	Author    string `json:"author`
// 	Title     string `json:"title"`
// 	Publisher string `json:"publisher`
// }

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) createBook(c *gin.Context) {
	book := models.Books{}
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Post request not created"})

	}
	err = r.DB.Create(&book).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Post request not created"})

	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "added"})

}

func (r *Repository) getbooks(c *gin.Context) {
	bookmodel := &[]models.Books{}

	err := r.DB.Find(bookmodel).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not get books"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "books fetch successfully", "data": bookmodel})

}

func (r *Repository) getBookById(c *gin.Context) {
	id := c.Param("id")
	bookmodel := &models.Books{}
	if id != "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "no id found"})

	}

	err := r.DB.Where("id =?", id).First(bookmodel).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "could not get id"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "id fetched", "data": bookmodel})

}

func (r *Repository) deleteBook(c *gin.Context) {

	bookmodel := models.Books{}
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "no id found"})

	}
	err := r.DB.Delete(bookmodel, id)
	if err.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "could not delete the id"})

	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": " deleted successfully"})

}

func (r *Repository) setuproutes(app *gin.Engine) {
	api := app.Group("/api")
	api.POST("/createbooks", r.createBook)
	api.DELETE("/deletebook/:id", r.deleteBook)
	api.GET("/getbooksbyid/:id", r.getBookById)
	api.GET("/getbooks", r.getbooks)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inside main")

	config := &storage.Config{

		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_User"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Host:     os.Getenv("DB_HOST"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		fmt.Println("inside error")
		log.Fatal(err)
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal(err)
	}

	r := Repository{
		DB: db,
	}

	app := gin.Default()
	r.setuproutes(app)
	app.Run("localhost:8080")

}
