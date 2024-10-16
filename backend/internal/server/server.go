package server

import (
	"github.com/briantkho/bling/internal/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewServer() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())
	setupRoutes(router)

	return router
}

/*
Define routes here
*/
func setupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "we are so live rn") })
	userRoutes := router.Group("/api/user")
	{
		userRoutes.GET("/:id", user.GetUser)
	}
}
