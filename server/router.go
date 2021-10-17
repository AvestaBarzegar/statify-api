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
	authRouter.GET("/authorize", controllers.LoginUser)
	router.Run(":" + PORT)
	return router
}
