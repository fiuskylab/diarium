package diarium

import (
	"regexp"
	"testing"
)

func TestNetwork(t *testing.T) {
	tts := []testCases{}

	{
		network := getNetwork()
		tts = append(tts, testCases{
			name:      "Network IP Regex",
			got:       network.IP,
			regexRule: regexp.MustCompile(`(\d{1,3}).(\d{1,3}).(\d{1,3}).(\d{1,3})`),
			testType:  Regex,
		})

		tts = append(tts, testCases{
			name:      "Network Hostname Regex",
			got:       network.IP,
			regexRule: regexp.MustCompile(`(.{1,})`),
			testType:  Regex,
		})
	}

	{
		ip := getIP()
		tts = append(tts, testCases{
			name:      "getIP",
			got:       ip,
			regexRule: regexp.MustCompile(`(\d{1,3}).(\d{1,3}).(\d{1,3}).(\d{1,3})`),
			testType:  Regex,
		})
	}

	runTests(t, tts)
}
