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

	route.GET("/user/:id", controllers.GetUser)
	route.GET("/users", controllers.GetAllUsers)
	route.POST("/user", controllers.CreateUser)
	route.PATCH("/user/:id", controllers.UpdateUser)
	route.DELETE("/user/:id", controllers.DeleteUser)

	route.POST("usertobook", controllers.UserToBook)
	route.Run()
}
