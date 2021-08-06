package diarium

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
			got:      newFile(w),
			testType: Equal,
		})
	}

	{
		w, _ := os.OpenFile("tmp/output_file_test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		defer w.Close()
		l := newFile(w)
		tts = append(tts, testCases{
			name:     "output",
			got:      l.output(testInterface),
			testType: Nil,
		})
	}

	return tts
}

func TestFile(t *testing.T) {
	runTests(t, getFileTests())
}
