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

var _spd *speechd.SpeechdSession = nil

//export Initialize
func Initialize() int {
	spd, err := speechd.Open()
	if err != nil {
		print("\nError encountered while initializing speech dispatcher:\n\t" + err.Error() + "\n")
		return -1
	}

	_spd = spd
	return 1
}

//export Speak
func Speak(text string, interrupt bool) int {
	if _spd == nil {
		return -1
	}

	if interrupt {
		_spd.Stop()
		_spd.Cancel()
	}

	_, err := _spd.Speak(text)
	if err != nil {
		print("\nError encountered while speaking with speech dispatcher:\n\t" + err.Error() + "\n")
		return -1
	}

	return 1
}

//export Close
func Close() int {
	if _spd == nil {
		return -1
	}

	_spd.Close()
	return 1
}

func main() {
	Initialize()

	Speak("131", false)
	Speak("shoaib", true)

	Close()
}
