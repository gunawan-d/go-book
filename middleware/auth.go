package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"

)

var JwtKey []byte

func init() {
	JwtKey = []byte(os.Getenv("JWT_SECRET"))
	if len(JwtKey) == 0 {
		panic("JWT_SECRET is not set in environment")
	}
}
func extractToken(c *gin.Context) string {
    bearerToken := c.GetHeader("Authorization")
    if len(bearerToken) > 7 && strings.ToUpper(bearerToken[0:7]) == "BEARER " {
        return bearerToken[7:]
    }
    return ""
}

func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := extractToken(c)
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
            c.Abort()
            return
        }

        claims := &struct {
            Username string `json:"username"`
            UserID   int    `json:"user_id"`
            jwt.StandardClaims
        }{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return JwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
            c.Abort()
            return
        }

        // Set user ID in context
        c.Set("user_id", claims.UserID)
        c.Next()
    }
}

// validateToken Verify token JWT
func validateToken(tokenString, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
