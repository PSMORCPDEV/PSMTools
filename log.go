package main

import (
	"fyne.io/fyne/v2/widget"
)

func logging(label *widget.Label) {
	label.SetText("Логирование активно")
	label.SetText(label.Text + "\nМеня зовут коннор и проверю\nкак это работает на самом деле")
}
