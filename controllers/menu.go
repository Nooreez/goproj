package controllers

import (
	"goproj/models"
	"goproj/utils"

	"github.com/gin-gonic/gin"
)

func ShowMenu(c *gin.Context) {
	var menu []models.Menu

	models.DB.Find(&menu)

	c.JSON(200, gin.H{"Menu": menu})
}

func AddMenu(c *gin.Context) {

	var menu models.Menu

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized as admin"})
		return
	}

	var existingMenu models.Menu

	models.DB.Where("item_name = ?", menu.ItemName).First(&existingMenu)

	if existingMenu.ID != 0 {
		c.JSON(400, gin.H{"error": "it's already in menu list"})
		return
	}

	models.DB.Create(&menu)

	c.JSON(200, gin.H{"success": "menu created"})
}

func DeleteMenu(c *gin.Context) {
	var menu models.Menu

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized as admin"})
		return
	}

	result := models.DB.Delete(&models.Menu{}, menu)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "menu not found"})
		return
	}
	c.JSON(200, gin.H{"OK": "menu deleted"})
}
