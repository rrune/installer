package main

import (
	_ "embed"
	"encoding/json"

	"github.com/gen2brain/dlgs"
	"github.com/rrune/installer/installer"
	. "github.com/rrune/installer/util"
)

//go:embed config.json
var configJson []byte

func main() {
	installer := installer.New()
	err := json.Unmarshal(configJson, &installer)
	Check(err)

	ok, err := dlgs.Question("<Programm>", "Möchtest du <Programm> installieren?", true)
	Check(err)
	if ok {
		dest, _, err := dlgs.File("Wähle das Installationsverzeichnis:", "", true)
		Check(err)
		installer.SetDest(dest)

		installer.Install()
		_, err = dlgs.Info("<Programm>", "<Programm> wurde installiert.")
		Check(err)
	}
}
