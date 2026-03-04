package middlewares

import (
	"net/http"

	"github.com/bits-and-atoms/Go_REST_API/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		// ctx.JSON(http.StatusUnauthorized, gin.H{
		// 	"message": "not authorized login first",
		// })
		// return
		//this is a middleware if we write code like above then in case of error it will stop here but we might want to go further
		//because another guy using this middleware dont need it but need the below one

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		//aborts current request
		return
	}
	uID, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized",
		})
		return
	}
	ctx.Set("userId",uID)
	ctx.Next()
	//next handler can execute correctly
}
