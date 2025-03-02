package controllers

import (
	"go-books/repository"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var JwtKey []byte

func init() {
	JwtKey = []byte(os.Getenv("JWT_SECRET"))
	if len(JwtKey) == 0 {
		panic("JWT_SECRET is not set in environment")
	}
}

func LoginHandler(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		user, err := userRepo.GetUserByUsername(req.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// Generate JWT token
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &struct {
			Username string `json:"username"`
			UserID   int    `json:"user_id"`
			jwt.StandardClaims
		}{
			Username: user.Username,
			UserID:   user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(JwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

// RegisterHandler
func RegisterHandler(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		err := userRepo.RegisterUser(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	}
}
