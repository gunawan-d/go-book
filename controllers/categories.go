package controllers

import (
	"net/http"
	"strconv"
	"go-books/repository"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	Repo *repository.CategoryRepository
}

func NewCategoryController(repo *repository.CategoryRepository) *CategoryController {
	return &CategoryController{Repo: repo}
}

// Get all categories
func (cc *CategoryController) GetCategories(c *gin.Context) {
	categories, err := cc.Repo.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// Get category by ID
func (cc *CategoryController) GetCategoryByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
        return
    }

	category, err := cc.Repo.GetCategoryByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"category": category})
}


// Create category
func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	exists, err := cc.Repo.CategoryExists(request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error check category exist"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Category already exists"})
		return
	}

	// Get user ID from JWT
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDInt, ok := userID.(int) // Pastikan tipe data sesuai
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id type"})
		return
	}

	if err := cc.Repo.CreateCategory(request.Name, userIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
 
	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})

}

// Delete category
func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	err = cc.Repo.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// Get books by category
func (cc *CategoryController) GetBooksByCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	books, err := cc.Repo.GetBooksByCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

