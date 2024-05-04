package main

import (
	"github.com/gin-gonic/gin"
	"go-learn/internal/controllers"
	"go-learn/internal/database"
)

func main() {
	database.RecreateDb()
	//os.Exit(1)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", controllers.HelloWorld)
	router.GET("/accounts", controllers.PrintAllAccounts)
	router.POST("/login", controllers.Login)
	router.Run("localhost:8080")

}
