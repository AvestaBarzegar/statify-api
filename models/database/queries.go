package database

func GetHistoricalTopArtists(spotify_user_id string) ([]TopArtists, error) {
	stmt := `SELECT (created_at, time_span, items) FROM artists WHERE spotify_user_id=$1`
	db.
}