package diarium

import (
	"encoding/json"
	"fmt"
	"os"
)

// file writes the logging into a given file
type file struct {
	w *os.File
}

// output writes given log into file
func (f *file) output(i LogBody) error {
	b, _ := json.Marshal(i)
	_, err := fmt.Fprintln(f.w, string(b))
	return err
}

// newTerminal add the output option to print
// logs into given file
func newFile(w *os.File) Output {
	return Output(&file{
		w: w,
	})
}
