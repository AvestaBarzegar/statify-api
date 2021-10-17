package helpers

import (
	"net/url"
	"os"
)

var SPOTIFY_CLIENT_ID = os.Getenv("SPOTIFY_CLIENT_ID")
var SPOTIFY_REDIRECT_URI = os.Getenv("SPOTIFY_REDIRECT_URI")

func RedirectURL() string {
	scope := "user-top-read,user-read-recently-played,user-read-private,user-read-email"
	params := url.Values{}
	params.Add("client_id", SPOTIFY_CLIENT_ID)
	params.Add("response_type", "code")
	params.Add("redirect_uri", SPOTIFY_REDIRECT_URI)
	params.Add("scope", scope)
	params.Add("show_dialog", "true")

	query := params.Encode()
	url := url.URL{
		Scheme:   "https",
		Host:     "accounts.spotify.com",
		Path:     "authorize",
		RawQuery: query,
	}

	return url.String()
}
