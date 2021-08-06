package diarium

import "testing"

func getLevelsTest() (tts []testCases) {
	{
		tts = append(tts, testCases{
			name:     "level.toString",
			want:     "Emergency",
			got:      Emergency.toString(),
			testType: Equal,
		})
	}

	{
		tts = append(tts, testCases{
			name:     "level.toIndex",
			want:     0,
			got:      Emergency.toIndex(),
			testType: Equal,
		})
	}

	{
		tts = append(tts, testCases{
			name:     "enum.getString",
			want:     "Emergency",
			got:      getString(Emergency),
			testType: Equal,
		})
	}

	{
		tts = append(tts, testCases{
			name:     "enum.getString",
			want:     0,
			got:      getIndex(Emergency),
			testType: Equal,
		})
	}
	return tts
}

func TestLevels(t *testing.T) {
	runTests(t, getLevelsTest())
}
