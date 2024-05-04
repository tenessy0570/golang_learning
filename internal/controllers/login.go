package controllers

import (
	"github.com/gin-gonic/gin"
	"go-learn/internal/database"
	"go-learn/internal/database/models"
	"go-learn/internal/utils"
	"net/http"
)

type LoginInfo struct {
	Login    string
	Password string
}

// Login returns bearer token bound to user whose credentials were passed in post body
func Login(context *gin.Context) {
	var loginInfo LoginInfo
	var bindError = context.ShouldBindJSON(&loginInfo)

	if bindError != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "provide login and password fields in post body"},
		)
		return
	}

	if len(loginInfo.Login) == 0 || len(loginInfo.Password) == 0 {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "values should not be empty"},
		)
		return
	}

	var foundUser models.Account
	database.DB.
		Where("login = ?", loginInfo.Login).
		Where("password = ?", loginInfo.Password).
		First(&foundUser)

	if len(foundUser.Login) == 0 {
		context.JSON(
			http.StatusNotFound,
			gin.H{"message": "bad credentials"},
		)
		return
	}

	var newToken = utils.CreateRandomString(32)
	foundUser.BearerToken = newToken
	database.DB.Save(&foundUser)

	context.JSON(
		http.StatusOK,
		gin.H{"token": newToken},
	)
}
