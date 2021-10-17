package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AvestaBarzegar/statify-api/helpers"
)

// This file handles the base path of /v1/api/account

// GET /v1/api/account/login
func LoginUser(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, helpers.RedirectURL())
}

// GET /v1/api/account/token
func AuthenticateUserWithToken(c *gin.Context) {

}
