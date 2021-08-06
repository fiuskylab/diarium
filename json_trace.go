package diarium

import "runtime"

type trace struct {
	Filename  string `json:"filename"`
	Line      int    `json:"line"`
	ScopeName string `json:"scope_name"`
}

// TRACERUNTIME - a
const TRACERUNTIME = 2

// TRACESCOPE
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
