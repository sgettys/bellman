package hub

import "github.com/sgettys/bellman/pkg/criers"

// Config allows configuration
type Config struct {
	LogLevel  string                  `yaml:"loglevel"`
	LogFormat string                  `yaml:"logformat"`
	Receivers []criers.ReceiverConfig `yaml:"receivers"`
}

func (c *Config) Validate() error {
	return nil
}
