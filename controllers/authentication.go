package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AvestaBarzegar/statify-api/bindings"
	"github.com/AvestaBarzegar/statify-api/helpers"
	"github.com/AvestaBarzegar/statify-api/helpers/consts"
)

// POST /v1/api/token
func ProvideAccessToken(c *gin.Context) {
	body := bindings.RequestAccessTokenQuery{}

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	if body.Code == "" || body.GrantType != "authorization_code" || body.RedirectURI == consts.SpotifyTokenURL {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	res, err := helpers.ExchangeCodeForToken(body.GrantType, body.RedirectURI, body.Code)

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}
