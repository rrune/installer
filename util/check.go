package util

import "github.com/ncruces/zenity"

func Check(err error) {
	if err != nil {
		zenity.Error("Ein Fehler ist aufgetreten.",
			zenity.Title("<Programm>"),
			zenity.ErrorIcon)
		panic(err)
	}
}
