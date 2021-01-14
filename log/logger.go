package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type LevelCode int32

// Defining all the log levels and assigning them
// a unique digit which can be used for enabling
// the log level
const (
	DebugLevel string = "debug"
	InfoLevel string = "info"
	WarningLevel string = "warn"
	ErrorLevel string = "error"
	PanicLevel string = "panic"
	DebugCode LevelCode = 5
	InfoCode LevelCode = 4
	WarningCode LevelCode = 3
	PanicCode LevelCode = 2
	ErrorCode LevelCode = 1
)

var logLevel = InfoCode

func init() {
	level := parseLevel(os.Getenv("WABA_LOG_LEVEL"))
	setLevel(level)
}

func parseLevel(level string) LevelCode {
	switch strings.ToLower(level) {
	case DebugLevel:
		return DebugCode
	case InfoLevel:
		return InfoCode
	case WarningLevel:
		return WarningCode
	case ErrorLevel:
		return ErrorCode
	case PanicLevel:
		return PanicCode
	default:
		return InfoCode
	}
}

// setLevel changes the level in global variable
// default log level is INFO
func setLevel(code LevelCode) {
	logLevel = code
}

// checkLevel performs basic math to check which
// log level is enabled, the idea is to check for
// positive numbers like if the log level warn
// is enabled then it should execute warn, error
// and panic levels
func checkLevel(level LevelCode) bool {
	if level - logLevel <= 0 {
		return true
	}
	return false
}

func processMsg(msg string, msgs ...interface{}) string {
	for _, v := range msgs {
		msg += fmt.Sprintf("%v ", v)
	}
	return msg
}

func Info(msgs ...interface{}) {
	if !checkLevel(InfoCode){
		return
	}
	fmt.Println(msgs ...)
}

func Warn(msgs ...interface{}) {
	if !checkLevel(WarningCode){
		return
	}
	fmt.Println(processMsg("Warning!! ", msgs ...))
}

//
func Debug(msgs ...interface{}) {
	if !checkLevel(DebugCode){
		return
	}
	t := time.Now()
	cur := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println(processMsg("[" + cur + "] debug: ", msgs ...))
}

// Error should always print to screen and exit after that
func Error(msgs ...interface{}) {
	fmt.Println(msgs ...)
	os.Exit(1)
}

// Panic can be enabled for debugging and development purposes
func Panic(msgs ...interface{}) {
	if !checkLevel(PanicCode){
		return
	}
	panic(processMsg("", msgs ...))
}

