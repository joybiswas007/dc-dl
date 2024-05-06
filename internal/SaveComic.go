package internal

import (
	"dc/client"
	"dc/cookies"
	"dc/utils"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func SaveComic(cmd *cobra.Command, args []string) {
	s := client.DCSession{}
	comicURL, _ := cmd.Flags().GetString("url")

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "💥",
	})

	if comicURL == "" {
		logger.Warn("URL can't be empty (｡•́︿•̀｡)")
		os.Exit(1)
	}
	s.ComicURL = comicURL

	if !utils.CheckURL(comicURL) {
		logger.Warn("URL is invalid (·•᷄∩•᷅ )")
		os.Exit(1)
	}

	cookieHeader, err := cookies.ParseCookies("cookies/cookies.txt")
	if err != nil {
		logger.Warn("Cookies.txt file NOT found inside cookies diretory (´•︵•`).")
		os.Exit(1)
	}

	s.Cookieheader = cookieHeader

	data, err := s.GetComicData(comicURL)
	if err != nil {
		logger.Warn("Failed to fetch comic data ˙◠˙.")
		os.Exit(1)
	}

	if data[0].WpNonce == "" {
		logger.Error("(⊙ _ ⊙ ) You need premium account to download comics.")
		os.Exit(1)
	}

	token, err := s.GetToken(data[0].ComicID, data[0].WpNonce)
	if err != nil {
		logger.Warn("Failed to fetch tokens ( • ᴖ • ｡).")
		os.Exit(1)
	}

	s.Token = token.Token

	logger.Infof("Downloading comic: %s by author: %s.", data[0].Title, data[0].Author)

	err = s.DownloadComic(data[0].Title)
	if err != nil {
		logger.Warn("Failed to download comic (ᴗ_ ᴗ。).")
		os.Exit(1)
	}

	logger.Infof("Comic: %s by author: %s downloaded.", data[0].Title, data[0].Author)
}
