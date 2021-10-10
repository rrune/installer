package util

import "github.com/gen2brain/dlgs"

func Check(err error) {
	if err != nil {
		_, err := dlgs.Error("Fehler", "Es ist ein Fehler aufgetreten.")
		if err != nil {
			panic(err)
		}
		panic(err)
	}
}
