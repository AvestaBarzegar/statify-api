package http

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

type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
