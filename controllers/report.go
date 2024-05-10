package controllers

import (
	"goproj/models"
	"goproj/utils"

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
	var requestBody struct {
		TrackCode   string `json:"track_code"`
		Rating      int    `json:"rating"`
		Description string `json:"description"`
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

	var user models.User
	models.DB.First(&user, claims.Id)

	if order.Status != "Finished" {
		c.JSON(400, gin.H{"error": "order status must be finished to update rating"})
		return
	}

	if order.UserID != user.ID {
		c.JSON(403, gin.H{"error": "you are not allowed to update the rating for this order"})
		return
	}

	var rating = requestBody.Rating

	if rating < 1 || rating > 5 {
		c.JSON(400, gin.H{"error": "rating value must be between 1 and 5"})
		return
	}

	order.Rating = rating
	order.Description = requestBody.Description
	models.DB.Save(&order)

	c.JSON(200, gin.H{"success": "rating updated successfully"})
}
