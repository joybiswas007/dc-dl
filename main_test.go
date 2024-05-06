package main

import (
	"dc/client"
	"dc/cookies"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	s := client.DCSession{}
	cookieHeader, err := cookies.ParseCookies("cookies/cookies.txt")
	if err != nil {
		t.Log(err)
		os.Exit(1)
	}
	s.Cookieheader = cookieHeader

	comicURL := "https://www.dhakacomics.com/digital-comics/durjoy-1"
	// comicURL := "https://www.dhakacomics.com/digital-comics/durjoy-15"

	data, err := s.GetComicData(comicURL)
	if err != nil {
		t.Log(err)
	}

	if data[0].WpNonce == "" {
		t.Log("You can't download this comic")
		os.Exit(1)
	}

	token, err := s.GetToken(data[0].ComicID, data[0].WpNonce)
	if err != nil {
		t.Log(err)
	}

	if token.Token == "" {
		t.Log("not tokens")
	}

}
