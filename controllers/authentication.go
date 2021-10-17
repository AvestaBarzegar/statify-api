package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AvestaBarzegar/statify-api/helpers"
)

// GET /v1/api/account/login
func AuthenticateUser(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, helpers.RedirectURL())
}
