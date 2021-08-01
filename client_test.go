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
	{
		cfg := NewDefaultConfig()
		tts = append(tts, testCases{
			name: "NewClient",
			want: &Client{cfg: cfg},
			got:  NewClient(cfg),
		})
	}
	return tts
}

func TestClient(t *testing.T) {
	runTests(t, getClientTests())
}
