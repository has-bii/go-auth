package card

import (
	"errors"
	"go-auth/initializers"
	"go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InsertCard(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }

	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		ListIDRaw   string `json:"list_id" binding:"required"`
		Order       int    `json:"order" binding:"required"`
	}

	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data!"})
		return
	}

	var listID uuid.UUID

	listID, err = uuid.Parse(input.ListIDRaw)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid workspace ID!", "data": nil})
		return
	}

	card := &models.Card{ListID: listID, Name: input.Name, Order: input.Order, Description: input.Description}
	if err = initializers.DB.Create(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "New card created successfully",
		"data":    card,
	})
}

func UpdateCard(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }
	var err error

	cardIDRaw := c.Param("id")

	var cardID uuid.UUID

	cardID, err = uuid.Parse(cardIDRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid card ID!", "data": nil})
		return
	}

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Order       int    `json:"order"`
	}

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data!"})
		return
	}

	var card models.Card
	if err = initializers.DB.Model(&models.Card{}).Where("id", cardID).First(&card).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Card not found!", "data": nil})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!", "data": nil})
		return
	}

	if input.Name != "" {
		card.Name = input.Name
	}

	if input.Description != "" {
		card.Description = input.Description
	}

	if input.Order > 0 {
		card.Order = input.Order
	}

	if err = initializers.DB.Save(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Card updated successfully",
		"data":    card,
	})
}

func DeleteCard(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }
	var err error

	cardIDRaw := c.Param("id")

	var cardID uuid.UUID

	cardID, err = uuid.Parse(cardIDRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid card ID!", "data": nil})
		return
	}

	if err = initializers.DB.Where("id = ?", cardID).Delete(&models.Card{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Card deleted successfully",
		"data":    nil,
	})
}
