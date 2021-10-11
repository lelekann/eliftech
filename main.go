package main

import (
	"eliftech_1/controllers"
	"eliftech_1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello word"})
	})

	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("users/:id", controllers.DeleteUser)

	r.GET("/posts", controllers.FindPosts)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/:id", controllers.FindPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
