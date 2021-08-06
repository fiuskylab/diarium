package diarium

import (
	"os"
	"testing"
)

func getClientTests() (tts []testCases) {
	{
		tts = append(tts, testCases{
			name:     "NewDefaultClient",
			want:     &Client{},
			got:      NewDefaultClient(),
			testType: Equal,
		})
	}
	{
		cfg := NewDefaultConfig()
		tts = append(tts, testCases{
			name:     "NewClient",
			want:     &Client{cfg: cfg},
			got:      NewClient(cfg),
			testType: Equal,
		})
	}
	{
		w, _ := os.OpenFile("tmp/client_test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		defer w.Close()
		f := newFile(w)
		client := NewClient(NewDefaultConfig().addOutput(f))
		tts = append(tts, testCases{
			name:     "NewClient",
			want:     nil,
			got:      client.Log(testInterface),
			testType: Equal,
		})
	}
	{
		w, _ := os.OpenFile("tmp/client_test_closed", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		w.Close()
		f := newFile(w)
		client := NewClient(NewDefaultConfig().addOutput(f))
		tts = append(tts, testCases{
			name:     "NewClient",
			want:     "write tmp/client_test_closed: file already closed",
			got:      client.Log(testInterface).Error(),
			testType: Equal,
		})
	}
	return tts
}

func TestClient(t *testing.T) {
	runTests(t, getClientTests())
}
