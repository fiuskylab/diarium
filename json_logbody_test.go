package diarium

import (
	"regexp"
	"testing"
)

func TestLogBody(t *testing.T) {

	var tts []testCases

	logBody := newLogBody(Warning.toString(), "Foo bar", nil)

	tts = append(tts, testCases{
		name:     "Level",
		want:     Warning.toString(),
		got:      logBody.Level,
		testType: Equal,
	})

	tts = append(tts, testCases{
		name:     "Message",
		want:     "Foo bar",
		got:      logBody.Message,
		testType: Equal,
	})

	tts = append(tts, testCases{
		name:     "Nil Extra",
		got:      logBody.Extra,
		testType: Nil,
	})

	tts = append(tts, testCases{
		name:      "Timestamp",
		got:       logBody.Timestamp,
		regexRule: regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})-(\d{2}):(\d{2})`),
		testType:  Regex,
	})

	runTests(t, tts)
}
