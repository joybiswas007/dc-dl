package client

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (s *DCSession) GetComicData(comicURL string) ([]ComicData, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", comicURL, nil)
	if err != nil {
		return []ComicData{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:125.0) Gecko/20100101 Firefox/125.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.dhakacomics.com/my-account")
	req.Header.Set("Cookie", s.Cookieheader)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("TE", "trailers")

	resp, err := client.Do(req)
	if err != nil {
		return []ComicData{}, err
	}
	defer resp.Body.Close()

	// Parse the page content
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []ComicData

	// Extract the page banner data
	doc.Find("#page-banner").Each(func(i int, s *goquery.Selection) {
		var data ComicData

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
