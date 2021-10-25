package helpers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/AvestaBarzegar/statify-api/helpers/consts"
	"github.com/AvestaBarzegar/statify-api/middleware"
)

func ExchangeCodeForToken(grantType string, redirectUri string, code string) (*http.Response, error) {
	authToken := "Basic " + consts.EncodeClientSecretAndId()
	body := url.Values{}
	body.Set("grant_type", grantType)
	body.Set("redirect_uri", redirectUri)
	body.Set("code", code)
	encodedBody := body.Encode()
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodPost, consts.SpotifyTokenURL, strings.NewReader(encodedBody))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", authToken)
	res, err := client.Do(r)
	return res, err
}

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
