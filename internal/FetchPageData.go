package dhakacomics

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// PageData represents the structured data you want to extract from the page.
type PageData struct {
	Title     string
	Author    string
	AuthorURL string
	CoverURL  string
	ComicID   string
	WpNonce   string
}

// FetchPageData will fetch the page data from the specified URL.
func FetchPageData(url string) ([]PageData, error) {
	UserAgent := os.Getenv("USER_AGENT")
	Cookie := os.Getenv("COOKIE")
	// Initialize HTTP client
	client := &http.Client{}

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("authority", "www.dhakacomics.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("cookie", Cookie)
	req.Header.Set("referer", "https://www.dhakacomics.com/my-account")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", UserAgent)
	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the page content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []PageData

	// Extract the page banner data
	doc.Find("#page-banner").Each(func(i int, s *goquery.Selection) {
		var data PageData

		data.Title = s.Find(".page-profile-title").Find("h1").Text()
		data.Author = s.Find(".page-profile-title").Find("h2").Text()
		var exists bool
		data.AuthorURL, exists = s.Find(".page-profile-title").Find("h2 a").Attr("href")
		if !exists {
			log.Println("Author URL not found")
		}

		coverURL, exists := s.Attr("style")
		if exists {
			re, err := regexp.Compile(`url\((.*?)\)`)
			if err != nil {
				log.Println(err)
			} else {
				match := re.FindStringSubmatch(coverURL)
				if len(match) >= 2 {
					data.CoverURL = match[1]
				}
			}
		} else {
			log.Println("Cover image not found.")
		}

		results = append(results, data)
	})

	// Extract the main content data
	doc.Find("section#main-content").Each(func(i int, s *goquery.Selection) {
		comicID, exist := s.Find(".post-ratings").Attr("id")
		if exist {
			id := strings.Replace(comicID, "post-ratings-", "", 1)
			if len(results) > i {
				results[i].ComicID = id
			}
		} else {
			log.Println("No ID found.")
		}
	})

	// Extract script data for wp_nonce
	dataScript := doc.Find("script").Eq(8).Text()
	nonceRegex := regexp.MustCompile(`wp_nonce:\s*"([^"]+)"`)
	nonceMatches := nonceRegex.FindStringSubmatch(dataScript)
	if len(nonceMatches) > 1 {
		wpNonce := nonceMatches[1]
		for i := range results {
			results[i].WpNonce = wpNonce
		}
	}

	return results, nil
}
