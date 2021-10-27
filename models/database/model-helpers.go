package database

import (
	"database/sql/driver"
	"fmt"
)

func (t Track) Value() (driver.Value, error) {
	s := fmt.Sprintf("(%s, %s, %s, %s, %d)",
		t.TrackName,
		t.ArtistName,
		t.SpotifyTrackId,
		t.TrackImageURL,
		t.TrackRank)
	return []byte(s), nil
}

func (a Artist) Value() (driver.Value, error) {
	s := fmt.Sprintf("(%s, %s, %s, %d)",
		a.ArtistName,
		a.SpotifyArtistId,
		a.ArtistImageURL,
		a.ArtistRank)
	return []byte(s), nil
}
