package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", HelloWorld)
	router.GET("/calculate", CalculatePow)
	router.Run("localhost:8080")

}
