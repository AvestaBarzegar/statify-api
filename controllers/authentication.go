package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AvestaBarzegar/statify-api/bindings"
	"github.com/AvestaBarzegar/statify-api/helpers"
)

// This file handles the base path of /v1/api/account

// POST /v1/api/token
func ProvideAccessToken(c *gin.Context) {
	body := bindings.RequestAccessTokenBody{}

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	if body.Code == "" || body.GrantType == "" || body.RedirectURI == "" {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	res, e := helpers.ExchangeCodeForToken(body.GrantType, body.RedirectURI, body.Code)

	if e != nil {
		c.JSON(res.StatusCode, gin.H{"error": e.Error()})
		return
	}

	c.JSON(200, res)
}
