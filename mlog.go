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

// These functions write to the standard logger and from go1.19.5:src/log/log.go.

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...any) {
	std.Print(v...)
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...any) {
	std.Printf(format, v...)
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...any) {
	std.Println(v...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...any) {
	std.Fatal(v...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...any) {
	std.Fatalf(format, v...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...any) {
	std.Fatalln(v...)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...any) {
	std.Panic(v...)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...any) {
	std.Panicf(format, v...)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...any) {
	std.Panicln(v...)
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is the count of the number of
// frames to skip when computing the file name and line number
// if Llongfile or Lshortfile is set; a value of 1 will print the details
// for the caller of Output.
func Output(calldepth int, s string) error {
	return std.Output(calldepth+1, s) // +1 for this frame.
}
