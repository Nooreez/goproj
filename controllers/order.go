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
	if order.TrackCode == "" {
		c.JSON(403, gin.H{"Bad request": "no such order"})
		return
	}
	models.DB.First(&order)
	c.JSON(200, gin.H{"Order": order})
}

func AddOrder(c *gin.Context) {
	var order models.Order

	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "user" && claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var user models.User
	models.DB.First(&user, claims.Id)

	var menu []models.Menu

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(400, gin.H{"error": menu})
		return
	}
	totalPrice := 0.0
	for _, item := range menu {
		totalPrice += item.Price
	}

	if totalPrice > user.Balance {
		c.JSON(499, gin.H{"error": "not enough money"})
		return
	}

	user.Balance -= totalPrice

	for _, item := range menu {
		var track models.Track
		track.MenuItem = item.ItemName
		track.TrackCode = order.TrackCode
		models.DB.Create(&track)
	}

	order.UserID = user.ID
	order.Status = "Pending"
	order.TrackCode, _ = utils.GenerateTrackCode("999", 10)
	models.DB.Save(&user)
	models.DB.Create(&order)

	c.JSON(200, gin.H{"success": "order created. Your track code is " + order.TrackCode, "menu": menu})
}

func UpdateStatus(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	var requestBody struct {
		TrackCode string `json:"track_code"`
		Status    string `json:"status"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	var order models.Order
	if err := models.DB.Where("track_code = ?", &requestBody.TrackCode).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "failed to find order"})
		return
	}

	switch requestBody.Status {
	case "Preparing":
		order.Status = "Preparing"
	case "Finished":
		order.Status = "Finished"
	default:
		c.JSON(400, gin.H{"error": &requestBody.TrackCode})
		return
	}

	if err := models.DB.Save(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to update status"})
		return
	}

	c.JSON(200, gin.H{"message": "status updated successfully"})
}
