package diarium

// Config provides the main settings for
// diarium logging tool
type Config struct {
	// AllowedLevels store all type of logging
	// that the client is allowed to output
	AllowedLevels []level
	// Outputs store all configured output options
	// available to Diarium's clients
	Outputs []Output
}

// NewDefaultConfig returns a Config
// with default values.
// By default the config comes with:
// 		- All levels allowed
// 		- 1 Output (terminal)
func NewDefaultConfig() *Config {
	return &Config{
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
	}
}

// NewConfig return an empty config struct
func NewConfig() *Config {
	return &Config{}
}

// setLevels overwrites AllowedLevels field
func (c *Config) setLevels(lvls ...level) *Config {
	c.AllowedLevels = lvls
	return c
}

// addOutput append a new output into config Outputs
func (c *Config) addOutput(o Output) *Config {
	c.Outputs = append(c.Outputs, o)
	return c
}
