package diarium

import "runtime"

type trace struct {
	Filename  string `json:"filename"`
	Line      int    `json:"line"`
	ScopeName string `json:"scope_name"`
}

// TRACESKIP - a
const TRACESKIP = 2

func getTrace() trace {
	pc := make([]uintptr, 1)

	runtime.Callers(TRACESKIP, pc)

	f := runtime.FuncForPC(pc[0])

	_, filename, line, _ := runtime.Caller(TRACESKIP)

	return trace{
		ScopeName: f.Name(),
		Filename:  filename,
		Line:      line,
	}
}
