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
				Outputs: []Output{
					newTerminal(),
				},
			},
			got:      NewDefaultConfig(),
			testType: Equal,
		})
	}

	{
		tts = append(tts, testCases{
			name:     "NewConfig",
			want:     &Config{},
			got:      NewConfig(),
			testType: Equal,
		})
	}

	{
		tts = append(tts, testCases{
			name: "Config.setLevels",
			want: &Config{
				AllowedLevels: []level{Emergency, Alert},
			},
			got:      NewConfig().setLevels(Emergency, Alert),
			testType: Equal,
		})
	}

	{
		tts = append(tts, testCases{
			name: "Config.addOutput",
			want: &Config{
				Outputs: []Output{Output(newTerminal())},
			},
			got:      NewConfig().addOutput(Output(newTerminal())),
			testType: Equal,
		})
	}
	return tts
}

func TestConfig(t *testing.T) {
	runTests(t, getConfigTests())
}
