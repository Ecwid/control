package witness

import (
	"fmt"
	"log"
)

type wlog struct {
	hook  func(string)
	Level level
}

// Level logging level
type level int64

// log levels for client logging
const (
	LevelFatal           level = 0x01
	LevelProtocolErrors  level = 0x02 | LevelFatal
	LevelProtocolMessage level = 0x04 | LevelProtocolErrors
	LevelProtocolEvents  level = 0x08 | LevelProtocolErrors
	LevelProtocolVerbose level = LevelProtocolErrors | LevelProtocolMessage | LevelProtocolEvents
)

// SetHook set logging output hook
func (w wlog) SetHook(hookf func(string)) {
	w.hook = hookf
}

// Printf Arguments are handled in the manner of fmt.Printf
func (w wlog) Printf(forlevel level, format string, v ...interface{}) {
	if forlevel&w.Level == forlevel {
		w.print(fmt.Sprintf(format, v...))
	}
}

// Print Arguments are handled in the manner of fmt.Print
func (w wlog) Print(forlevel level, v ...interface{}) {
	if forlevel&w.Level == forlevel {
		w.print(fmt.Sprint(v...))
	}
}

func (w wlog) print(message string) {
	if w.hook != nil {
		w.hook(message)
		return
	}
	log.Println(message)
}
