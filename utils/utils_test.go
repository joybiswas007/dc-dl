package utils

import (
	"os"
	"testing"
)

func TestChecURL(t *testing.T) {
	if !CheckURL("https://www.dhakacomics.com/digital-comics/durjoy-1") {
		t.Log("Invalid url")
		os.Exit(1)
	}
	t.Log("Valid url")
}
