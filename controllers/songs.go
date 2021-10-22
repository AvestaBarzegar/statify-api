package controllers

import (
	"github.com/AvestaBarzegar/statify-api/helpers"
	"github.com/gin-gonic/gin"
)

/*
 GET v1/song/{id}/analysis
 headers: {
	Authorization: `Bearer "Token`
	Accept: `application/json`
	Content-Type: `applicaton/json`
}

returns:
{
	// The bit below is acquired from calling GET https://api.spotify.com/v1/audio-analysis?id={id}&artist={artist}&title={title}
	track audio analysis: {
		"danceability": 0.847,
		"energy": 0.529,
		"key": 5,
		"loudness": -5.321,
		"mode": 0,
		"speechiness": 0.175,
		"acousticness": 0.481,
		"instrumentalness": 0.00348,
		"liveness": 0.0952,
		"valence": 0.797,
		"tempo": 92.685,
	},
	// The bit below is acquired from calling GET https://api.lyrics.ovh/v1/{artist}/{title}
	lyrics: "string"
}


*/

// GET v1/song/{id}/analysis

func GetLyrics(c *gin.Context) {

	res, err := helpers.GetLyrics("Frank Ocean", "Nights")

	if err != nil {
		c.JSON(res.StatusCode, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/json", res.Body, nil)
}
