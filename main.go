package main

import (
	_ "embed"
	"encoding/json"

	"github.com/ncruces/zenity"
	"github.com/rrune/installer/installer"
	. "github.com/rrune/installer/util"
)

//go:embed config.json
var configJson []byte

func main() {
	installer := installer.New()
	err := json.Unmarshal(configJson, &installer)
	Check(err)

	err = zenity.Question("Möchtest du <Programm> installieren?",
		zenity.Title("<Programm>"),
		zenity.QuestionIcon)

	if err == nil {
		installer.Dest, err = zenity.SelectFile(
			zenity.Filename(""),
			zenity.Directory())
		Check(err)

		installer.Install()

		zenity.Info("Installiert.",
			zenity.Title("<Programm>"),
			zenity.InfoIcon)
	}
}
