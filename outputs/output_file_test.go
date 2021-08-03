package outputs

import (
	"os"
	"testing"
)

func getFileTests() (tts []testCases) {
	{
		w, _ := os.OpenFile("tmp/output_file_test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		defer w.Close()
		tts = append(tts, testCases{
			name: "newFile",
			want: Output(&file{
				w: w,
			}),
			got: NewFile(w),
		})
	}

	{
		w, _ := os.OpenFile("tmp/output_file_test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		defer w.Close()
		l := NewFile(w)
		tts = append(tts, testCases{
			name: "output",
			want: nil,
			got:  l.output(testInterface),
		})
	}

	return tts
}

func TestFile(t *testing.T) {
	runTests(t, getFileTests())
}
