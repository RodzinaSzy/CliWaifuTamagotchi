package utils

import (
	"sync"

	"github.com/rivo/tview"
)

var (
	Happiness        = 1000                 // Initial happiness value
	CurrentBar       = ""                   // Place holder for the current bar state
	happinessMutex   sync.Mutex             // Mutex to protect concurrent access to Happiness
	HappinessBarRef  *tview.TextView        // Link to the happiness bar itself so we dynamically update it
	HeadASCII        *string                // Current head
	BlinkHeadASCII   *string                // Current blinking head
)

// ==============================
// Load all expressions once
// ==============================
var (
	neutral       = LoadASCII(BasePath + "/expressions/neutral")
	neutralBlink  = LoadASCII(BasePath + "/expressions/neutral-blink")
	confused      = LoadASCII(BasePath + "/expressions/confused")
	confusedBlink = LoadASCII(BasePath + "/expressions/confused-blink")
	bored         = LoadASCII(BasePath + "/expressions/bored")
	boredBlink    = LoadASCII(BasePath + "/expressions/bored-blink")
	sad           = LoadASCII(BasePath + "/expressions/sad")
	sadBlink      = LoadASCII(BasePath + "/expressions/sad-blink")
)

func setExpression(head, blink string) {
	if HeadASCII != nil && BlinkHeadASCII != nil && *HeadASCII != head {
		*HeadASCII = head
		*BlinkHeadASCII = blink
	}
}

// ==============================
// Decrease Happiness
// ==============================
func DecreaseHappiness(n int) {
	happinessMutex.Lock()
	defer happinessMutex.Unlock()

	if Happiness > 0 {
		Happiness -= n
		if Happiness < 0 {
			Happiness = 0
		}
		// Optimize it
		if Happiness%2 == 0{
			updateBar()
		}
	}
}

// ==============================
// Increase Happiness
// ==============================
func IncreaseHappiness(n int) {
	happinessMutex.Lock()
	defer happinessMutex.Unlock()

	if Happiness < 1000 {
		Happiness += n
		if Happiness > 1000 {
			Happiness = 1000
		}
		updateBar()
	}
}

// ==============================
// Returns visual bar string
// ==============================
func GetHappinessBar() {
	switch {
	case Happiness > 900:
		CurrentBar = "██████████"
		setExpression(neutral, neutralBlink)
	case Happiness > 800:
		CurrentBar =  "█████████░"
		setExpression(neutral, neutralBlink)
	case Happiness > 700:
		CurrentBar =  "████████░░"
		setExpression(confused, confusedBlink)
	case Happiness > 600:
		CurrentBar =  "███████░░░"
		setExpression(confused, confusedBlink)
	case Happiness > 500:
		CurrentBar =  "██████░░░░"
		setExpression(bored, boredBlink)
	case Happiness > 400:
		CurrentBar =  "█████░░░░░"
		setExpression(bored, boredBlink)
	case Happiness > 300:
		CurrentBar =  "████░░░░░░"
		setExpression(bored, boredBlink)
	case Happiness > 200:
		CurrentBar =  "███░░░░░░░"
		setExpression(sad, sadBlink)
	case Happiness > 100:
		CurrentBar =  "██░░░░░░░░"
		setExpression(sad, sadBlink)
	case Happiness > 0:
		CurrentBar =  "█░░░░░░░░░"
		setExpression(sad, sadBlink)
	default:
		CurrentBar =  "░░░░░░░░░░"
		setExpression(sad, sadBlink)
	}
}

// ==============================
// Internal UI update
// ==============================
func updateBar() {
	if HappinessBarRef != nil && UIEventsChan != nil {
		GetHappinessBar()
		barText := CurrentBar
		UIEventsChan <- func() {
			HappinessBarRef.SetText(barText)
		}
	}
}
