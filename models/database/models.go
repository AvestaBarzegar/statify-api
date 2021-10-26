package database

const (
	FourWeeks = "four weeks"
	SixMonths = "six months"
	AllTime   = "all time"
)

type Track struct {
	TrackName      string `json:"track_name" db:"track_name`
	ArtistName     string `json:"artist_name" db:"artist_name"`
	SpotifyTrackId string `json:"spotify_track_id" db:"spotify_track_id`
	TrackImageURL  string `json:"track_image_url" db:"track_image_url"`
	TrackRank      int    `json:"track_rank" db:"track_rank"`
}

type Artist struct {
	ArtistName      string `json:"artist_name"`
	SpotifyArtistId string `json:"spotify_track_id"`
	ArtistImageURL  string `json:"track_image_url"`
	ArtistRank      int    `json:"track_rank"`
}

type TopTracks struct {
	SpotifyUserId string `json:"spotify_user_id"`
	CreatedAt     int64  `json:"created_at"`
	// Is either FourWeeks, SixMonths, or AllTime
	TimeSpan string `json:"time_span"`
	Tracks   Track  `json:"items"`
}

type TopArtists struct {
	SpotifyUserId string `json:"spotify_user_id"`
	CreatedAt     int64  `json:"created_at"`
	// Is either FourWeeks, SixMonths, or AllTime
	TimeSpan string `json:"time_span"`
	Artists  Artist `json:"items"`
}

type TopInfoResponse struct {
	TopTracks  []TopTracks  `json:"top_tracks"`
	TopArtists []TopArtists `json:"top_artists"`
}
