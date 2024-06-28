package router

import (
	"data-manage/commons/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRouter is used to get the router configured with the middlewares and the routes
func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	v1Routes := router.Group("v1")
	{
		// Health Check
		v1Routes.GET(constants.HealthCheck, func(c *gin.Context) {
			response := map[string]string{
				"message": "API is up and running",
			}
			c.JSON(http.StatusOK, response)
		})
		v1Routes.POST(constants.HealthCheck, getBrokerRecommendationsController.HandleGetBrokerRecommendations)

	}
	return router
}
