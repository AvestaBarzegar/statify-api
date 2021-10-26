package consts

import (
	b64 "encoding/base64"
	"os"
)

// Common URLS TODO: Refactor and made methods that build URLs
const SpotifyTokenURL = "https://accounts.spotify.com/api/token"
const LyricsURL = "https://api.lyrics.ovh/v1/"
const AudioFeaturesURL = "https://api.spotify.com/v1/audio-features/"

// Spotify Client Information
var ClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
var ClientId = os.Getenv("SPOTIFY_CLIENT_ID")

// DB Configuration information
var HOST = os.Getenv("DB_HOST")
var PORT = os.Getenv("DB_PORT")
var USER = os.Getenv("POSTGRES_USER")
var PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var DB = os.Getenv("POSTGRES_DB")

func EncodeClientSecretAndId() string {
	toBeEncoded := ClientId + ":" + ClientSecret
	return b64.StdEncoding.EncodeToString([]byte(toBeEncoded))
}
