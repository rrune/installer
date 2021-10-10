package main

import (
	_ "embed"
	"encoding/json"

	"github.com/rrune/installer/installer"
	. "github.com/rrune/installer/util"
	"github.com/sqweek/dialog"
)

//go:embed config.json
var configJson []byte

func main() {
	installer := installer.New()
	err := json.Unmarshal(configJson, &installer)
	Check(err)

	ok := dialog.Message("%s", "Möchtest du <Programm> installiren?").Title("<Programm>").YesNo()
	if ok {
		installer.Dest, err = dialog.Directory().Title("Wähle das Installationsverzeichnis aus:").Browse()
		Check(err)
		installer.Install()
		dialog.Message("%s", "<Programm> wurde installiert,").Title("<Programm>").Info()
	}
}
