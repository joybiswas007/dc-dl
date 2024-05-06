package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func CheckURL(comicURL string) bool {
	parsedURL, err := url.Parse(comicURL)
	if err != nil {
		fmt.Println(err)
	}
	path := strings.Trim(parsedURL.Path, "/")
	segments := strings.Split(path, "/")
	for _, s := range segments {
		if s == "digital-comics" {
			return true
		}
	}
	return false
}
