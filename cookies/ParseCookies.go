package cookies

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// parseCookies reads cookies from a file and returns the cookie header string
func ParseCookies(filename string) (string, error) {
	// Open the cookies file
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice to store the cookies
	var cookies []*http.Cookie

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines and lines starting with "#" (comments)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse the line into a cookie
		cookie, err := parseCookie(line)
		if err != nil {
			return "", fmt.Errorf("error parsing cookie: %v", err)
		}

		// Add the cookie to the slice
		cookies = append(cookies, cookie)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error scanning file: %v", err)
	}

	// Construct the cookie header string
	cookieHeader := ""
	for _, cookie := range cookies {
		if cookieHeader != "" {
			cookieHeader += "; "
		}
		cookieHeader += fmt.Sprintf("%s=%s", cookie.Name, cookie.Value)
	}

	return cookieHeader, nil
}

// parseCookie parses a line from the cookies.txt file into a Cookie struct
func parseCookie(line string) (*http.Cookie, error) {
	parts := strings.Split(line, "\t")

	// Check if the cookie line has enough parts
	if len(parts) < 7 {
		return nil, fmt.Errorf("invalid cookie line: %s", line)
	}

	secure := parts[2] == "TRUE"
	httpOnly := parts[3] == "TRUE"

	cookie := &http.Cookie{
		Name:     parts[5],
		Value:    parts[6],
		Domain:   parts[0],
		Path:     parts[2],
		Secure:   secure,
		HttpOnly: httpOnly,
	}

	return cookie, nil
}
