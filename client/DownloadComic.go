package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
)

func (s *DCSession) DownloadComic(comicName string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.dhakacomics.com/view-comic?token=%s", s.Token), nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:125.0) Gecko/20100101 Firefox/125.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Referer", s.ComicURL)
	req.Header.Set("DNT", "1")
	req.Header.Set("Sec-GPC", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", s.Cookieheader)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download comic. Status code: %d", resp.StatusCode)
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	downloadsDir := filepath.Join(pwd, "DC-DL")
	if _, err := os.Stat(downloadsDir); os.IsNotExist(err) {
		os.Mkdir(downloadsDir, 0755)
	}
	filePath := filepath.Join(downloadsDir, comicName+".pdf")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	template := `{{ red " ⃝'<" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}} {{counters . "%s / %s" | rndcolor}} {{string . "my_green_string" | green}} {{string . "my_blue_string" | blue}}`

	bar := pb.ProgressBarTemplate(template).Start64(resp.ContentLength)

	reader := bar.NewProxyReader(resp.Body)

	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}
	bar.Finish()
	return nil
}
