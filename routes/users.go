package routes

import (
	"fmt"
	"net/http"

	"example.com/project/models"
	"example.com/project/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest,gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not save user."})
		return
	}
	context.JSON(http.StatusCreated,gin.H{"message": "User created!", "user":user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest,gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized,gin.H{"message": err.Error()})
		return
	}
	//generate token after login is success
	token, err := utils.GenerateToken(user.Email,user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError,gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK,gin.H{"message": "Login successful", "token": token})
}