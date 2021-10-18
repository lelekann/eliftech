package controllers

import (
	"eliftech_1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GET posts/
//Get all posts

func FindPosts(c *gin.Context) {
	var posts []models.Post

	if err := models.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

//POST posts/
//Create new post

func CreatePost(c *gin.Context) {
	var input models.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Content: input.Content, CreatedBY: input.CreatedBy}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

//GET posts/:id
//Get single post

func FindPost(c *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id=?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

//PUT /posts/:id
//Update post

func UpdatePost(c *gin.Context) {
	var post models.Post

	if err := models.DB.Where("id=?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found!"})
		return
	}

	var input models.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&post).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

//DELETE posts/:id
//Delete post

func DeletePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id=?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found!"})
		return
	}

	models.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
