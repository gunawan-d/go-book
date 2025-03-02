package repository

import (
    "database/sql"
    "errors"
    "fmt"
    "go-books/structs"
)

type BookRepository struct {
    DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
    return &BookRepository{DB: db}
}

// Get all books
func (r *BookRepository) GetAllBooks() ([]structs.Book, error) {
    rows, err := r.DB.Query("SELECT id, title, author, release_year, total_page, thickness, category_id FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []structs.Book
    for rows.Next() {
        var book structs.Book
        if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ReleaseYear, &book.TotalPage, &book.Thickness, &book.CategoryID); err != nil {
            return nil, err
        }
        books = append(books, book)
    }
    return books, nil
}

// Get book by ID
func (r *BookRepository) GetBookByID(id int) (*structs.Book, error) {
    var book structs.Book
    err := r.DB.QueryRow("SELECT id, title, author, release_year, total_page, thickness, category_id FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Author, &book.ReleaseYear, &book.TotalPage, &book.Thickness, &book.CategoryID)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("book not found")
        }
        return nil, err
    }
    return &book, nil
}

// Create a new book
func (r *BookRepository) CreateBook(book structs.Book) error {
	// Check if the Book already exists
	exists, err := r.BookExists(book.Title)
	if err != nil {
		return fmt.Errorf("failed to check if book exists: %v", err)
	}
	if exists{
		return fmt.Errorf("book with title '%s' already exists", book.Title)
	}

    thickness := "tipis"
    if book.TotalPage > 100 {
        thickness = "tebal"
    }

    _, err = r.DB.Exec("INSERT INTO books (title, author, release_year, total_page, thickness, category_id) VALUES ($1, $2, $3, $4, $5, $6)", book.Title, book.Author, book.ReleaseYear, book.TotalPage, thickness, book.CategoryID)
    if err != nil {
        return fmt.Errorf("failed to insert book: %v", err)
    }
    return nil
}

// Delete book
func (r *BookRepository) DeleteBook(id int) error {
    res, err := r.DB.Exec("DELETE FROM books WHERE id = $1", id)
    if err != nil {
        return err
    }
    rowsAffected, _ := res.RowsAffected()
    if rowsAffected == 0 {
        return errors.New("book not found")
    }
    return nil
}

// Check if books exists by title
func (r * BookRepository) BookExists(title string) (bool, error) {
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE title = $1)", title).Scan(&exists)
	if err != nil {
		return false, err 
	}
	return exists, nil
}