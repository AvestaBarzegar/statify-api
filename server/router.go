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
	setupSongRouter(router)
	router.Run(":" + PORT)
	return router
}

func setupAuthRouter(router *gin.Engine) {
	authRouter := router.Group("/v1/")
	authRouter.POST("/v1/api/token", controllers.ProvideAccessToken)
}

func setupSongRouter(router *gin.Engine) {
	songRouter := router.Group("/v1/songs")
	songRouter.Use(middleware.RefreshAuthToken())
	songRouter.GET("/", controllers.GetLyrics)
}
