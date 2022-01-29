package main

import (
	"C"

	"github.com/ilyapashuk/go-speechd"
)

var _spd *speechd.SpeechdSession

//export Initialize
func Initialize() {
	spd, err := speechd.Open()
	if err != nil {
		print("Error encountered")
		return
	}

	_spd = spd
}

//export Speak
func Speak(text string, interrupt bool) {
	if interrupt {
		_spd.Stop()
	}

	_spd.Speak(text)
}

//export Close
func Close() {
	_spd.Close()
}

func main() {
	Initialize()

	Speak("hello", false)
	Speak("shoaib", false)

	Close()
}
