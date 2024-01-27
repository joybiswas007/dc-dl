package dhakacomics

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
)

func DownloadComic(name, token, refererUrl string) error {
	UserAgent := os.Getenv("USER_AGENT")
	Cookie := os.Getenv("COOKIE")
	ViewComicUri := os.Getenv("VIEW_COMIC")

	client := &http.Client{}
	url := fmt.Sprintf("%s?token=%s", ViewComicUri, token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("authority", "www.dhakacomics.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("cookie", Cookie)
	req.Header.Set("referer", refererUrl)
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the response status code indicates success
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download comic. Status code: %d", resp.StatusCode)
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	downloadsDir := filepath.Join(pwd, "Comics Downloads")
	if _, err := os.Stat(downloadsDir); os.IsNotExist(err) {
		os.Mkdir(downloadsDir, 0755)
	}
	filePath := filepath.Join(downloadsDir, name+".pdf")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// Initialize the progress bar

	bar := pb.Full.Start64(resp.ContentLength)

	// Create a proxy reader that updates the progress bar

	reader := bar.NewProxyReader(resp.Body)

	// Copy the response body to the file
	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}
	// Finish the progress bar
	bar.Finish()
	fmt.Printf("Comic '%s' downloaded successfully.\n", name)
	return nil
}
