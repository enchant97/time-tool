package tui

import (
	"time"

	"github.com/enchant97/time-tool/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type configHandler struct {
	config core.Config
}

func (c *configHandler) Replace(newConfig core.Config) {
	c.config = newConfig
}

func (c *configHandler) Get() core.Config {
	return c.config
}

var globalConfig configHandler

func createPage(inner tview.Primitive, name string) tview.Primitive {
	return tview.NewFrame(inner).
		SetBorders(2, 2, 2, 2, 4, 4).
		AddText("TUI Time Tool", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(name, true, tview.AlignCenter, tcell.ColorRed).
		AddText("Use <CTRL+C> to exit", false, tview.AlignCenter, tcell.ColorGreen).
		AddText("Use <TAB> and <RETURN> to navigate", false, tview.AlignCenter, tcell.ColorGreen)
}

func createConfigPage(switchToMenu func()) tview.Primitive {
	storedConfig := globalConfig.Get()
	form := tview.NewForm().
		AddInputField("Location", storedConfig.Location, 20, nil, func(text string) { storedConfig.Location = text }).
		AddInputField("Layout", storedConfig.Layout, 20, nil, func(text string) { storedConfig.Layout = text }).
		AddButton("Save", func() {
			storedConfig.DefaultUnset()
			err := core.WriteConfig(storedConfig)
			if err != nil {
				panic(err)
			}
			globalConfig.Replace(storedConfig)
			switchToMenu()
		}).
		AddButton("Menu", switchToMenu).
		AddButton("Reset", func() {
			newConfig := core.Config{}
			newConfig.DefaultUnset()
			err := core.WriteConfig(newConfig)
			if err != nil {
				panic(err)
			}
			globalConfig.Replace(storedConfig)
			switchToMenu()
		})
	return createPage(form, "Config")
}

func createTimePage(app *tview.Application, tt *time.Ticker, switchToMenu func()) tview.Primitive {
	inner := tview.NewModal().
		AddButtons([]string{"Menu"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) { switchToMenu() })
	go func() {
		for {
			t := <-tt.C
			app.QueueUpdateDraw(func() {
				config := globalConfig.Get()
				timeString, err := core.TimeToHuman(t, config)
				if err != nil {
					panic(err)
				}
				inner.SetText(timeString)
			})
		}
	}()
	return createPage(inner, "Current Time")
}

func createMenuPage(app *tview.Application, pages *tview.Pages) tview.Primitive {
	inner := tview.NewForm().
		AddButton("Time", func() {
			tt := time.NewTicker(time.Second)
			timePage := createTimePage(app, tt, func() {
				tt.Stop()
				pages.RemovePage("time")
			})
			pages.AddAndSwitchToPage("time", timePage, true)
		}).
		AddButton("Config", func() {
			configPage := createConfigPage(func() { pages.RemovePage("config") })
			pages.AddAndSwitchToPage("config", configPage, true)
		}).
		AddButton("Quit", func() { app.Stop() })
	return createPage(inner, "Menu")
}

func Entrypoint(initialConfig core.Config) error {
	globalConfig = configHandler{config: initialConfig}
	app := tview.NewApplication()
	app.SetTitle("TUI Time Tool")
	pages := tview.NewPages()

	menu := createMenuPage(app, pages)
	pages.AddPage("menu", menu, true, true)

	return app.
		SetRoot(pages, true).
		EnableMouse(true).
		EnablePaste(true).
		SetFocus(pages).
		Run()
}
