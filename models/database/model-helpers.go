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
	// Make a string that is parseable
	source := string(bytes)

	// Replace all unnecessary characters
	for _, old := range []string{"\\\"", "\"", "{", "}", "(", "[", "]"} {
		source = strings.ReplaceAll(source, old, "")
	}
	var res Artists
	// Split up the whole string into an array of strings that represent each artist object
	artists := strings.Split(source, "),")
	// Remove all unnecessary braces
	artists[len(artists)-1] = strings.TrimRight(artists[len(artists)-1], ")")
	// If the column is empty then we can return an empty array
	if len(artists) == 1 && artists[0] == "" {
		a = &Artists{}
		return nil
	}
	for _, artist := range artists {
		artistRawData := strings.Split(artist, ",")
		i, err := strconv.Atoi(artistRawData[3])
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
