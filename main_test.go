package main

import (
	"testing"
	"time"
)

func TestNormal(t *testing.T) {
	Initialize()
	Speak("hello", false)
	Close()
}

func TestInterrupt(t *testing.T) {
	Initialize()
	Speak("This is a very very very long message.", false)
	Speak("I have the power to interrupt", true)
	Close()
}

func TestInterruptWithDelay(t *testing.T) {
	Initialize()
	Speak("This is a very very very long message.", false)
	time.Sleep(time.Millisecond * 500)
	Speak("I have the power to interrupt", true)
	Close()
}
