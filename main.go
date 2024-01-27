package main

import (
	dhakacomics "dhakacomics-go/internal"
	"fmt"

	"strings"

	"net/url"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error loading .env file")
	}
	var userInput string
	fmt.Println("Welcome to DhakaComics Downloader.")
	fmt.Println("Enter digital comics url: ")
	fmt.Scanln(&userInput)

	for {
		if strings.ToLower(strings.TrimSpace(userInput)) == "q" {
			fmt.Println("Exiting the downloader.")
			break
		}

		parsedUrl, err := url.Parse(userInput)
		if err != nil {
			fmt.Println("Error parsing url", err)
			break
		}

		if !dhakacomics.UrlValidOrNot(parsedUrl, "digital-comics") {
			fmt.Println("Invalid url")
			break
		}

		data, err := dhakacomics.FetchPageData(userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		token, err := dhakacomics.GetComicToken(data[0].ComicID, data[0].WpNonce, userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = dhakacomics.DownloadComic(data[0].Title, token, userInput)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Enter another digital comics url (or 'q' to quit): ")
		fmt.Scanln(&userInput)
	}

}
