package diarium

import (
	"time"
)

// LogBody stores all data related
// to the logging
type LogBody struct {
	// Level is which level was logged
	Level string `json:"level"`
	// Message store any message that is important
	// for later debugging, e.g err.Error()
	Message string `json:"message"`
	// Timestamp is when it was logged
	Timestamp string `json:"timestamp"`
	// Trace store file information related
	// to the log execution
	Trace trace `json:"trace"`
	// Network store network information related
	// to the log execution
	Network network `json:"network"`
	// Extra is any extra info that might be
	// helpful
	Extra map[string]string `json:"extra"`
}

// newLogBody generates a log with filled fields
func newLogBody(lvl, message string, extra map[string]string) LogBody {
	return LogBody{
		Level:     lvl,
		Message:   message,
		Extra:     extra,
		Timestamp: time.Now().Format(time.RFC3339),
		Trace:     getTrace(),
		Network:   getNetwork(),
	}
}
