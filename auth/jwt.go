package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("rahasia")

type Claims struct {
	Email string `json:"username"`
	Role  string `json:"role"`
	ID    uint   `json:"ID"`
	jwt.StandardClaims
}

func GenerateToken(email, role string, id uint) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Email: email,
		Role:  role,
		ID:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, err
}

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, err := ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get claims from token"})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("id", claims.ID)
		c.Next()
	}
}
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, err := ValidateToken(tokenString)
		if err != nil || !token.Valid {
			fmt.Println("Invalid Token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			fmt.Println("Failed to get claims from token")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get claims from token"})
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			fmt.Println("Unauthorized role:", claims.Role) // Add this line for debugging
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("id", claims.ID)
		c.Next()
	}
}
