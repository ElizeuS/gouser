package controllers

import (
	"github.com/ElizeuS/gouser/database"
	"github.com/ElizeuS/gouser/models"
	"github.com/gin-gonic/gin"
)

func ShowUser(c *gin.Context) {
	id := c.Param("id")

	db := database.GetDatabase()

	var user models.User
	err := db.First(&user, id).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user",
		})

		return
	}

	c.JSON(200, user)

}

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "cannot create user: " + err.Error(),
		})

		return
	}

	c.JSON(200, user.ID)

}

func ShowUsers(c *gin.Context) {
	db := database.GetDatabase()

	var users []models.User
	err := db.Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "cannot list users: " + err.Error(),
		})
	}
}
