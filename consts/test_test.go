package consts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	name string
	want interface{}
	got  interface{}
}

var TestInterface = struct {
	name string `json:name`
}{
	name: "Foo",
}

func runTests(t *testing.T, tts []TestCases) {
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.got)
		})
	}
}
