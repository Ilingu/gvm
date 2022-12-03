package console

import (
	"fmt"
	"log"
)

type MsgType string

const (
	reset = "\x1b[0m"

	// bright     = "\x1b[1m"
	// dim        = "\x1b[2m"

	underscore = "\x1b[4m"
	// blink      = "\x1b[5m"
	// reverse    = "\x1b[7m"
	// hidden     = "\x1b[8m"

	// Foreground Colors

	fgBlack   = "\x1b[30m"
	fgRed     = "\x1b[31m"
	fgGreen   = "\x1b[32m"
	fgYellow  = "\x1b[33m"
	fgBlue    = "\x1b[34m"
	fgMagenta = "\x1b[35m"
	fgCyan    = "\x1b[36m"
	fgWhite   = "\x1b[37m"

	// Background Colors
	// bgBlack   = "\x1b[40m"
	// bgRed     = "\x1b[41m"
	// bgGreen   = "\x1b[42m"
	// bgYellow  = "\x1b[43m"
	// bgBlue    = "\x1b[44m"
	// bgMagenta = "\x1b[45m"
	// bgCyan    = "\x1b[46m"
	// bgWhite   = "\x1b[47m"

	// Type

	INFO    MsgType = fgCyan
	SUCCESS MsgType = fgGreen
	WARNING MsgType = fgYellow
	ERROR   MsgType = fgRed
	NEUTRAL MsgType = fgWhite
)

var MsgTypeToString = map[MsgType]string{INFO: "Info", SUCCESS: "Success", WARNING: "Warning", ERROR: "Error", NEUTRAL: "Neutral"}

// It logs a message to the standar output with colors and flags
//
// Usage:
//
//	LogMsg("couldn't open a connection to db", ERROR)
//
// The 3rd argument is facultative, it's whether you want to underline the message output or not
func LogMsg(msg any, priority MsgType, underline ...bool) {
	msgType := NEUTRAL
	if isOkColor(priority) {
		msgType = priority
	}

	MsgTypeText := []any{string(fgMagenta), fmt.Sprintf("[%s]", MsgTypeToString[msgType]), string(reset)}
	Message := []any{string(msgType), msg, string(reset)}
	if len(underline) == 1 && underline[0] {
		Message = append([]any{string(underscore)}, Message...)
	}

	LogMessage := append(MsgTypeText, Message...)
	log.Println(LogMessage...)
}

// Short and Handy Version of LogMsg() to log Info messages
//
// Under the hood:
//
//	LogMsg(your_msg, INFO, underline...)
func Log(msg any, underline ...bool) {
	LogMsg(msg, INFO, underline...)
}

// Short and Handy Version of LogMsg() to log Neutral messages
//
// Under the hood:
//
//	LogMsg(your_msg, NEUTRAL, underline...)
func Neutral(msg any, underline ...bool) {
	LogMsg(msg, NEUTRAL, underline...)
}

// Short and Handy Version of LogMsg() to log Success messages
//
// Under the hood:
//
//	LogMsg(your_msg, SUCCESS, underline...)
func Success(msg any, underline ...bool) {
	LogMsg(msg, SUCCESS, underline...)
}

// Short and Handy Version of LogMsg() to log Warning messages
//
// Under the hood:
//
//	LogMsg(your_msg, WARNING, underline...)
func Warn(msg any, underline ...bool) {
	LogMsg(msg, WARNING, underline...)
}

// Short and Handy Version of LogMsg() to log Error messages
//
// Under the hood:
//
//	LogMsg(your_msg, ERROR, underline...)
func Error(msg any, underline ...bool) {
	LogMsg(msg, ERROR, underline...)
}

func isOkColor(priority MsgType) bool {
	switch priority {
	case INFO, SUCCESS, WARNING, ERROR, NEUTRAL:
		return true
	default:
		return false
	}
}
