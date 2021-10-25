package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/AvestaBarzegar/statify-api/helpers/consts"
	jsonModels "github.com/AvestaBarzegar/statify-api/models/http"
	"github.com/gin-gonic/gin"
)

// read this
const grantType = "client_credentials"

var ServerAuthToken *string = nil
var expiryDate = time.Now().Unix() - 3000

func RefreshAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentTime := time.Now().Unix()
		// If the authToken is not expired then we need to do nothing
		if currentTime <= expiryDate || ServerAuthToken != nil {
			return
		}
		// Make HTTP request and update expiryDate along with the AuthToken
		resp, err := getServerAuthToken()
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var result jsonModels.AuthTokenResponse
		if err := json.Unmarshal(body, &result); err != nil {
			c.AbortWithStatus(http.StatusFailedDependency)
			return
		}
		expiryDate = time.Now().Unix() + int64(result.ExpiresIn) - 1800
		ServerAuthToken = &result.AccessToken
	}
}

func getServerAuthToken() (*http.Response, error) {
	authToken := "Basic " + consts.EncodeClientSecretAndId()
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
