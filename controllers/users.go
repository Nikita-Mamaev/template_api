package controllers

import (
	"github.com/Nikita-Mamaev/template_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
	Book string `json:"book"`
}

type UpdateUserInput struct {
	Name string `json:"name"`
	Book string `json:"book"`
}

func GetAllUsers(context *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	context.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(context *gin.Context) {
	var input CreateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Name: input.Name,
		Book: input.Book,
	}
	models.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUser(context *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не найден"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(context *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не найден"})
		return
	}

	var input UpdateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Update(input)

	context.JSON(http.StatusOK, gin.H{"user": user})
}

func DeleteUser(context *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", context.Param("id")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не найден"})
		return
	}

	models.DB.Delete(&user)

	context.JSON(http.StatusOK, gin.H{"user": true})
}
