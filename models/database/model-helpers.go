package database

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

type Artists []Artist

func (a *Artists) Scan(value interface{}) error {
	bytes, ok := value.([]byte) // input example 👉🏻 {"(david,38,url,1)","(david2,2,\"url 2\",2)"}
	if !ok {
		return errors.New("incompatible type from bytes")
	}
	source := string(bytes)
	var res Artists
	artists := strings.Split(source, "\",\"")
	for _, artist := range artists {
		for _, old := range []string{"\\\"", "\"", "{", "}", "(", ")"} {
			artist = strings.ReplaceAll(artist, old, "")
		}
		artistRawData := strings.Split(artist, ",")
		i, err := strconv.Atoi(artistRawData[1])
		if err != nil {
			return fmt.Errorf("parce ArtistRank raw data (%s) in %d iteration error: %v", artist, i, err)
		}
		res = append(res, Artist{
			ArtistName:      artistRawData[0],
			SpotifyArtistId: artistRawData[1],
			ArtistImageURL:  artistRawData[2],
			ArtistRank:      i,
		})
	}
	*a = res
	return nil
}
