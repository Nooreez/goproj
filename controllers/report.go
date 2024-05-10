package controllers

import (
	"goproj/models"
	"goproj/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRating(c *gin.Context) {
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

	var order models.Order
	if err := models.DB.Where("track_code = ?", c.Param("trackCode")).First(&order).Error; err != nil {
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	if order.Status != "Finished" {
		c.JSON(400, gin.H{"error": "order status must be finished to update rating"})
		return
	}

	claimsID, err := strconv.ParseUint(claims.Id, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	if order.UserID != uint(claimsID) {
		c.JSON(403, gin.H{"error": "you are not allowed to update the rating for this order"})
		return
	}

	var rating int
	if err := c.BindJSON(&rating); err != nil {
		c.JSON(400, gin.H{"error": "invalid rating value"})
		return
	}

	if rating < 1 || rating > 5 {
		c.JSON(400, gin.H{"error": "rating value must be between 1 and 5"})
		return
	}

	order.Rating = rating
	models.DB.Save(&order)

	c.JSON(200, gin.H{"success": "rating updated successfully"})
}
