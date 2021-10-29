package helpers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/AvestaBarzegar/statify-api/helpers/consts"
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
