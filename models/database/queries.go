package database

import (
	"log"
)

func GetHistoricalTopArtists(spotify_user_id string) ([]TopArtists, error) {
	db, e := Initialize()
	var topArtists []TopArtists
	if e != nil {
		log.Printf("Something went wrong %v", e)
		return topArtists, e
	}
	conn := db.Conn
	stmtStr := `SELECT created_at, time_span, items FROM artists WHERE spotify_user_id=$1`
	stmt, err := conn.Prepare(stmtStr)
	if err != nil {
		log.Print(err)
	}
	rows, errStmt := stmt.Query(spotify_user_id)
	if errStmt != nil {
		log.Printf("Something went wrong %v", errStmt)
	}
	defer rows.Close()
	for rows.Next() {
		topArtist := new(TopArtists)
		err := rows.Scan(&topArtist.CreatedAt, &topArtist.TimeSpan, &topArtist.Artists)
		if err != nil {
			log.Printf("Something went wrong %v", err)
		}
		topArtists = append(topArtists, *topArtist)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Something went wrong %v", err)
	}
	log.Println(topArtists)
	return topArtists, nil
}

func GetHistoricalTopTracks(spotify_user_id string) ([]TopTracks, error) {
	db, e := Initialize()
	var topTracks []TopTracks
	if e != nil {
		log.Printf("Something went wrong %v", e)
		return topTracks, e
	}
	conn := db.Conn
	stmtStr := `SELECT created_at, time_span, items FROM tracks WHERE spotify_user_id=$1`
	stmt, err := conn.Prepare(stmtStr)
	if err != nil {
		log.Print(err)
	}
	rows, errStmt := stmt.Query(spotify_user_id)
	if errStmt != nil {
		log.Printf("Something went wrong %v", errStmt)
	}
	defer rows.Close()
	for rows.Next() {
		topTrack := new(TopTracks)
		err := rows.Scan(&topTrack.CreatedAt, &topTrack.TimeSpan, &topTrack.Tracks)
		if err != nil {
			log.Printf("Something went wrong %v", err)
		}
		topTracks = append(topTracks, *topTrack)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Something went wrong %v", err)
	}
	log.Println(topTracks)
	return topTracks, nil
}
