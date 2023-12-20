package main

import (
	_ "challenge-2/docs"
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

	r.Run(":8080")
}
