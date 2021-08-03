package diarium

import (
	"testing"

	"github.com/fiuskylab/diarium/consts"
	"github.com/fiuskylab/diarium/outputs"
)

func getConfigTests() (tts []testCases) {
	{
		tts = append(tts, testCases{
			name: "NewDefaultConfig",
			want: &Config{
				AllowedLevels: []consts.Level{
					consts.Emergency,
					consts.Critical,
					consts.Alert,
					consts.Error,
					consts.Warning,
					consts.Notice,
					consts.Info,
					consts.Debug,
				},
				Outputs: []outputs.Output{
					outputs.NewTerminal(),
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
				AllowedLevels: []consts.Level{consts.Emergency, consts.Alert},
			},
			got: NewConfig().setLevels(consts.Emergency, consts.Alert),
		})
	}

	{
		tts = append(tts, testCases{
			name: "Config.addOutput",
			want: &Config{
				Outputs: []outputs.Output{outputs.Output(outputs.NewTerminal())},
			},
			got: NewConfig().addOutput(outputs.Output(outputs.NewTerminal())),
		})
	}
	return tts
}

func TestConfig(t *testing.T) {
	runTests(t, getConfigTests())
}
