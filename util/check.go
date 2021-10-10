package util

import "github.com/sqweek/dialog"

func Check(err error) {
	if err != nil {
		dialog.Message("%s", "Ein Fehler ist aufgetreten").Title("<Programm>").Info()
		panic(err)
	}
}
