package utils

import (
	"fmt"
	"time"
	"embed"

	"github.com/rivo/tview"
)

// Reference to the channel for UI updates
var UIEventsChan chan func()
// Define the avatar's arts via their paths
var BasePath = GetBasePath()

// ==============================
// EMBEDS
// ==============================

//go:embed ascii-arts/**
var ASCIIFS embed.FS

//go:embed assets/**
var ASSETSFS embed.FS

// ==============================
// ASCII ART LOADING
// ==============================

// LoadASCII loads ASCII art from a file and returns it as a string
func LoadASCII(path string) string {
	content, err := ASCIIFS.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Failed to load %s: %v", path, err))
		return ""
	}
	return string(content)
}

// ==============================
// BLINKING WAIFU
// ==============================

// StartBlinking starts a blinking animation for waifu ASCII art.
// Returns a stop channel to terminate the blinking.
func StartBlinking(app *tview.Application, waifuArt *tview.TextView,
	head, blinkHead, body *string, interval time.Duration) chan bool {

	stop := make(chan bool, 1)
	var last string

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-stop:
				return
			case <-ticker.C:
				// Decrease Happiness
				DecreaseHappiness(1)
				// Show blink frame
				blinkText := *blinkHead + "\n" + *body
				if blinkText != last && UIEventsChan != nil {
					UIEventsChan <- func() {
						waifuArt.SetText(blinkText)
					}
					last = blinkText
				}

				// Restore normal frame after short delay
				normalText := *head + "\n" + *body
				go time.AfterFunc(200*time.Millisecond, func() {
					if normalText != last && UIEventsChan != nil {
						UIEventsChan <- func() {
							waifuArt.SetText(normalText)
						}
						last = normalText
					}
				})
			}
		}
	}()

	return stop
}
