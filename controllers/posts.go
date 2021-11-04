package controllers

import (
	"eliftech_1/models"
	"eliftech_1/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GET posts/
//Get all posts

func FindPosts(c *gin.Context) {
	var s services.PostService
	res, err := s.GetPosts()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Posts not found!",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

//POST posts/
//Create new post

func CreatePost(c *gin.Context) {
	var s services.PostService
	var input models.CreatePostInput
	var post *models.Post

	post.Content = input.Content
	post.CreatedBY = input.CreatedBy

	res, err := s.CreatePost(post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Post cannot created",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

//GET posts/:id
//Get single post

func FindPost(c *gin.Context) {
	var s services.PostService
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)

	res, err := s.GetPostById(uint(uintId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Post not found!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

//DELETE posts/:id
//Delete post

func DeletePost(c *gin.Context) {
	var s services.UserService
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)

	err = s.DeleteUser(uint(uintId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Post not found!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
