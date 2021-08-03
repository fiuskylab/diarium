package diarium

import (
	"os"
	"testing"

	"github.com/fiuskylab/diarium/outputs"
)

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
	{
		w, _ := os.OpenFile("tmp/client_test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		defer w.Close()
		f := outputs.NewFile(w)
		client := NewClient(NewDefaultConfig().addOutput(f))
		tts = append(tts, testCases{
			name: "NewClient",
			want: nil,
			got:  client.Log(testInterface),
		})
	}
	{
		w, _ := os.OpenFile("tmp/client_test_closed", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		w.Close()
		f := outputs.NewFile(w)
		client := NewClient(NewDefaultConfig().addOutput(f))
		tts = append(tts, testCases{
			name: "NewClient",
			want: "write tmp/client_test_closed: file already closed",
			got:  client.Log(testInterface).Error(),
		})
	}
	return tts
}

func TestClient(t *testing.T) {
	runTests(t, getClientTests())
}
