package controllers

import (
	"net/http"
	"net/url"

	"github.com/AvestaBarzegar/statify-api/helpers"
	"github.com/gin-gonic/gin"
)

var SPOTIFY_CLIENT_ID = helpers.GoDotEnvVar("SPOTIFY_CLIENT_ID")
var SPOTIFY_REDIRECT_URI = helpers.GoDotEnvVar("SPOTIFY_REDIRECT_URI")

func redirectURL() string {
	scope := "user-top-read,user-read-recently-played,user-read-private,user-read-email"
	params := url.Values{}
	params.Add("client_id", SPOTIFY_CLIENT_ID)
	params.Add("response_type", "code")
	params.Add("redirect_uri", SPOTIFY_REDIRECT_URI)
	params.Add("scope", scope)
	params.Add("show_dialog", "true")

	query := params.Encode()
	url := url.URL{
		Scheme:   "https",
		Host:     "accounts.spotify.com",
		Path:     "authorize",
		RawQuery: query,
	}

	return url.String()
}

// GET /v1/api/account/login
func AuthenticateUser(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, redirectURL())
}
