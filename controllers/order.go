package controllers

import (
	"goproj/models"
	"goproj/utils"
	"net/http"

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
	totalPrice := 0.0
	for _, item := range order.Menu {
		totalPrice += item.Price
	}

	var user models.User
	models.DB.First(&user, claims.Id)
	if totalPrice > user.Balance {
		c.JSON(499, gin.H{"error": "not enough money"})
		return
	}

	user.Balance -= totalPrice
	order.TrackCode = utils.GenerateTrackCode("999", 10)
	order.Status = "Pending"

	models.DB.Create(&order)

	models.DB.Save(&user)

	c.JSON(200, gin.H{"success": "order created. Your track code is " + order.TrackCode})
}

func DeleteOrder(c *gin.Context) {
	trackCode := c.Param("trackCode")
	var order models.Order
	result := models.DB.Where("track_code = ?", trackCode).Delete(&order)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "failed to delete order"})
		return
	}
	c.JSON(200, gin.H{"success": "Deleted successfully"})
}

func UpdateStatus(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)
	if err != nil || claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized as admin"})
		return
	}

	trackCode := c.Param("trackCode")
	status := c.Query("status")

	var order models.Order
	if err := models.DB.Where("track_code = ?", trackCode).First(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find order"})
		return
	}

	switch status {
	case "Preparing":
		order.Status = "Preparing"
	case "Finished":
		order.Status = "Finished"
	default:
		c.JSON(400, gin.H{"error": "invalid status"})
		return
	}

	if err := models.DB.Save(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status updated successfully"})
}
