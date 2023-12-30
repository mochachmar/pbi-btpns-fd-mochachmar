package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/app"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/database"
)

func CreatePhoto(c *gin.Context) {
	var newPhoto app.Photo

	if err := c.ShouldBindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newPhoto)

	c.JSON(http.StatusCreated, gin.H{"message": "Photo created successfully"})
}

func GetPhotos(c *gin.Context) {
	var photos []app.Photo

	database.DB.Find(&photos)

	c.JSON(http.StatusOK, photos)
}

func GetPhotoByID(c *gin.Context) {
	photoID := c.Param("photoId")

	var photo app.Photo

	result := database.DB.First(&photo, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func UpdatePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	var updatedPhoto app.Photo

	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingPhoto app.Photo
	result := database.DB.First(&existingPhoto, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	database.DB.Model(&existingPhoto).Updates(&updatedPhoto)

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

func DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	var photo app.Photo
	result := database.DB.First(&photo, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	database.DB.Delete(&photo, photoID)

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
