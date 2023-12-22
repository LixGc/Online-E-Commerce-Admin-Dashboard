package main

import (
	"e-commerce/auth"
	"e-commerce/config"
	_ "e-commerce/docs"
	"e-commerce/handlers"
	"e-commerce/helpers"
	"e-commerce/models"
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := config.ConnectDatabase()
	r := gin.Default()
	insertSampleDataGorm(db)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/users/register", handlers.Register(db))
	r.POST("/users/login", handlers.Login(db))
	r.PATCH("/users/topup", auth.AuthenticationMiddleware(), handlers.UpdateUserBalance(db))
	r.GET("/categories", auth.AuthorizationMiddleware(), handlers.GetCategories(db))
	r.POST("/categories", auth.AuthorizationMiddleware(), handlers.CreateCategory(db))
	r.PATCH("/categories/:categoryId", auth.AuthorizationMiddleware(), handlers.UpdateCategory(db))
	r.DELETE("/categories/:categoryId", auth.AuthorizationMiddleware(), handlers.DeleteCategory(db))
	r.GET("/products", auth.AuthenticationMiddleware(), handlers.GetProducts(db))
	r.POST("/products", auth.AuthorizationMiddleware(), handlers.CreateProduct(db))
	r.PUT("/products/:productId", auth.AuthorizationMiddleware(), handlers.UpdateProduct(db))
	r.DELETE("/products/:productId", auth.AuthorizationMiddleware(), handlers.DeleteProduct(db))
	r.POST("/transactions", auth.AuthenticationMiddleware(), handlers.CreateTransaction(db))
	r.GET("/transactions/my-transactions", auth.AuthenticationMiddleware(), handlers.GetMyTransaction(db))
	r.GET("/transactions/user-transactions", auth.AuthorizationMiddleware(), handlers.GetTransaction(db))
	r.Run(":8080")
}
func insertSampleDataGorm(db *gorm.DB) {
	sampleUsers := []models.User{
		{FullName: "Admin", Email: "admin@example.com", Role: "admin", Balance: 100000, Password: "adminpassword"},
	}

	for _, user := range sampleUsers {
		hashedPassword, err := helpers.HashPassword(user.Password)
		if err != nil {
			log.Fatal(err)
		}

		newUser := models.User{
			FullName: user.FullName,
			Email:    user.Email,
			Password: hashedPassword,
			Balance:  user.Balance,
			Role:     user.Role,
		}

		result := db.Create(&newUser)
		if result.Error != nil {
			log.Fatalf("Error creating user: %v", result.Error)
		}

		fmt.Println("User created successfully:", newUser.ID)
	}
	fmt.Println("Insert data success")
}
