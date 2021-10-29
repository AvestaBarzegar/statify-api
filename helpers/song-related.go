package helpers

import (
	"net/http"

	"github.com/AvestaBarzegar/statify-api/helpers/consts"
	"github.com/AvestaBarzegar/statify-api/middleware"
)

func GetLyrics(artist string, track string) (*http.Response, error) {
	newPath := consts.LyricsURL + artist + "/" + track
	res, err := http.Get(newPath)
	return res, err
}

func GetSongDetails(id string) (*http.Response, error) {
	newPath := consts.AudioFeaturesURL + id + "/"
	authHeader := "Bearer " + *middleware.ServerAuthToken
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodGet, newPath, nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", authHeader)
	res, err := client.Do(r)
	return res, err
}
