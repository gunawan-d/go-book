package controllers

import (
    "net/http"
    "strconv"
    "go-books/repository"
    "go-books/structs"

    "github.com/gin-gonic/gin"
)

type BookController struct {
    Repo *repository.BookRepository
}

func NewBookController(repo *repository.BookRepository) *BookController {
    return &BookController{Repo: repo}
}

// Get all books
func (bc *BookController) GetAllBooks(c *gin.Context) {
    books, err := bc.Repo.GetAllBooks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving books"})
        return
    }
    c.JSON(http.StatusOK, books)
}

// Get book by ID
func (bc *BookController) GetBookByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    book, err := bc.Repo.GetBookByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"book": book})
}

// Create book
func (bc *BookController) CreateBook(c *gin.Context) {
    var book structs.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
        return
    }

    if err := bc.Repo.CreateBook(book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

// Delete book
func (bc *BookController) DeleteBook(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    err = bc.Repo.DeleteBook(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}