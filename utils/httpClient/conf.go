package httpClient

import "time"

type Client struct {
	Service string        `yaml:"service"`
	Domain  string        `yaml:"domain"`
	Timeout time.Duration `yaml:"timeout"`
	Retry   int           `yaml:"retry"`
}
