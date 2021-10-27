package database

import (
	"log"

	"github.com/lib/pq"
)

// Note that time_span is either FourWeeks, SixMonths, or AllTime
func InsertTrackRow(spotifyUserId string, time_span string, tracks []Track) bool {
	db, e := Initialize()
	if e != nil {
		log.Fatalf("Something went wrong %v", e)
		return false
	}
	conn := db.Conn
	stmt := `INSERT INTO tracks 
			(spotify_user_id, time_span, items) 
			VALUES 
			($1, $2, $3::track[]);`
	_, err := conn.Exec(stmt, spotifyUserId, time_span, pq.Array(tracks))

	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func InsertArtistRow(spotifyUserId string, time_span string, artists []Artist) bool {
	db, e := Initialize()
	if e != nil {
		log.Fatalf("Something went wrong %v", e)
		return false
	}
	conn := db.Conn
	stmt := `INSERT INTO artists 
			(spotify_user_id, time_span, items) 
			VALUES 
			($1, $2, $3::artist[]);`
	_, err := conn.Exec(stmt, spotifyUserId, time_span, pq.Array(artists))

	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
