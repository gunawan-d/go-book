package main

import (
	"database/sql"
	"fmt"
	"go-books/controllers"
	"go-books/database"
	"go-books/middleware"
	"go-books/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Configuration database
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set in environment")
	}

	// Initiation Connection database
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer DB.Close()

	// Check Connection database
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	// Migration database
	database.DBMigrate(DB)

	fmt.Println("Successfully connected to database!")

	// Initiation repository
	userRepo := repository.NewUserRepository(DB)
	categoryRepo := repository.NewCategoryRepository(DB)
	bookRepo := repository.NewBookRepository(DB)

	// Initiation controller
	categoryController := controllers.NewCategoryController(categoryRepo)
	bookController := controllers.NewBookController(bookRepo)

	// Initiation router Gin
	r := gin.Default()

	// Protected route dengan middleware JWT
	authRoutes := r.Group("/api")
	authRoutes.Use(middleware.JWTMiddleware())

	// Routes Category  (need auth)
	authRoutes.GET("/categories", categoryController.GetCategories)
	authRoutes.POST("/categories", categoryController.CreateCategory)
	authRoutes.GET("/categories/:id", categoryController.GetCategoryByID)
	authRoutes.DELETE("/categories/:id", categoryController.DeleteCategory)
	authRoutes.GET("/categories/:id/books", categoryController.GetBooksByCategory)

	// Book Routes (requere Authentication)
	authRoutes.GET("/books", bookController.GetAllBooks)
	authRoutes.POST("/books", bookController.CreateBook)
	authRoutes.GET("/books/:id", bookController.GetBookByID)
	authRoutes.DELETE("/books/:id", bookController.DeleteBook)

	// Routes user (not with auth)
	r.POST("/api/users/login", controllers.LoginHandler(*userRepo))
	r.POST("/api/users/register", controllers.RegisterHandler(*userRepo))

	// Running server on port 8080
	r.Run(":8080")

}
