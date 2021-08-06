package diarium

import (
	"regexp"
	"testing"
)

func FooBar() trace {
	return getTrace()
}

func TestTrace(t *testing.T) {

	tts := []testCases{}

	trace := FooBar()

	{
		tts = append(tts, testCases{
			name:      "Filename",
			got:       trace.Filename,
			regexRule: regexp.MustCompile(`(.{1,}).go`),
			testType:  Regex,
		})
	}

	{
		tts = append(tts, testCases{
			name:     "Line",
			got:      trace.Line,
			want:     19,
			testType: Equal,
		})
	}

	{
		tts = append(tts, testCases{
			name:      "Scope",
			got:       trace.ScopeName,
			regexRule: regexp.MustCompile(`(.{0,})(FooBar)(.{0,})`),
			testType:  Regex,
		})
	}

	runTests(t, tts)
}
