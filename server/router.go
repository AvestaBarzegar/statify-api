package server

import (
	"os"

	"github.com/AvestaBarzegar/statify-api/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	PORT := os.Getenv("PORT")
	authRouter := router.Group("/v1/api/account")
	authRouter.GET("/login", controllers.AuthenticateUser)
	router.Run(":" + PORT)
	return router
}
