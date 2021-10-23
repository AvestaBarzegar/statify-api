package middleware

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/AvestaBarzegar/statify-api/helpers"
	"github.com/AvestaBarzegar/statify-api/helpers/consts"
)

const grantType = "client_credentials"

var ServerAuthToken *string = nil
var expiryDate = time.Now().Unix() - 3000

func RefreshAuthToken() {
	currentTime := time.Now().Unix()
	// If the authToken is expired
	if currentTime >= expiryDate || ServerAuthToken == nil {
		// Make HTTP request and update expiryDate
	}
}

func getServerAuthToken() (*http.Response, error) {
	authToken := "Basic " + helpers.EncodeClientSecretAndId()
	body := url.Values{}
	body.Set("grant_type", grantType)
	encodedBody := body.Encode()
	client := &http.Client{}

	r, _ := http.NewRequest(http.MethodPost, consts.SpotifyTokenURL, strings.NewReader(encodedBody))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", authToken)
	res, err := client.Do(r)
	return res, err
}