package mlog

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync/atomic"

	"github.com/mattn/go-isatty"
)

const (
	gray = 37
)

// A Logger is a standard log.Logger with Debug family functions.
type Logger struct {
	log.Logger
	isDebug int32 // atomic boolean
}

// New creates a new Logger and enables the debug log mode when found
// DEBUG environment variable set to one of the following values:
//   - yes
//   - true
//   - 1
func New() *Logger {
	var isDebug int32
	switch strings.ToLower(os.Getenv("DEBUG")) {
	case "yes", "true", "1":
		isDebug = 1
	}
	return &Logger{
		Logger:  *log.Default(),
		isDebug: isDebug,
	}
}

var std = New()

// EnableDebug enables the debug log mode and has higher priority than DEBUG environment variable.
func (l *Logger) EnableDebug() {
	atomic.StoreInt32(&l.isDebug, 1)
}

// DisableDebug disables the debug log mode and has higher priority than DEBUG environment variable.
func (l *Logger) DisableDebug() {
	atomic.StoreInt32(&l.isDebug, 0)
}

// Debug prints with "[DEBUG] " prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Debug(v ...any) {
	if atomic.LoadInt32(&l.isDebug) == 1 {
		levelText := "DEBUG"
		if isatty.IsTerminal(os.Stdout.Fd()) {
			levelText = colorize(gray, levelText)
		}
		l.Print(fmt.Sprintf("[%s] ", levelText), fmt.Sprint(v...))
	}
}

// Debugf prints with "[DEBUG] " prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, v ...any) {
	l.Debug(fmt.Sprintf(format, v...))
}

// Debugln prints with "[DEBUG] " prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Debugln(v ...any) {
	l.Debug(fmt.Sprintln(v...))
}

func colorize(color int, text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, text)
}

// EnableDebug enables the debug log mode and has higher priority than DEBUG environment variable.
func EnableDebug() {
	std.EnableDebug()
}

// DisableDebug disables the debug log mode and has higher priority than DEBUG environment variable.
func DisableDebug() {
	std.DisableDebug()
}

// Debug prints with "[DEBUG] " prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Print.
func Debug(v ...any) {
	std.Debug(v...)
}

// Debugf prints with "[DEBUG] " prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...any) {
	std.Debugf(format, v...)
}

// Debugln prints with "[DEBUG] " prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Println.
func Debugln(v ...any) {
	std.Debugln(v...)
}
