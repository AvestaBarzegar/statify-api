package bindings

type RequestAccessTokenHeaders struct {
	ContentType   string `header:"Content-Type"`
	Authorization string `header:"Authorization"`
}
