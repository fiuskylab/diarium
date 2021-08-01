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
			got: NewDefaultConfig(),
		})
	}

	{
		tts = append(tts, testCases{
			name: "NewConfig",
			want: &Config{},
			got:  NewConfig(),
		})
	}

	{
		tts = append(tts, testCases{
			name: "Config.setLevels",
			want: &Config{
				AllowedLevels: []level{Emergency, Alert},
			},
			got: NewConfig().setLevels(Emergency, Alert),
		})
	}

	{
		tts = append(tts, testCases{
			name: "Config.addOutput",
			want: &Config{
				Outputs: []Output{Output(newTerminal())},
			},
			got: NewConfig().addOutput(Output(newTerminal())),
		})
	}
	return tts
}

func TestConfig(t *testing.T) {
	runTests(t, getConfigTests())
}
