package main

import (
	"github.com/Nikita-Mamaev/template_api/controllers"
	"github.com/Nikita-Mamaev/template_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()

	models.ConnectDB()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello world"})
	})

	route.GET("/book/:id", controllers.GetBook)
	route.GET("/books", controllers.GetAllBooks)
	route.POST("/book", controllers.CreateBook)
	route.PATCH("/book/:id", controllers.UpdateBook)
	route.DELETE("/book/:id", controllers.DeleteBook)
	route.Run()
}
