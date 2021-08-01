package diarium

import "testing"

func getConfigTests() (tts []testCases) {
	{
		tts = append(tts, testCases{
			name: "NewDefaultConfig",
			want: &Config{
				AllowedLevels: []level{
					Emergency,
					Critical,
					Alert,
					Error,
					Warning,
					Notice,
					Info,
					Debug,
				},
			},
			got: NewDefaultConfig(),
		})
	}
	return tts
}

func TestConfig(t *testing.T) {
	runTests(t, getConfigTests())
}
