package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the routes for the Gin server.
func RegisterRoutes(server *gin.Engine) {
	// Route for inserting a new agency
	server.POST("/agency", insertAgency)

	// Route for updating an existing agency by ID
	server.PUT("/agency/:agencyId", updateAgency)

	// Route for getting an agency by ID
	server.GET("/agency/:agencyId", getAgencyById)
}
