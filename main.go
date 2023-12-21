package main

import (
	"e-commerce/auth"
	"e-commerce/config"
	"e-commerce/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := config.ConnectDatabase()
	r := gin.Default()
	r.POST("/users/register", handlers.Register(db))
	r.POST("/users/login", handlers.Login(db))
	r.PATCH("/users/topup", auth.AuthenticationMiddleware(), handlers.UpdateUserBalance(db))
	r.GET("/categories", auth.AuthorizationMiddleware(), handlers.GetCategories(db))
	r.POST("/categories", auth.AuthorizationMiddleware(), handlers.CreateCategory(db))
	r.PATCH("/categories/:id", auth.AuthorizationMiddleware(), handlers.UpdateCategory(db))
	r.DELETE("/categories/:id", auth.AuthorizationMiddleware(), handlers.DeleteCategory(db))
	r.GET("/products", auth.AuthenticationMiddleware(), handlers.GetProducts(db))
	r.POST("/products", auth.AuthorizationMiddleware(), handlers.CreateProduct(db))
	r.PUT("/products/:id", auth.AuthorizationMiddleware(), handlers.UpdateProduct(db))
	r.DELETE("/products/:id", auth.AuthorizationMiddleware(), handlers.DeleteProduct(db))
	r.POST("/transactions", auth.AuthenticationMiddleware(), handlers.CreateTransaction(db))
	r.GET("/transactions/my-transactions", auth.AuthenticationMiddleware(), handlers.GetMyTransaction(db))
	r.GET("/transactions/user-transactions", auth.AuthorizationMiddleware(), handlers.GetTransaction(db))
	r.Run(":8080")
}
