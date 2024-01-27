package dhakacomics

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// TokenResponse is the structure that matches the JSON response from the API.
type TokenResponse struct {
	Token string `json:"token"`
}

// getToken makes an API call and returns a token or an error.
func GetComicToken(id, wpNonce, refererUrl string) (string, error) {
	UserAgent := os.Getenv("USER_AGENT")
	Cookie := os.Getenv("COOKIE")
	UmpToken := os.Getenv("X_CSRF_UMP_TOKEN")
	ViewComicUri := os.Getenv("VIEW_COMIC")

	client := &http.Client{}
	payload := fmt.Sprintf("id=%s&wp_nonce=%s", id, wpNonce)
	var data = strings.NewReader(payload)
	req, err := http.NewRequest("POST", ViewComicUri, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "www.dhakacomics.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("cookie", Cookie)
	req.Header.Set("origin", "https://www.dhakacomics.com")
	req.Header.Set("referer", refererUrl)
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", UserAgent)
	req.Header.Set("x-csrf-ump-token", UmpToken)
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// Parse the JSON response.
	var tokenResp TokenResponse
	err = json.Unmarshal(bodyText, &tokenResp)
	if err != nil {
		return "", err
	}

	// Return the token.
	return tokenResp.Token, nil
}
