package handlers

import (
	"e-commerce/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

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
func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
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
