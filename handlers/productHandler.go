package handlers

import (
	"e-commerce/helpers"
	"e-commerce/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInput struct {
			Title      string `json:"title"`
			Price      int    `json:"price"`
			Stock      int    `json:"stock"`
			CategoryID int    `json:"category_id"`
		}
		// Bind only the specified fields from the JSON request
		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userInput.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
			return
		}
		if userInput.Price == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Price can't be empty or zero"})
			return
		}
		// Custom validation for the 'Price' field
		if err := helpers.ValidatePrice(userInput.Price); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userInput.Stock == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Stock can't be empty or zero"})
			return
		}
		if userInput.CategoryID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category can't be empty or zero"})
			return
		}
		var existingProduct models.Product
		if err := db.Where("title = ?", userInput.Title).First(&existingProduct).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Title already exists!"})
			return
		}
		newProduct := models.Product{
			Title:      userInput.Title,
			Price:      userInput.Price,
			Stock:      userInput.Stock,
			CategoryID: uint(userInput.CategoryID),
		}
		if err := db.Create(&newProduct).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}
		c.JSON(http.StatusCreated, newProduct)
	}
}
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		db.Preload("Products").Find(&products)
		if len(products) == 0 {
			c.JSON(http.StatusOK, []string{})
		} else {
			c.JSON(http.StatusOK, products)
		}
	}
}
func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var existingProduct models.Product
		if err := db.Where("id = ?", id).First(&existingProduct).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		var userInput struct {
			Title      string `json:"title"`
			Price      int    `json:"price"`
			Stock      int    `json:"stock"`
			CategoryID int    `json:"category_id"`
		}

		// Bind only the specified fields from the JSON request
		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userInput.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
			return
		}
		if userInput.Price == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Price can't be empty or zero"})
			return
		}
		// Custom validation for the 'Price' field
		if err := helpers.ValidatePrice(userInput.Price); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userInput.Stock == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Stock can't be empty or zero"})
			return
		}
		if userInput.CategoryID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category can't be empty or zero"})
			return
		}

		existingProduct.Title = userInput.Title
		existingProduct.Price = userInput.Price
		existingProduct.Stock = userInput.Stock
		existingProduct.CategoryID = uint(userInput.CategoryID)
		// Save the changes to the database
		if err := db.Save(&existingProduct).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		type ProductResponse struct {
			Title      string `json:"title"`
			Price      string `json:"price"`
			Stock      int    `json:"stock"`
			CategoryID int    `json:"category_id"`
		}
		response := ProductResponse{
			Title:      existingProduct.Title,
			Price:      helpers.FormatRupiah(existingProduct.Price),
			Stock:      existingProduct.Stock,
			CategoryID: int(existingProduct.CategoryID),
		}
		c.JSON(http.StatusOK, gin.H{"product": response})
	}
}
func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var existingProduct models.Product

		if err := db.First(&existingProduct, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := db.Delete(&existingProduct, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product has been successfully deleted"})
	}
}
