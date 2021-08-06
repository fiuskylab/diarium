package diarium

// Client
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
func (c *Client) Log(i LogBody) error {
	for _, o := range c.cfg.Outputs {
		err := o.output(i)
		if err != nil {
			return err
		}
	}
	return nil
}
