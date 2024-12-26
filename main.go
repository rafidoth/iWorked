package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Stopwatch")

	var startTime time.Time
	var elapsedTime time.Duration
	var running bool
	var ticker *time.Ticker
	var updateTimer *time.Timer

	stopwatchTime := widget.NewLabel("00:00:00")

	updateTime := func() {
		if running {
			elapsedTime = time.Since(startTime)
			stopwatchTime.SetText(fmt.Sprintf("%02d:%02d:%02d", int(elapsedTime.Hours())%24, int(elapsedTime.Minutes())%60, int(elapsedTime.Seconds())%60))
		}
	}

	togglePauseStartBtn := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)

	togglePauseStartBtn.OnTapped = func() {
		if running {
			running = false
			ticker.Stop()
			if updateTimer != nil {
				updateTimer.Stop() // Stop the update timer
			}
			togglePauseStartBtn.SetIcon(theme.MediaPlayIcon())
		} else {
			// Start the stopwatch
			running = true
			startTime = time.Now().Add(-elapsedTime) // Adjust to account for time passed during pause
			ticker = time.NewTicker(time.Second)
			go func() {
				for range ticker.C {
					updateTime()
				}
			}()
			togglePauseStartBtn.SetIcon(theme.MediaPauseIcon())
		}
	}

	resetBtn := widget.NewButtonWithIcon("", theme.MediaReplayIcon(), func() {
		if running {
			ticker.Stop()
			running = false
		}
		elapsedTime = 0
		stopwatchTime.SetText("00:00:00")
		togglePauseStartBtn.SetIcon(theme.MediaPlayIcon())
	})

	Container := container.NewHBox(stopwatchTime, togglePauseStartBtn, resetBtn)
	window.SetContent(Container)
	window.ShowAndRun()
}
