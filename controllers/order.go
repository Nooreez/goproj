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
	order.TrackCode := generateTrackCode("999",10)
	models.DB.Create(&order)

	c.JSON(200, gin.H{"success": "order created. Your track code is " + order.TrackCode})
}

func DeleteMenu(c *gin.Context) {
	var trackCode string
	
	if err := c.ShouldBindJSON(&track_code); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	order := models.DB.First()

	result := models.DB.Delete(&models.Menu{}, menu)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "menu not found"})
		return
	}
	c.JSON(200, gin.H{"OK": "menu deleted"})
}