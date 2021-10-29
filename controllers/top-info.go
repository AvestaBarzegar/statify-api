package controllers

import (
	"net/http"

	httpModels "github.com/AvestaBarzegar/statify-api/models/http"
	"github.com/gin-gonic/gin"
)

// GETS TOP TRACKS
// GET /v1/me/top/tracks
// QUERYPARAM: spotify_user_id="alrighticantfindagoodname"
// HEADERS: [Content-Type: application/json], [Authorization: `Bearer {token}`]
func GetTopTracks(c *gin.Context) {
	h := httpModels.RequestHeaders{}
	err := c.ShouldBindHeader(&h)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	auth := h.Authorization
	// Make request to get spotify top tracks for all 3 timespans
	// Get the data from the database
	// After making such a request, check in the database to make sure an insert is allowed

}
