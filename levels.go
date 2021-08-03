package diarium

type level int

const (
	// Emergency - Unusable conditions
	Emergency level = iota
	// Alert - A condition that a action
	// must be taken immediatly
	Alert
	// Critical - Critical condition
	Critical
	// Error - A error were handled, must
	// investigate if it's a common behavior
	// or not
	Error
	// Warning - Warning, something happened
	// but the system continued running normally
	Warning
	// Notice - Normal messages, but with
	// significance
	Notice
	// Info - Informational messages, not
	// related to errors
	Info
	// Debug - For debbuging purposes
	Debug
)

var mapLevel = map[level]string{
	Emergency: "Emergency",
	Critical:  "Critical",
	Alert:     "Alert",
	Error:     "Error",
	Warning:   "Warning",
	Notice:    "Notice",
	Info:      "Info",
	Debug:     "Debug",
}

// Enum manages the converting functions
// and others, related to ours consts
type enum interface {
	toString() string
	toIndex() int
}

// String return Level as string
func (l level) toString() string {
	return mapLevel[l]
}

// Index return Level's index(int) value
func (l level) toIndex() int {
	return int(l)
}

// getString retrives Enum's string value
func getString(e enum) string {
	return e.toString()
}

// getString retrives Enum's int value
func getIndex(e enum) int {
	return e.toIndex()
}
