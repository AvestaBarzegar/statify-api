package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/AvestaBarzegar/statify-api/models/database"
)

// POST /v1/api/token
func ProvideAccessToken(c *gin.Context) {
	tracks := []database.Track{
		database.Track{
			TrackName:      "Blonde",
			TrackImageURL:  "https://sugma",
			TrackRank:      1,
			SpotifyTrackId: "skajdlkajsd",
			ArtistName:     "Frank Ocean",
		},
		database.Track{
			TrackName:      "Blonde",
			TrackImageURL:  "https://sugma",
			TrackRank:      2,
			SpotifyTrackId: "skajdlkajsd",
			ArtistName:     "PNG Ocean",
		},
		database.Track{
			TrackName:      "Blonde",
			TrackImageURL:  "https://sugma",
			TrackRank:      3,
			SpotifyTrackId: "skajdlkajsd",
			ArtistName:     "JPEG Ocean",
		},
	}

	database.InsertTrackRow("alrighticantfindagoodname", database.FourWeeks, tracks)
	c.String(200, "testing")
	// body := httpModels.RequestAccessTokenQuery{}

	// err := c.ShouldBind(&body)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	log.Println(err)
	// 	return
	// }

	// if body.Code == "" || body.GrantType != "authorization_code" || body.RedirectURI == consts.SpotifyTokenURL {
	// 	c.String(http.StatusBadRequest, "Bad Request")
	// 	return
	// }
	// res, err := helpers.ExchangeCodeForToken(body.GrantType, body.RedirectURI, body.Code)

	// if err != nil {
	// 	c.JSON(res.StatusCode, gin.H{"error": err.Error()})
	// 	return
	// }
	// defer res.Body.Close()
	// c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}
