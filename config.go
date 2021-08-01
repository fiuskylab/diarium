package diarium

// Config provides the main settings for
// diarium logging tool
type Config struct {
	AllowedLevels []level
}

// NewDefaultConfig returns a Config
// with default values
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
	}
}
