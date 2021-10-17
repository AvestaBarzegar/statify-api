package server

import (
	"github.com/AvestaBarzegar/statify-api/controllers"
	"github.com/AvestaBarzegar/statify-api/helpers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	PORT := helpers.GoDotEnvVar("PORT")
	authRouter := router.Group("/v1/api/account/")
	authRouter.GET("/login", controllers.AuthenticateUser)
	router.Run(":" + PORT)
	return router
}
