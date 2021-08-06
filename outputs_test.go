package diarium

import (
	"os"
	"testing"
)

func getOutputTests() (tts []testCases) {
	testInterface := struct {
		Name string `json:"name"`
	}{
		Name: "Foo",
	}

	w, _ := os.OpenFile("tmp/outputs_test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	f := newFile(w)
	{
		tts = append(tts, testCases{
			name:     "output",
			got:      output(f, testInterface),
			testType: Nil,
		})
	}

	return tts
}

func TestOutput(t *testing.T) {
	runTests(t, getOutputTests())
}
