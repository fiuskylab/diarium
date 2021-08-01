package diarium

import "testing"

func getClientTests() (tts []testCases) {
	{
		tts = append(tts, testCases{
			name: "NewDefaultClient",
			want: &Client{},
			got:  NewDefaultClient(),
		})
	}
	return tts
}

func TestClient(t *testing.T) {
	runTests(t, getClientTests())
}
