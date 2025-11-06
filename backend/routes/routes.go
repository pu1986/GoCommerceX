package routes

import (
	"gocommerce/controllers"
	"gocommerce/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", controllers.Home)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	r.PUT("/make-admin/:email", controllers.MakeAdmin)

	// ✅ Image upload route
	r.POST("/upload", controllers.UploadImage)

	// ✅ Admin routes (protected)
	admin := r.Group("/admin")
	admin.Use(middleware.AdminMiddleware())
	{
		admin.POST("/products", controllers.CreateProduct)
		admin.PUT("/products/:id", controllers.UpdateProduct)
		admin.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
