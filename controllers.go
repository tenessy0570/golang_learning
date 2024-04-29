package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

func CalculatePow(c *gin.Context) {
	var result = 1
	go func() {
		var myVar = math.Pow(2, 64)
		fmt.Println(myVar)
	}()

	c.JSON(http.StatusOK, result)
}
