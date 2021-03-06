package controllers

import (
	"github.com/Nikita-Mamaev/template_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type UserToBookInput struct {
	UserID uint `json:"user_id" binding:"required"`
	BookID uint `json:"book_id" binding:"required"`
}

func GetAllBooks(context *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	context.JSON(http.StatusOK, gin.H{"books": books})
}

func CreateBook(context *gin.Context) {
	var input CreateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
		Year:   input.Year,
	}
	models.DB.Create(&book)

	context.JSON(http.StatusOK, gin.H{"books": book})
}

func GetBook(context *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Книга не найдена"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"book": book})
}

func UpdateBook(context *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Книга не найдена"})
		return
	}

	var input UpdateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Update(input)

	context.JSON(http.StatusOK, gin.H{"book": book})
}

func DeleteBook(context *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Книга не найдена"})
		return
	}

	models.DB.Delete(&book)

	context.JSON(http.StatusOK, gin.H{"book": true})
}

func UserToBook(context *gin.Context) {

	var input UserToBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userToBook := models.UsersBooks{
		UserId: input.UserID,
		BookId: input.BookID,
	}
	models.DB.Create(&userToBook)

	context.JSON(http.StatusOK, gin.H{"status": true})

}
func GetUserBooks(context *gin.Context) {
	var book models.UsersBooks
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не найден"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"books": book})
}
