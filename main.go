package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var mainWindow fyne.Window

func main() {
	autostort()
	appItself := app.New()
	mainWindow = appItself.NewWindow("PSMTools")
	mainWindow.Resize(fyne.NewSize(400, 190))
	mainWindow.SetFixedSize(true)

	appItself.Settings().SetTheme(theme.DarkTheme())
	logs := widget.NewLabel("PSMTools\n\n\n\n\n\n\nPSMDEV\n©2025-2026")
	logbutt := widget.NewButtonWithIcon(("Log (β)"), theme.FileTextIcon(), func() {})
	buttons := container.NewVBox(
		widget.NewButtonWithIcon(("Terminal"), theme.DesktopIcon(), func() {}),
		logbutt,
		widget.NewButtonWithIcon(("Utilities"), theme.MoreHorizontalIcon(), func() {}),
		widget.NewButtonWithIcon(("Registry"), theme.GridIcon(), func() {}),
		widget.NewButtonWithIcon(("Help"), theme.QuestionIcon(), func() {}),
	)
	SettingButton := widget.NewButtonWithIcon(("Settings (does nothing bruh)"), theme.SettingsIcon(), func() {
	})
	content := container.NewBorder(
		nil, SettingButton, logs, buttons,
	)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
