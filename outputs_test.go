package diarium

import (
	"os"
	"testing"
)

func getOutputTests() (tts []testCases) {
	testLogBody := newLogBody(Debug.toString(), "Foo bar", nil)

	w, _ := os.OpenFile("tmp/outputs_test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	f := newFile(w)
	{
		tts = append(tts, testCases{
			name:     "output",
			got:      output(f, testLogBody),
			testType: Nil,
		})
	}

	return tts
}

func TestOutput(t *testing.T) {
	runTests(t, getOutputTests())
}
