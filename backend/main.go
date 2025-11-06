package main

import (
	"time"
	"gocommerce/config"
	"gocommerce/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// ✅ Allow CORS (frontend: localhost:3000)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ✅ Connect to MongoDB
	config.ConnectDB()

	// ✅ Serve static files from the "uploads" folder (for images)
	r.Static("/uploads", "./uploads")

	// ✅ Register all routes (including /upload)
	routes.RegisterRoutes(r)

	// ✅ Start the server
	r.Run(":8080")
}
