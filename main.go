/*
	Installation:-
		Copy the libspeechdwrapper.so(in the lib folder) file to /usr/lib like this:-
			sudo cp ./lib/libspeechdwrapper.so /usr/lib

	Building Shared Library(Optional):-
		Use the following command to generate the shared library:-
			go build -buildmode=c-shared -o ./lib/libspeechdwrapper.so .

*/

package main

import "C"

import (
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

	Speak("131", false)
	Speak("shoaib", false)

	Close()
}
