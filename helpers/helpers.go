package helpers

import (
	b64 "encoding/base64"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/AvestaBarzegar/statify-api/helpers/consts"
)

func encodeClientSecretAndId() string {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	toBeEncoded := clientId + ":" + clientSecret
	return b64.StdEncoding.EncodeToString([]byte(toBeEncoded))
}

func ExchangeCodeForToken(grantType string, redirectUri string, code string) (*http.Response, error) {
	authToken := "Basic " + encodeClientSecretAndId()
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
