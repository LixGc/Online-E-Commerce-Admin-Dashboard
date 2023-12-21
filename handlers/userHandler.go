package handlers

import (
	"e-commerce/helpers"
	"e-commerce/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Updates User's balance
// @Summary Updates the user's balance
// @Produce json
// @Consumes json
// @Param email body int true "Balance"
// @Header 200 {string} Token "API token for authentication"
// @Success 201 {object} models.User "User registered successfully"
// @Router /users/register [post]
func UpdateUserBalance(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve user email from the context
		userEmail, exists := c.Get("email")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User email not found in the context"})
			return
		}

		// Retrieve existing user from the database using email
		var existingUser models.User
		if err := db.Where("email = ?", userEmail).First(&existingUser).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Create a struct to hold only the necessary fields for update
		var userInput struct {
			Balance int `json:"balance"`
		}

		// Bind only the specified fields from the JSON request
		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userInput.Balance == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Balance cannot be empty or zero"})
			return
		}

		// Custom validation for the 'Balance' field
		if err := helpers.ValidateBalance(userInput.Balance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update the user's balance
		existingUser.Balance = userInput.Balance

		// Save the changes to the database
		if err := db.Save(&existingUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Respond with the updated user
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Your balance has been successfully updated to Rp %d", existingUser.Balance)})
	}
}
