package diarium

import "testing"

func getTerminalTests() (tts []testCases) {
	testInterface := struct {
		Name string `json:"name"`
	}{
		Name: "Foo",
	}

	{
		tts = append(tts, testCases{
			name:     "newTerminal",
			want:     Output(&terminal{}),
			got:      newTerminal(),
			testType: Equal,
		})
	}

	{
		t := newTerminal()
		tts = append(tts, testCases{
			name:     "output",
			got:      t.output(testInterface),
			testType: Nil,
		})
	}
	return tts
}

func TestTerminal(t *testing.T) {
	runTests(t, getTerminalTests())
}
