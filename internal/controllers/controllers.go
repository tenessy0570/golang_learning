package controllers

import (
	"github.com/gin-gonic/gin"
	"go-learn/internal/database"
	"go-learn/internal/database/models"
	"net/http"
)

func PrintAllAccounts(context *gin.Context) {
	var accountsList []models.Account
	database.DB.Find(&accountsList)
	context.JSON(http.StatusOK, accountsList)
}
