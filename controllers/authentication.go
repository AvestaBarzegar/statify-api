package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AvestaBarzegar/statify-api/bindings"
)

// This file handles the base path of /v1/api/account

// POST /v1/api/token
func ProvideAccessToken(c *gin.Context) {
	body := bindings.RequestAccessTokenBody{}

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Code == "" || body.GrantType == "" || body.RedirectURI == "" {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	c.String(200, "ok\n")
}
