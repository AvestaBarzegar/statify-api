package server

import (
	"os"

	"github.com/AvestaBarzegar/statify-api/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	PORT := os.Getenv("PORT")
	authRouter := router.Group("/")
	authRouter.POST("/v1/api/token", controllers.ProvideAccessToken)
	router.Run(":" + PORT)
	return router
}
