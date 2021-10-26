package server

import (
	"os"

	"github.com/AvestaBarzegar/statify-api/controllers"
	"github.com/AvestaBarzegar/statify-api/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	PORT := os.Getenv("PORT")
	setupAuthRouter(router)
	setupTrackRouter(router)
	router.Run(":" + PORT)
	return router
}

func setupAuthRouter(router *gin.Engine) {
	authRouter := router.Group("/v1")
	authRouter.POST("/api/token", controllers.ProvideAccessToken)
}

func setupTrackRouter(router *gin.Engine) {
	trackRouter := router.Group("/v1/tracks")
	trackRouter.Use(middleware.RefreshAuthToken())
	trackRouter.GET("/traits/", controllers.GetTrackTraits)
}
