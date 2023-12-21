package handlers

import (
	"e-commerce/helpers"
	"e-commerce/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Get user's transactions
// @Description Retrieve transactions for the authenticated user
// @Tags Transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} models.TransactionHistory "List of user's transactions"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 404 {object} ErrorResponse "Transactions not found for the user"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /transactions/my-transactions [get]
func GetMyTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User id not found in the context"})
			return
		}

		var transactions []models.TransactionHistory
		if err := db.Preload("Product").Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Transactions not found for the user"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transactions)
	}
}

// @Summary Get all transactions
// @Description Retrieve all transactions (admin access)
// @Tags Transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Security ApiKeyAuth
// @Success 200 {array} models.TransactionHistory "List of all transactions"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /transactions/user-transactions [get]
func GetTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var TransactionHistory []models.TransactionHistory
		db.Preload("Products").Find(&TransactionHistory)
		if len(TransactionHistory) == 0 {
			c.JSON(http.StatusOK, []string{})
		} else {
			c.JSON(http.StatusOK, TransactionHistory)
		}
	}
}

// @Summary Create a new transaction
// @Description Purchase a product and create a transaction record
// @Tags Transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param product_id body int true "Product ID to purchase"
// @Param quantity body int true "Quantity of the product to purchase"
// @Success 200 {string} string "Purchase successfull"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 404 {object} ErrorResponse "Product not found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /transactions [post]
func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInput struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		}
		email, exists := c.Get("email")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User email not found in the context"})
			return
		}
		// Bind only the specified fields from the JSON request
		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userInput.Quantity == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity can't be 0 or empty"})
			return
		}
		if userInput.ProductID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product can't be empty"})
			return
		}
		var existingProduct models.Product
		if err := db.Where("id = ?", userInput.ProductID).First(&existingProduct).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		if existingProduct.Stock == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product is out of stock"})
			return
		}
		if userInput.Quantity > existingProduct.Stock {
			errorMessage := fmt.Sprintf("Insufficient stock. Only %d stocks left.", existingProduct.Stock)
			c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
			return
		}
		var existingUser models.User
		if err := db.Where("email = ?", email).First(&existingUser).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		totalPrice := userInput.Quantity * existingProduct.Price
		if existingUser.Balance < totalPrice {
			errorMessage := fmt.Sprintf("Insufficient balance. Total price: %s, your balance: %s", helpers.FormatRupiah(totalPrice), helpers.FormatRupiah(existingUser.Balance))
			c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
			return
		}
		existingProduct.Stock -= userInput.Quantity
		if err := db.Save(&existingProduct).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		existingUser.Balance -= totalPrice
		if err := db.Save(&existingUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		updatedTransaction := models.TransactionHistory{
			ProductID:  existingProduct.ID,
			UserID:     existingUser.ID,
			Quantity:   userInput.Quantity,
			TotalPrice: totalPrice,
		}
		if err := db.Create(&updatedTransaction).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var existingCategory models.Category
		if err := db.Model(&existingProduct).Association("Category").Find(&existingCategory); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		existingCategory.SoldProductAmount += userInput.Quantity

		if err := db.Save(&existingCategory).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "You have successfully purchased the product",
			"transaction_bill": gin.H{
				"total_price":   totalPrice,
				"quantity":      userInput.Quantity,
				"product_title": existingProduct.Title,
			},
		})
	}
}
