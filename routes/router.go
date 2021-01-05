package routes

import (
	"github.com/SteakBarbare/Partiel-MT4/controllers"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes set up the routes for the server
func InitializeRoutes(r *gin.Engine) {
	// HTML routes for the GUI
	r.GET("/", controllers.LoginPage)
	r.GET("/login", controllers.LoginPage)

	// // API routes
	api := r.Group("/api")
	{
		api.POST("/login", controllers.SooS)
	}

	// Assets routes
	r.Static("views/assets", "./views/assets")
}
