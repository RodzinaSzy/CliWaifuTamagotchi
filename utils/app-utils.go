package utils

import (
	"bufio"
	"embed"
	"fmt"
	"time"

	"github.com/rivo/tview"
)

// Reference to the channel for UI updates
var UIEventsChan chan func()

// ==============================
// ASCII ART LOADING
// ==============================

// Embed the entire ascii-arts directory recursively
//
//go:embed ascii-arts/**
var ASCIIFS embed.FS

// LoadASCII loads ASCII art from a file and returns it as a string
func LoadASCII(path string) string {
	content, err := ASCIIFS.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Failed to load %s: %v", path, err))
		return ""
	}
	return string(content)
}

//go:embed assets/**
var ASSETSFS embed.FS

// LoadEncouragements loads a slice of lines from a text file
func LoadEncouragements(path string) ([]string, error) {
	file, err := ASSETSFS.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed reading %s: %w", path, err)
	}

	return lines, nil
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
