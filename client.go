package diarium

// Client is the main part of the logging
// tool. It's responsible for every public info and action
// related to it.
type Client struct {
	cfg *Config
}

// NewDefaultClient return a Client
// with default configs
func NewDefaultClient() *Client {
	return &Client{
		cfg: NewDefaultConfig(),
	}
}

// NewClient return a Client
// with given Config
func NewClient(cfg *Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

// Log prints a given interface through all
// configured outputs
func (c *Client) Log(lvl, message string, extra map[string]string) error {
	for _, o := range c.cfg.Outputs {
		err := o.output(newLogBody(lvl, message, extra))
		if err != nil {
			return err
		}
	}
	return nil
}
