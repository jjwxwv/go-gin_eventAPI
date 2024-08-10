package middlewares

import (
	"fmt"
	"net/http"

	"example.com/project/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	//get token from incoming request
	//token is a part of header
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		//middlewarte is execute in the middle of a request
		//AbortWithStatusJSON is abort the current request if somthing went wrong and no other code on the server runs
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message": "Not authorized"})
		return
	}
	//verify token
	userId, err := utils.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message": "Not authorized"})
		return
	}
	//Set method is allow to add some data to context value and can be used anywhere where the context is available
	context.Set("userId", userId)
	context.Next()
}