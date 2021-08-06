package diarium

import (
	"time"
)

// LogBody
type LogBody struct {
	Level     string            `json:"level"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	Trace     trace             `json:"trace"`
	Network   network           `json:"network"`
	Extra     map[string]string `json:"extra"`
}

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
