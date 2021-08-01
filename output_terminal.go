package diarium

import (
	"encoding/json"
	"fmt"
)

// terminal writes the logging into terminal's
type terminal struct {
}

// output prints given log into terminal's buffer
func (f *terminal) output(i interface{}) error {
	b, _ := json.Marshal(i)
	fmt.Println(string(b))

	return nil
}

// newTerminal add the output option to print
// logs into terminal's buffer
func newTerminal() Output {
	return Output(&terminal{})
}
