package controllers

import (
	"eliftech_1/models"
	"eliftech_1/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GET users/
//Get all users

func FindUsers(c *gin.Context) {
	var s services.UserService
	res, err := s.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Users not found!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

//POST users/
//Create new user

func CreateUser(c *gin.Context) {
	var s services.UserService
	var input models.CreateUserInput
	var user *models.User
	user.Name = input.Name
	user.Password = input.Password
	user.Email = input.Email

	res, err := s.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User cannot created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

//GET users/:id
//Get single user

func FindUser(c *gin.Context) {
	var s services.UserService
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)

	res, err := s.GetUserById(uint(uintId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

// //PUT users/:id
// //Update user

// func UpdateUser(c *gin.Context) {
// 	var user models.User

// 	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
// 		return
// 	}

// 	var input models.CreateUserInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	models.DB.Model(&user).Updates(input)

// 	c.JSON(http.StatusOK, gin.H{"data": user})
// }

//DELETE users/:id
//Delete user

func DeleteUser(c *gin.Context) {
	var s services.UserService
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)

	err = s.DeleteUser(uint(uintId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
