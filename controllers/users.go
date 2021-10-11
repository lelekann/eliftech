package controllers

import (
	"eliftech_1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GET users/
//Get all users

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

//POST users/
//Create new user

func CreateUser(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

//GET users/:id
//Get single user

func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

//PUT users/:id
//Update user

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

//DELETE users/:id
//Delete user

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
