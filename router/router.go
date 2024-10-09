package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// Initialize the router
	r := gin.Default()

	// Define the routes
	initRoutes(r)

	// Start the server
	r.Run(":8080")
}
