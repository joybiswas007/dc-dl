package cookies

import "testing"

func TestCookies(t *testing.T) {
	cookieHeader, err := ParseCookies("cookies.txt")
	if err != nil {
		t.Log(err)
	}
	if cookieHeader != "" {
		t.Log("Reading cookie file success!")
	}
}
