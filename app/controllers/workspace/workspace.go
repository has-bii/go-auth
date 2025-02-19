package workspace

import (
	"errors"
	"go-auth/initializers"
	"go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkspaceDTOWithMembers struct {
	models.WorkspaceDTO
	Members []models.UserDTO `json:"members"`
}

type result struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	OwnerID uuid.UUID `json:"owner_id"`
}

func GetWorkspace(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	results := []result{}
	if err := initializers.DB.Table("workspaces").
		Select("workspaces.id, workspaces.name, workspaces.owner_id").
		Joins("left join workspace_users wu on workspaces.id = wu.workspace_id").
		Where("workspaces.owner_id = ? OR wu.user_id = ?", userID, userID).
		Scan(&results).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get workspace data!", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workspace data retrieved.", "data": results})
}

func GetWorkspaceByID(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	workspaceID := c.Param("id")

	var workspace models.WorkspaceDTO
	if err := initializers.DB.Model(&models.Workspace{}).Where("id", workspaceID).Preload("Owner", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "email")
	}).First(&workspace).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Workspace not found!", "data": nil})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get workspace data!", "data": nil})
		return
	}

	workspaceWithMembers := WorkspaceDTOWithMembers{
		WorkspaceDTO: workspace,
		Members:      []models.UserDTO{},
	}

	if err := initializers.DB.Raw(`SELECT u.id, u.name, u.email FROM workspace_users w INNER JOIN users u ON w.user_id = u.id WHERE w.workspace_id = ?`, workspaceID).
		Scan(&workspaceWithMembers.Members).
		Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could not find members of this workspace!", "data": nil})
	}

	isRelated := workspaceWithMembers.OwnerID == userID

	for i := 0; i < len(workspaceWithMembers.Members); i++ {
		if userID == workspaceWithMembers.Members[i].ID {
			isRelated = true
		}
	}

	if !isRelated {
		c.JSON(http.StatusForbidden, gin.H{"message": "You are not either the owner or a member of this workspace!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workspace data retrieved.", "data": workspaceWithMembers})
}

func InsertWorkspace(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Extract user ID from context (set in middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	workspace := models.Workspace{Name: input.Name, OwnerID: userID.(uuid.UUID)}

	result := initializers.DB.Create(&workspace)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "New workspace created successfully",
		"data":    nil,
	})
}

func UpdateWorkspace(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "data is required!"})
		return
	}

	// Extract user ID from context (set in middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Get id from param
	workspaceID := c.Param("id")

	var workspace models.Workspace
	if err := initializers.DB.
		Model(models.Workspace{}).
		Where("id = ?", workspaceID).
		First(&workspace).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Workspace not found!"})
			return

		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
		return
	}

	// check owner
	if userID != workspace.OwnerID {
		c.JSON(http.StatusForbidden, gin.H{"message": "You are not the owner of this workspace!"})
		return
	}

	workspace.Name = input.Name
	initializers.DB.Save(workspace)

	c.JSON(http.StatusOK, gin.H{"message": "Workspace updated successfully."})
}

func DeleteWorkspace(c *gin.Context) {
	// Extract user ID from context (set in middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Get id from param
	workspaceID := c.Param("id")

	var workspace result
	if err := initializers.DB.Table("workspaces").
		Where("id = ?", workspaceID).
		First(&workspace).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Workspace not found!"})
			return

		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
		return
	}

	// check owner
	if userID != workspace.OwnerID {
		c.JSON(http.StatusForbidden, gin.H{"message": "You are not the owner of this workspace!"})
		return
	}

	if errorDelete := initializers.DB.Where("id = ?", workspaceID).Delete(&models.Workspace{}).Error; errorDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workspace deleted successfully."})
}

func AddMember(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required"`
	}

	var errs error

	if errs = c.ShouldBindJSON(&input); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data!"})
		return
	}

	// Extract user ID from context (set in middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Get id from param
	workspaceIDString := c.Param("id")

	workspaceID, err := uuid.Parse(workspaceIDString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid workspace ID"})
		return
	}

	// Get workspace
	var workspace models.Workspace
	if errs = initializers.DB.
		Model(models.Workspace{}).
		Where("id = ?", workspaceID).
		First(&workspace).
		Error; errs != nil {
		if errors.Is(errs, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Workspace not found!"})
			return

		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
		return
	}

	// check owner
	if userID != workspace.OwnerID {
		c.JSON(http.StatusForbidden, gin.H{"message": "You are not the owner of this workspace!"})
		return
	}

	// Get user
	var user models.User
	if errs = initializers.DB.Model(&models.User{}).Where("email = ?", input.Email).First(&user).Error; errs != nil {
		if errors.Is(errs, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Workspace not found!"})
			return

		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
		return
	}

	// Inserting member
	if errs := initializers.DB.Table("workspace_users").Create(models.WorkspaceUser{WorkspaceID: workspaceID, UserID: user.ID}).Error; errs != nil {
		if errors.Is(errs, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Mamber has already joined!"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "New member added successfully."})
}
