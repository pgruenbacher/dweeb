// Author: shorrockin/noted
// log is a simple abstraction to provide common logging packages
// with colored log output.
package log

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

const blue = "\x1B[0;34m"
const red = "\x1B[0;31m"
const green = "\x1B[0;32m"
const yellow = "\x1B[0;33m"
const white = ""
const reset = "\x1B[0m"
const teal = "\x1B[0;36m"
const gray = "\x1B[0;35m"

type LogLevel struct {
	color  string
	level  int
	name   string
	suffix string
	method func(string, ...interface{})
}

func init() {
	// only prints the log.go file lines...need to fix
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetFlags(0)
}

func (level LogLevel) prefix() string {
	return fmt.Sprintf("%v[%v%v%v] ", level.suffix, level.color, level.name, reset)
}

func (level LogLevel) log(format string, v ...interface{}) {
	fname, fline := trace()
	p := fmt.Sprintf("%v%10s:%d%v ", gray, fname, fline, reset)
	if level.level >= MinLevel.level {
		level.method(level.prefix()+p+format, v...)
	}
}

func (level LogLevel) String() string {
	return fmt.Sprintf("LogLevel[name: %v%v%v, level: %v]", level.color, level.name, reset, level.level)
}

var DebugLevel = LogLevel{blue, 0, "debug", "", log.Printf}
var InfoLevel = LogLevel{teal, 1, "info", " ", log.Printf}
var WarnLevel = LogLevel{yellow, 2, "warn", " ", log.Printf}
var ErrorLevel = LogLevel{red, 3, "error", "", log.Printf}
var PanicLevel = LogLevel{red, 4, "panic", "", log.Panicf}
var FatalLevel = LogLevel{red, 5, "fatal", "", log.Fatalf}

var MinLevel = DebugLevel

// Temp prints a string with all the colors of the rainbow so that
// you can easily identify them in logs.
func Temp(format string, v ...interface{}) {
	tmp := fmt.Sprintf("[%v!%v%v!%v%v!%v%v!%v%v!%v] ", red, reset, yellow, reset, blue, reset, green, reset, red, reset)
	log.Printf(tmp+format, v...)
}

// Debug prints to the output with the string 'debug' colored blue
func Debug(format string, v ...interface{}) {
	DebugLevel.log(format, v...)
}

// Info prints to the output with the string 'info' colored white
func Info(format string, v ...interface{}) {
	InfoLevel.log(format, v...)
}

// Warn prints to the output with the string 'warn' colored yellow
func Warn(format string, v ...interface{}) {
	WarnLevel.log(format, v...)
}

// Error prints to the output with the string 'error' colored red
func Error(format string, v ...interface{}) {
	ErrorLevel.log(format, v...)
}

// Panic prints to the output with the string 'panic' colored red
// followed by a call to panic.
func Panic(format string, v ...interface{}) {
	PanicLevel.log(format, v...)
}

// Fatal prints to the output with the string 'fatal' colored red
// followed by a call to os.Exit
func Fatal(format string, v ...interface{}) {
	FatalLevel.log(format, v...)
}

// FuncForPC returns a *Func describing the function that contains the given program counter address, or else nil.
// FileLine returns the file name and line number of the source code corresponding to the program counter pc. The result will not be accurate if pc is not a program counter within f.
func trace() (string, int) {
	pc := make([]uintptr, 3) // at least 1 entry needed.
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[1])
	file, line := f.FileLine(pc[1])

	nameString := file
	nameArray := strings.Split(nameString, "/")
	name := nameArray[len(nameArray)-1]

	return name, line
}
