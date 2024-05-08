package controllers

import (
	"goproj/models"
	"goproj/utils"

	"github.com/gin-gonic/gin"
)

func ShowOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	models.DB.First(&order)

	c.JSON(200, gin.H{"Order": order})
}

func AddOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if claims.Role != "user" && claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	order.TrackCode = utils.GenerateTrackCode("999", 10)
	models.DB.Create(&order)

	c.JSON(200, gin.H{"success": "order created. Your track code is " + order.TrackCode})
}
