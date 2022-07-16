package config

import "flag"

type config struct {
	UseBuffer bool `json:"buffer"`
	Index     int  `json:"index"`
}

// ParseFlags reads flags and sets values to the configuration properties.
func (c *config) ParseFlags() *config {
	flag.Parse()
	return c
}

// New creates new object of the app configuration.
func New() *config {
	var cfg = &config{
		UseBuffer: true,
		Index:     0,
	}

	flag.BoolVar(&cfg.UseBuffer, "buffer", true, "Using buffer is enabled for calculation")
	flag.IntVar(&cfg.Index, "index", 0, "The index of the number in the Fibonacci sequence.")

	return cfg.ParseFlags()
}
