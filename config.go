package diarium

import (
	"github.com/fiuskylab/diarium/consts"
	"github.com/fiuskylab/diarium/outputs"
)

// Config provides the main settings for
// diarium logging tool
type Config struct {
	// AllowedLevels store all type of logging
	// that the client is allowed to output
	AllowedLevels []consts.Level
	// Outputs store all configured output options
	// available to Diarium's clients
	Outputs []outputs.Output
}

// NewDefaultConfig returns a Config
// with default values.
// By default the config comes with:
// 		- All levels allowed
// 		- 1 Output (terminal)
func NewDefaultConfig() *Config {
	return &Config{
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
	}
}

// NewConfig return an empty config struct
func NewConfig() *Config {
	return &Config{}
}

// setLevels overwrites AllowedLevels field
func (c *Config) setLevels(lvls ...consts.Level) *Config {
	c.AllowedLevels = lvls
	return c
}

// addOutput append a new output into config Outputs
func (c *Config) addOutput(o outputs.Output) *Config {
	c.Outputs = append(c.Outputs, o)
	return c
}
