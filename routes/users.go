package routes

import (
	"example/restapi/models"
	utils "example/restapi/utlis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the requested data", "error": err})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created sucessfully."})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the requested data", "error": err})
		return
	}

	err = user.ValidateCredentails()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user", "error": err})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Sucessful!!", "token": token})
}
