package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (s *DCSession) GetToken(comicID, wpNonce string) (*Token, error) {
	client := &http.Client{}
	payload := fmt.Sprintf("id=%s&wp_nonce=%s", comicID, wpNonce)
	var data = strings.NewReader(payload)
	req, err := http.NewRequest("POST", "https://www.dhakacomics.com/view-comic", data)
	if err != nil {
		return &Token{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:125.0) Gecko/20100101 Firefox/125.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	// req.Header.Set("X-CSRF-UMP-TOKEN", "")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Origin", "https://www.dhakacomics.com")
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", s.ComicURL)
	req.Header.Set("Cookie", s.Cookieheader)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	if err != nil {
		return &Token{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Token{}, err
	}
	var token Token
	err = json.Unmarshal(bodyText, &token)
	if err != nil {
		return &Token{}, err
	}

	return &token, nil
}
