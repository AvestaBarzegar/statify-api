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

type TrackTraitsResponse struct {
	Danceability     float64 `json:"danceability"`
	Energy           float64 `json:"energy"`
	Key              int     `json:"key"`
	Loudness         float64 `json:"loudness"`
	Mode             int     `json:"mode"`
	Speechiness      float64 `json:"speechiness"`
	Acousticness     float64 `json:"acousticness"`
	Instrumentalness float64 `json:"instrumentalness"`
	Liveness         float64 `json:"liveness"`
	Valence          float64 `json:"valence"`
	Tempo            float64 `json:"tempo"`
	Type             string  `json:"type"`
	ID               string  `json:"id"`
	URI              string  `json:"uri"`
	TrackHref        string  `json:"track_href"`
	AnalysisURL      string  `json:"analysis_url"`
	DurationMs       int     `json:"duration_ms"`
	TimeSignature    int     `json:"time_signature"`
}

type LyricsResponseModel struct {
	Lyrics string `json:"lyrics"`
}

type TrackTraitsClientResponseModel struct {
	Lyrics      string              `json:"lyrics"`
	TrackTraits TrackTraitsResponse `json:"traits"`
}
