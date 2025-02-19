package list

import (
	"errors"
	"go-auth/initializers"
	"go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetList(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }

	var err error

	workspaceIDRaw := c.Param("id")

	workspaceID, parseErr := uuid.Parse(workspaceIDRaw)

	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID!", "data": nil})
		return
	}

	var lists []models.List
	if err = initializers.DB.Model(&models.List{}).Preload("Cards", func(DB *gorm.DB) *gorm.DB {
		return DB.Order("cards.order ASC")
	}).Where("workspace_id", workspaceID).Order("lists.order ASC").Find(&lists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get list data!", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List retrieved successfully.", "data": lists})
}

func InsertList(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }

	var input struct {
		Name        string `json:"name" binding:"required"`
		WorkspaceID string `json:"workspace_id" binding:"required"`
		Order       int    `json:"order" binding:"required"`
	}

	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data!"})
		return
	}

	var parsedWorkspaceID uuid.UUID

	parsedWorkspaceID, err = uuid.Parse(input.WorkspaceID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid workspace ID!", "data": nil})
		return
	}

	list := &models.List{Name: input.Name, WorkspaceID: parsedWorkspaceID, Order: input.Order}
	if err = initializers.DB.Create(&list).Preload("Cards").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!", "data": nil})
		return
	}

	if list.Cards == nil {
		list.Cards = []models.Card{}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "New list created successfully",
		"data":    list,
	})
}

func UpdateList(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }
	var err error

	listIDRaw := c.Param("id")

	var listID uuid.UUID

	listID, err = uuid.Parse(listIDRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid list ID!", "data": nil})
		return
	}

	var input struct {
		Name  string `json:"name"`
		Order int    `json:"order"`
	}
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data!"})
		return
	}

	list := &models.List{}
	if err = initializers.DB.Model(&models.List{}).Where("id", listID).Preload("Cards").First(&list).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "List not found!", "data": nil})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!", "data": nil})
		return
	}

	if list.Cards == nil {
		list.Cards = []models.Card{}
	}

	if input.Name != "" {
		list.Name = input.Name
	}

	if input.Order > 0 {
		list.Order = input.Order
	}

	if err = initializers.DB.Save(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "List updated successfully",
		"data":    list,
	})
}

func DeleteList(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	// userID, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }
	var err error

	listIDRaw := c.Param("id")

	var listID uuid.UUID

	listID, err = uuid.Parse(listIDRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid list ID!", "data": nil})
		return
	}

	if err = initializers.DB.Where("id = ?", listID).Delete(&models.List{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "List deleted successfully",
		"data":    nil,
	})
}
