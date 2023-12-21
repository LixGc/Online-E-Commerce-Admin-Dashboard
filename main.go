package main

import (
	"e-commerce/auth"
	"e-commerce/config"
	_ "e-commerce/docs"
	"e-commerce/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := config.ConnectDatabase()
	r := gin.Default()
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
