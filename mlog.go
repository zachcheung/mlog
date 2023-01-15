package mlog

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	gray = 37
)

// A Logger is a standard `log.Logger` with `Debug` family functions.
type Logger struct {
	log.Logger
	isDebug bool
}

// New creates a new Logger and enables the debug log mode when found
// `DEBUG` environment variable set to one of the following values:
//   - yes
//   - true
//   - 1
func New() *Logger {
	var isDebug bool
	switch strings.ToLower(os.Getenv("DEBUG")) {
	case "yes", "true", "1":
		isDebug = true
	}
	return &Logger{
		Logger:  *log.Default(),
		isDebug: isDebug,
	}
}

var std = New()

// EnableDebug enables the debug log mode and has higher priority than `DEBUG` environment variable.
func (l *Logger) EnableDebug() {
	l.isDebug = true
}

// DisableDebug disables the debug log mode and has higher priority than `DEBUG` environment variable.
func (l *Logger) DisableDebug() {
	l.isDebug = false
}

// Debug prints with `[DEBUG] ` prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Debug(v ...any) {
	if l.isDebug {
		l.Print(fmt.Sprintf("[%s] ", colorize(gray, "DEBUG")), fmt.Sprint(v...))
	}
}

// Debugf prints with `[DEBUG] ` prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, v ...any) {
	l.Debug(fmt.Sprintf(format, v...))
}

// Debugln prints with `[DEBUG] ` prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Debugln(v ...any) {
	l.Debug(fmt.Sprintln(v...))
}

func colorize(color int, text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, text)
}

// EnableDebug enables the debug log mode and has higher priority than `DEBUG` environment variable.
func EnableDebug() {
	std.EnableDebug()
}

// DisableDebug disables the debug log mode and has higher priority than `DEBUG` environment variable.
func DisableDebug() {
	std.DisableDebug()
}

// Debug prints with `[DEBUG] ` prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Print.
func Debug(v ...any) {
	std.Debug(v...)
}

// Debugf prints with `[DEBUG] ` prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...any) {
	std.Debugf(format, v...)
}

// Debugln prints with `[DEBUG] ` prefix when in the debug mode.
// Arguments are handled in the manner of fmt.Println.
func Debugln(v ...any) {
	std.Debugln(v...)
}
