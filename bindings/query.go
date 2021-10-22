package bindings

type RequestAccessTokenQuery struct {
	GrantType   string `form:"grant_type"`
	Code        string `form:"code"`
	RedirectURI string `form:"redirect_uri"`
}

// ?id={id}&artist={artist}&track={track}
type SongInformationQuery struct {
	SongId string `form:"id"`
	Artist string `form:"artist"`
	Track  string `form:"track"`
}
