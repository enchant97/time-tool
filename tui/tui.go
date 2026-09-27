package tui

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func drawTime(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
	timeStr := time.Now().Format(time.RFC3339)
	tview.Print(screen, timeStr, x, height/2, width, tview.AlignCenter, tcell.ColorWhite)
	return 0, 0, 0, 0
}

func Entrypoint() error {
	app := tview.NewApplication()
	app.SetTitle("TUI Time Tool")

	view := tview.NewBox().
		SetDrawFunc(drawTime).
		SetBorder(true).
		SetTitle("Current Time")

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			app.Draw()
		}
	}()

	return app.SetRoot(view, true).Run()
}
