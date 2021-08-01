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

func (f *file) output(i interface{}) error {
	b, _ := json.Marshal(i)
	_, err := fmt.Fprintln(f.w, string(b))
	return err
}

func newFile(w *os.File) Output {
	return Output(&file{
		w: w,
	})
}
