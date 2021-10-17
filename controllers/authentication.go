package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AvestaBarzegar/statify-api/helpers"
)

// This file handles the base path of /v1/api/account

// GET /authorize
func LoginUser(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, helpers.RedirectURL())
}

// POST /v1/api/token
func ProvideAccessToken(c *gin.Context) {

}
