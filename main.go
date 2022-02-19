package main

import (
	_ "embed"
	"encoding/json"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/rrune/installer/installer"
)

//go:embed config.json
var configJson []byte

func main() {
	inst := installer.New()
	err := json.Unmarshal(configJson, &inst)
	if err != nil {
		panic(err)
	}

	a := app.New()
	w := a.NewWindow("Installer")

	dest := binding.NewString()
	out := binding.NewString()

	d := dialog.NewFolderOpen(func(lu fyne.ListableURI, e error) {
		if lu != nil {
			inst.SetDest(lu.Path())
			dest.Set(inst.Dest)
		}
	}, w)

	path := widget.NewLabel("Selected path:")
	input := widget.NewLabelWithData(dest)
	inputContainer := container.New(layout.NewHBoxLayout(), path, input)

	dlgBtn := widget.NewButton("Choose Path", func() {
		d.Show()
	})
	pathContainer := container.New(layout.NewVBoxLayout(), dlgBtn, inputContainer)

	outLabel := widget.NewLabelWithData(out)

	progress := widget.NewProgressBarInfinite()
	progress.Hide()

	finishBtn := widget.NewButton("Finish", func() {
		os.Exit(0)
	})
	finishBtn.Hide()

	installBtn := widget.NewButton("Install", func() {
		if inst.Dest != "" {
			progress.Show()

			out.Set("Downloading")
			err = inst.Download()
			if err != nil {
				out.Set("Error while downloading")
			}

			out.Set("Installing")
			err = inst.Unzip()
			if err != nil {
				out.Set("Error while installing")
			}

			out.Set("Cleaning up")
			err = os.Remove(inst.Temp)
			if err != nil {
				out.Set("Error while cleaning up")
			}

			out.Set("Done")
			progress.Stop()
			finishBtn.Show()
		} else {
			out.Set("Select a directory")
		}
	})

	content := container.New(layout.NewVBoxLayout(), pathContainer, installBtn, outLabel, progress, finishBtn)

	w.SetContent(content)
	w.ShowAndRun()
}
