package client

type DCSession struct {
	Cookieheader string
	Token        string
	ComicURL     string
}

type ComicData struct {
	Title     string
	Author    string
	AuthorURL string
	CoverURL  string
	ComicID   string
	WpNonce   string
}

type Token struct {
	Token string `json:"token"`
}
