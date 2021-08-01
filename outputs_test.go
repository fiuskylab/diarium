package diarium

import (
	"os"
	"testing"
)

func getOutputTests() (tts []testCases) {
	testInterface := struct {
		name string `json:name`
	}{
		name: "Foo",
	}

	w, _ := os.OpenFile("tmp/outputs_test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	f := newFile(w)
	{
		tts = append(tts, testCases{
			name: "output",
			want: nil,
			got:  output(f, testInterface),
		})
	}

	return tts
}

func TestOutput(t *testing.T) {
	runTests(t, getOutputTests())
}