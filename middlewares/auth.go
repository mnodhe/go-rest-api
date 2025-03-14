package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/utils"
	"net/http"
	"strings"
)

func Authenticate(context *gin.Context) {
	token := strings.TrimPrefix(context.Request.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	tokenClaims, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	context.Set("userId", tokenClaims["userId"])
	context.Next()
}
