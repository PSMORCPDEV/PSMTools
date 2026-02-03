package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	appItself := app.New()
	window := appItself.NewWindow("PSMTools")
	window.Resize(fyne.NewSize(400, 190))
	window.SetFixedSize(true)

	appItself.Settings().SetTheme(theme.DarkTheme())
	logs := widget.NewLabel("PSMT\n\n\n\n\n\nPSMDEV\n©2025-2026")
	logbutt := widget.NewButtonWithIcon("Log (β)", theme.FileTextIcon(), func() {})
	//logbutt := widget.NewButtonWithIcon("Логирование (β)", theme.FileTextIcon(), func() {
	//	logging(logs)
	//})
	buttons := container.NewVBox(
		widget.NewButtonWithIcon("Терминал", theme.DesktopIcon(), func() {}),
		logbutt,
		widget.NewButtonWithIcon("Utilities", theme.MoreHorizontalIcon(), func() {}),
		widget.NewButtonWithIcon("Regestry", theme.GridIcon(), func() {}),
		widget.NewButtonWithIcon("Help", theme.QuestionIcon(), func() {}),
	)
	SettingButton := widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {})
	content := container.NewBorder(
		nil, SettingButton, logs, buttons,
	)

	window.SetContent(content)
	window.ShowAndRun()
}
