package routes

import (
	"net/http"

	"github.com/bits-and-atoms/Go_REST_API/model"
	"github.com/bits-and-atoms/Go_REST_API/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"could not parse request data",
		})
		return
	}
	err = user.Save()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"could not register user",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":"user registered",
	})
}

func login(ctx *gin.Context){
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"message":"could not parse request data",
		})
		return
	}
	err = user.ValidateCreds()
	if err != nil{
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"message":"could not authenticate user",
		})
		return
	}
	token, err := utils.GenerateToken(user.Email,user.ID)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"message":"could not generate token",
		})
	}
	ctx.JSON(http.StatusOK,gin.H{
		"message":"login successful",
		"token": token,
	})
}