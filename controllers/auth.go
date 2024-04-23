package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jwt-gin/models"
	"github.com/jwt-gin/utils"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	var input RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}
	user.UserName = input.Username
	user.Password = input.Password
	_, err := user.SaveUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(ctx *gin.Context) {
	var input LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}
	user.UserName = input.Username
	user.Password = input.Password

	token, err := models.LoginCheck(user.UserName, user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username or password incorrect"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

func CurrentUser(ctx *gin.Context) {
	user_id, err := utils.ExtractTokenID(ctx)
	log.Println(user_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := models.GetUserByID(user_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func GetUsers(ctx *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
