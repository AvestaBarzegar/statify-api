package database

import (
	"database/sql/driver"
	"fmt"
)

func (v Track) Value() (driver.Value, error) {
	s := fmt.Sprintf("(%s, %s, %s, %s, %d)",
		v.TrackName,
		v.ArtistName,
		v.SpotifyTrackId,
		v.TrackImageURL,
		v.TrackRank)
	return []byte(s), nil
}
