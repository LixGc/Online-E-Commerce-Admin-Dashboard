package handlers

import (
	"e-commerce/auth"
	"e-commerce/helpers"
	"e-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ErrorResponse represents an error response in the API.
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents a success response in the API.
type SuccessResponse struct {
	Message string `json:"message"`
}

// @Summary Register a new user
// @Produce json
// @Consumes json
// @Param email body string true "Email"
// @Param full_name body string true "Full Name"
// @Param password body string true "Password"
// @Success 201 {object} models.User "User registered successfully"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 409 {object} ErrorResponse "Email already exists"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /users/register [post]
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a struct to hold only the necessary fields
		var newUserInput struct {
			Email    string `json:"email"`
			FullName string `json:"full_name"`
			Password string `json:"password"`
		}

		// Bind only the specified fields from the JSON request
		if err := c.ShouldBindJSON(&newUserInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check for empty fields using if statements
		if newUserInput.FullName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Full Name cannot be empty"})
			return
		}

		if newUserInput.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password cannot be empty"})
			return
		}

		// Check if the password length is less than 6
		if len(newUserInput.Password) < 6 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"})
			return
		}

		// Check if the email is in a valid format
		if err := helpers.IsValidEmail(newUserInput.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Ensure the email is unique before attempting to create the user
		var existingUser models.User
		if err := db.Where("email = ?", newUserInput.Email).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
		hashedPassword, err := helpers.HashPassword(newUserInput.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		// If all checks pass, create the new user in the database
		newUser := models.User{
			Email:    newUserInput.Email,
			FullName: newUserInput.FullName,
			Password: hashedPassword,
			Balance:  0,
			Role:     "customer",
		}
		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, newUser)
	}
}

// @Summary Logs user into the system
// @Produce json
// @Consumes json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {string} Token "API token for authentication"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /users/login [post]

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check for empty email and password
		if user.Email == "" || user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password cannot be empty"})
			return
		}

		// Check if the email is in a valid format
		if err := helpers.IsValidEmail(user.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Find the user by email
		var foundUser models.User
		result := db.Where("email = ?", user.Email).First(&foundUser)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Compare the provided password with the hashed password
		if err := helpers.ComparePassword(foundUser.Password, user.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Convert user.ID to uint before passing it to GenerateToken
		token, err := auth.GenerateToken(user.Email, foundUser.Role, uint(foundUser.ID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
			return
		}

		// Respond with the generated token
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
