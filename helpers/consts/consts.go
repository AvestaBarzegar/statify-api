package consts

import (
	b64 "encoding/base64"
	"os"
)

const SpotifyTokenURL = "https://accounts.spotify.com/api/token"
const LyricsURL = "https://api.lyrics.ovh/v1/"
const AudioFeaturesURL = "https://api.spotify.com/v1/audio-features"

var ClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
var ClientId = os.Getenv("SPOTIFY_CLIENT_ID")

func EncodeClientSecretAndId() string {
	toBeEncoded := ClientId + ":" + ClientSecret
	return b64.StdEncoding.EncodeToString([]byte(toBeEncoded))
}
