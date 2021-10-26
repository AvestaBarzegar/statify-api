package consts

import (
	b64 "encoding/base64"
	"os"
	"strconv"
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
var USER = os.Getenv("POSTGRES_USER")
var PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var DB = os.Getenv("POSTGRES_DB")

func EncodeClientSecretAndId() string {
	toBeEncoded := ClientId + ":" + ClientSecret
	return b64.StdEncoding.EncodeToString([]byte(toBeEncoded))
}

func GetPortDB() int64 {
	port, err := strconv.ParseInt(os.Getenv("DB_PORT"), 0, 64)
	if err != nil {
		panic(err)
	}
	return port
}
