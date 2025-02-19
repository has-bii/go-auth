package main

import (
	"go-auth/controllers"
	"go-auth/controllers/card"
	"go-auth/controllers/list"
	"go-auth/controllers/workspace"
	"go-auth/initializers"
	"go-auth/middleware"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	// Get the frontend URL from the environment variable
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		log.Fatal("FRONTEND_URL is not set in the .env file")
	}

	// Create Gin instance
	r := gin.Default()

	// Set CORS policy
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL}, // Use frontend URL from environment variable
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache preflight request
	}))

	// Helo world
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	auth := r.Group("/auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/auth", controllers.GetProfile)

	workspaceProtected := r.Group("/workspace")
	workspaceProtected.Use(middleware.AuthMiddleware())

	workspaceProtected.GET("/", workspace.GetWorkspace)
	workspaceProtected.GET("/:id", workspace.GetWorkspaceByID)
	workspaceProtected.POST("/", workspace.InsertWorkspace)
	workspaceProtected.PUT("/:id", workspace.UpdateWorkspace)
	workspaceProtected.DELETE("/:id", workspace.DeleteWorkspace)
	workspaceProtected.POST("/add-member/:id", workspace.AddMember)

	listProtected := r.Group("/list")
	listProtected.Use(middleware.AuthMiddleware())

	listProtected.GET("/:id", list.GetList)
	listProtected.POST("/", list.InsertList)
	listProtected.PUT("/:id", list.UpdateList)
	listProtected.DELETE("/:id", list.DeleteList)

	cardProtected := r.Group("/card")
	cardProtected.Use(middleware.AuthMiddleware())

	cardProtected.POST("", card.InsertCard)
	cardProtected.PUT("/:id", card.UpdateCard)
	cardProtected.DELETE("/:id", card.UpdateCard)

	port := os.Getenv("PORT")
	r.Run(`:` + port)
}
