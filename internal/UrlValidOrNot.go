package dhakacomics

import (
	"net/url"
	"strings"
)

func UrlValidOrNot(u *url.URL, segment string) bool {
	path := strings.Trim(u.Path, "/")
	segments := strings.Split(path, "/")
	for _, s := range segments {
		if s == segment {
			return true
		}
	}
	return false
}
