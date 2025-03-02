package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"go-books/structs"
)

type CategoryRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

// Get all categories
func (r *CategoryRepository) GetAllCategories() ([]structs.Category, error) {
	rows, err := r.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []structs.Category
	for rows.Next() {
		var category structs.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// Get category by ID
func (r *CategoryRepository) GetCategoryByID(id int) (*structs.Category, error) {
    var category structs.Category
    err := r.DB.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("Category with ID %d not found", id)
            return nil, errors.New("category not found")
        }
        // log.Printf("Error querying category with ID %d: %v", id, err)
        return nil, err
    }
    // log.Printf("Category with ID %d found: %+v", id, category)
    return &category, nil
}

// Create a new category
func (r *CategoryRepository) CreateCategory(name string, createdBy int) error {
	_, err := r.DB.Exec("INSERT INTO categories (name, created_by) VALUES ($1, $2)", name, createdBy)
	if err != nil {
		return fmt.Errorf("failed to insert category: %v", err)
	}
	return nil
}


// Delete category
func (r *CategoryRepository) DeleteCategory(id int) error {
	res, err := r.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}

// Get books by category
func (r *CategoryRepository) GetBooksByCategory(id int) ([]string, error) {
	rows, err := r.DB.Query("SELECT title FROM books WHERE category_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []string
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			return nil, err
		}
		books = append(books, title)
	}
	return books, nil
}

// Check if category exists by name
func (r *CategoryRepository) CategoryExists(name string) (bool, error) {
    var exists bool
    err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE name = $1)", name).Scan(&exists)
    if err != nil {
        log.Printf("Error checking if category exists: %v", err)
        return false, err
    }
    log.Printf("Category %s exists: %v", name, exists)
    return exists, nil
}
