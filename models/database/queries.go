package database

import (
	"log"
	"time"
)

func GetHistoricalTopArtists(spotifyUserId string) ([]TopArtists, error) {
	db, e := Initialize()
	var topArtists []TopArtists
	if e != nil {
		log.Printf("Something went wrong %v", e)
		return topArtists, e
	}
	conn := db.Conn
	stmtStr := `SELECT (EXTRACT(EPOCH FROM created_at))::bigint, time_span, items FROM artists WHERE spotify_user_id=$1`
	stmt, err := conn.Prepare(stmtStr)
	if err != nil {
		log.Print(err)
	}
	rows, errStmt := stmt.Query(spotifyUserId)
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

func GetHistoricalTopTracks(spotifyUserId string) ([]TopTracks, error) {
	db, e := Initialize()
	var topTracks []TopTracks
	if e != nil {
		log.Printf("Something went wrong %v", e)
		return topTracks, e
	}
	conn := db.Conn
	stmtStr := `SELECT (EXTRACT(EPOCH FROM created_at))::bigint, time_span, items FROM tracks WHERE spotify_user_id=$1`
	stmt, err := conn.Prepare(stmtStr)
	if err != nil {
		log.Print(err)
		return topTracks, err
	}
	rows, errStmt := stmt.Query(spotifyUserId)
	if errStmt != nil {
		log.Printf("Something went wrong %v", errStmt)
		return topTracks, errStmt
	}
	defer rows.Close()
	for rows.Next() {
		topTrack := new(TopTracks)
		err := rows.Scan(&topTrack.CreatedAt, &topTrack.TimeSpan, &topTrack.Tracks)
		if err != nil {
			log.Printf("Something went wrong %v", err)
			return topTracks, err
		}
		topTracks = append(topTracks, *topTrack)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Something went wrong %v", err)
		return topTracks, err
	}
	log.Println(topTracks)
	return topTracks, nil
}

func ShouldInsertNewTrack(table string, spotifyUserId string) bool {
	db, e := Initialize()
	if e != nil {
		log.Printf("Something went wrong %v", e)
		return true
	}
	conn := db.Conn
	curTime := time.Now().Unix()
	stmtStr := `SELECT (EXTRACT(EPOCH FROM created_at))::bigint FROM $1 WHERE spotify_user_id=$2 order by created_at DESC LIMIT 1;`
	stmt, err := conn.Prepare(stmtStr)
	if err != nil {
		log.Println(err)
		return true
	}
	rows, errStmt := stmt.Query(table, spotifyUserId)
	if errStmt != nil {
		log.Printf("Something went wrong %v", errStmt)
		return true
	}
	defer rows.Close()
	for rows.Next() {
		lastUpdated := new(int64)
		err := rows.Scan(&lastUpdated)
		if err != nil {
			log.Printf("Something went wrong %v", err)
			return true
		}
		oneDay := int64(86400)
		if curTime-oneDay > *lastUpdated {
			return true
		}
	}
	return false
}
