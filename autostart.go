package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/emersion/go-autostart"
)

func autostort() {
	if runtime.GOOS != "windows" {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		autostarapp := &autostart.App{
			Name:        "PSMTOOLS",
			DisplayName: "PSMTools",
			Exec:        []string{exPath, "--minimized"},
		}
		if !autostarapp.IsEnabled() {
			autostarapp.Enable()
		}
	} else {
		return
	}
	return
}
