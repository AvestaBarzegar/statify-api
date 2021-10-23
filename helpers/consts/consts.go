package consts

import "os"

const SpotifyTokenURL = "https://accounts.spotify.com/api/token"
const LyricsURL = "https://api.lyrics.ovh/v1/"
const AudioFeaturesURL = "https://api.spotify.com/v1/audio-features"

var ClientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
var ClientId = os.Getenv("SPOTIFY_CLIENT_ID")
