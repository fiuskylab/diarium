package diarium

import "testing"

func getTerminalTests() (tts []testCases) {
	testInterface := struct {
		name string `json:name`
	}{
		name: "Foo",
	}

	{
		tts = append(tts, testCases{
			name: "newTerminal",
			want: Output(&terminal{}),
			got:  newTerminal(),
		})
	}

	{
		t := newTerminal()
		tts = append(tts, testCases{
			name: "output",
			want: nil,
			got:  t.output(testInterface),
		})
	}
	return tts
}

func TestTerminal(t *testing.T) {
	runTests(t, getTerminalTests())
}
