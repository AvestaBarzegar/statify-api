package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AvestaBarzegar/statify-api/helpers"
	jsonModels "github.com/AvestaBarzegar/statify-api/models/http"
	"github.com/gin-gonic/gin"
)

// GET v1/track/traits
func GetTrackTraits(c *gin.Context) {
	queryParams := jsonModels.SongInformationQuery{}

	err := c.ShouldBind(&queryParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	if queryParams.Artist == "" || queryParams.SongId == "" || queryParams.Track == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res1, err := helpers.GetLyrics(queryParams.Artist, queryParams.Track)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer res1.Body.Close()
	body1, _ := ioutil.ReadAll(res1.Body)
	var lyrics jsonModels.LyricsResponseModel
	if err := json.Unmarshal(body1, &lyrics); err != nil {
		c.AbortWithStatus(http.StatusFailedDependency)
		return
	}

	res2, err := helpers.GetSongDetails(queryParams.SongId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer res2.Body.Close()
	body2, _ := ioutil.ReadAll(res2.Body)
	var traits jsonModels.TrackTraitsResponse
	if err := json.Unmarshal(body2, &traits); err != nil {
		c.AbortWithStatus(http.StatusFailedDependency)
		return
	}
	var trackTraits jsonModels.TrackTraitsClientResponseModel
	trackTraits.Lyrics = lyrics.Lyrics
	trackTraits.TrackTraits = traits
	c.JSON(http.StatusAccepted, trackTraits)
}
