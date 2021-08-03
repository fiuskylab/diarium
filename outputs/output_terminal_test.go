package outputs

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
			got:  NewTerminal(),
		})
	}

	{
		t := NewTerminal()
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
