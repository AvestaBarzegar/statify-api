package helpers

import (
	b64 "encoding/base64"
	"net/http"
	"net/url"
	"os"
	"strings"
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

	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodPost, spotifyTokenURL, strings.NewReader(body.Encode()))
	r.Header.Add("Authorization", authToken)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return client.Do(r)
}

var spotifyTokenURL = "https://accounts.spotify.com/api/token"
