package main

import (
	"github.com/ilyapashuk/go-speechd"
)

var _spd *speechd.SpeechdSession

func Initialize() {
	spd, err := speechd.Open()
	if err != nil {
		print("Error encountered")
		return
	}

	_spd = spd
}

func Speak(text string, interrupt bool) {
	if interrupt {
		_spd.Stop()
	}

	_spd.Speak(text)
}

func Close() {
	_spd.Close()
}

func main() {
	Initialize()

	Speak("hello", false)
	Speak("shoaib", false)

	Close()
}
