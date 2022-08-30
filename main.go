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

/**
* Description: Connects with the speechd socket. If no screen reader is activated, it fails to connect on the first try although if we again try to initialize after 3-4 seconds, it connects successfully.
* Return:      1 if successful and -1 if unsuccessful.
*/
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

/**
* Description: Speaks the given text.
* Parameters:  text = the text to speak. Important!! string variable is accessed in other languages in a special way, refer to https://github.com/vladimirvivien/go-cshared-examples for examples
*              interrupt = whether to cancel the previous speech or not.
* Return:      1 if successful and -1 if unsuccessful.
*/
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

/**
* Description: Disconnects with the speechd socket.
* Return:      1 if successful and -1 if unsuccessful.
*/
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

	Speak("Hello World", true)

	Close()
}