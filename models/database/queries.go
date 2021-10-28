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
	stmtStr := `SELECT * FROM artists WHERE spotify_user_id=$1`
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
		err := rows.Scan(&topArtist.Id, &topArtist.SpotifyUserId, &topArtist.CreatedAt, &topArtist.TimeSpan, &topArtist.Artists)
		if err != nil {
			log.Printf("Something went wrong %v", err)
		}
		topArtists = append(topArtists, *topArtist)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Something went wrong %v", err)
	}
	return topArtists, nil
}
