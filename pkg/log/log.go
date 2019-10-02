package log

import (
	"fmt"
	"log"
)

// Level logging level
type Level int64

// Logging logging level for cdp client
var Logging Level = LevelFatal | LevelProtocolErrors | LevelSession
var hook func(string)

// log levels for client logging
const (
	LevelFatal           = 0x01
	LevelProtocolErrors  = 0x02
	LevelProtocolMessage = 0x04 | LevelProtocolErrors
	LevelProtocolEvents  = 0x08 | LevelProtocolErrors
	LevelProtocolVerbose = LevelProtocolErrors | LevelProtocolMessage | LevelProtocolEvents
	LevelSession         = 0x10
)

// SetHook set hook function
func SetHook(hookf func(string)) {
	hook = hookf
}

// Printf Arguments are handled in the manner of fmt.Printf
func Printf(level Level, format string, v ...interface{}) {
	if level&Logging == level {
		print(fmt.Sprintf(format, v...))
	}
}

// Print Arguments are handled in the manner of fmt.Print
func Print(level Level, v ...interface{}) {
	if Logging&level != 0 {
		print(fmt.Sprint(v...))
	}
}

func print(message string) {
	if hook != nil {
		hook(message)
		return
	}
	log.Println(message)
}
