package handlers

import (
	"e-commerce/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCategory Creates a new category
// @Summary Creates a new category
// @Produce json
// @Consumes json
// @Param Authorization header string true "Bearer token for authentication"
// @Param type body string true "Category type"
// @Success 201 {object} models.Category "Category created successfully"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 403 {object} ErrorResponse "Forbidden"
// @Failure 409 {object} ErrorResponse "Conflict"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /categories [post]
func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInput struct {
			Type string `json:"type"`
		}
		// Bind only the specified fields from the JSON request
		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userInput.Type == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type cannot be empty"})
			return
		}

		var existingCategory models.Category
		if err := db.Where("type = ?", userInput.Type).First(&existingCategory).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Type already exists!"})
			return
		}
		newCategory := models.Category{
			Type:              userInput.Type,
			SoldProductAmount: 0,
		}
		if err := db.Create(&newCategory).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
			return
		}
		c.JSON(http.StatusCreated, newCategory)
	}
}

// GetCategories Gets all categories
// @Summary Gets all categories
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Success 200 {array} models.Category "List of categories"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /categories [get]
func GetCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.Category
		db.Preload("Products").Find(&categories)
		if len(categories) == 0 {
			c.JSON(http.StatusOK, []string{})
		} else {
			c.JSON(http.StatusOK, categories)
		}
	}
}

// UpdateCategory Updates a category type by ID
// @Summary Updates a category type by ID
// @Produce json
// @Consumes json
// @Param Authorization header string true "Bearer token for authentication"
// @Param id path int true "Category ID" Format(int64)
// @Param type body string true "Type"
// @Success 200 {object} models.Category "Updated category"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 404 {object} ErrorResponse "Category not found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /categories/{id} [patch]
func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("categoryId")

		var existingCategory models.Category
		if err := db.Where("id = ?", id).First(&existingCategory).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}

		var userInput struct {
			Type string `json:"type"`
		}

		// Bind only the specified fields from the JSON request
		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userInput.Type == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Type can't be empty"})
			return
		}

		existingCategory.Type = userInput.Type
		// Save the changes to the database
		if err := db.Save(&existingCategory).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, existingCategory)
	}
}

// DeleteCategory Deletes a category by ID
// @Summary Deletes a category by ID
// @Produce json
// @Param Authorization header string true "Bearer token for authentication"
// @Param id path int true "Category ID" Format(int64)
// @Success 200 {object} SuccessResponse "Category deleted successfully"
// @Failure 404 {object} ErrorResponse "Category not found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /categories/{id} [delete]
func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("categoryId")
		var existingCategory models.Category

		if err := db.First(&existingCategory, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Delete(&existingCategory, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Category has been successfully deleted"})
	}
}
