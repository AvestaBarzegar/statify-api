package bindings

type RequestAccessTokenQuery struct {
	GrantType   string `form:"grant_type"`
	Code        string `form:"code"`
	RedirectURI string `form:"redirect_uri"`
}

type SongInformation
