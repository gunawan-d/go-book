package repository

import (
    "database/sql"
    "errors"
    "time"

    "golang.org/x/crypto/bcrypt"
    "go-books/structs"
)

// UserRepository struct
type UserRepository struct {
    DB *sql.DB
}

func (u *UserRepository) RegisterUser(username string, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    createdAt := time.Now().Format("2006-01-02 15:04:05")

    query := `INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)`
    _, err = u.DB.Exec(query, username, string(hashedPassword), createdAt)
    if err != nil {
        return err
    }

    return nil
}

// NewUserRepository creates an instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{DB: db}
}

// GetUserByUsername retrieves a user by username
func (u *UserRepository) GetUserByUsername(username string) (*structs.User, error) {
    user := structs.User{}
    err := u.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).
        Scan(&user.ID, &user.Username, &user.Password)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    return &user, nil
}