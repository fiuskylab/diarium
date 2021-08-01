package diarium

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
