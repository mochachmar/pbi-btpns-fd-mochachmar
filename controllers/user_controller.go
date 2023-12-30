package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/app"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/database"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/helpers"
)

func RegisterUser(c *gin.Context) {
	var newUser app.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := helpers.ValidateUser(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.Password = helpers.HashPassword(newUser.Password)

	database.DB.Create(&newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
	var user app.User
	var loginCredentials app.User

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Where("email = ?", loginCredentials.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !helpers.VerifyPassword(user.Password, loginCredentials.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
	userID := helpers.GetUserIDFromToken(c)

	var updatedUser app.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser app.User
	result := database.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Model(&existingUser).Updates(&updatedUser)

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	userID := helpers.GetUserIDFromToken(c)

	var user app.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user, userID)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
