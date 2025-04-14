package middlewares

import (
	utils "example/restapi/utlis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticatie(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized", "error": err})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
