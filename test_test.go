package diarium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCases struct {
	name string
	want interface{}
	got  interface{}
}

func runTests(t *testing.T, tts []testCases) {
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.got)
		})
	}
}
