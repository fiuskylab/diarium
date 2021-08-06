package diarium

import "runtime"

// trace stores the app execution data
// related to files
type trace struct {
	// Filename the name of the file
	Filename string `json:"filename"`
	// Line in which line the log was called
	Line int `json:"line"`
	// ScopeName name of the function that
	// the log was called
	ScopeName string `json:"scope_name"`
}

// TRACERUNTIME - Define the trace depth
// into runtime.Callers
const TRACERUNTIME = 2

// TRACESCOPE - Define the trace depth
// into runtime.Callers
const TRACESCOPE = 0

func getTrace() trace {
	pc := make([]uintptr, 1)

	runtime.Callers(TRACESCOPE, pc)

	f := runtime.FuncForPC(pc[0])

	_, filename, line, _ := runtime.Caller(TRACERUNTIME)

	return trace{
		ScopeName: f.Name(),
		Filename:  filename,
		Line:      line,
	}
}
