package config

import (
	"github.com/caarlos0/env/v6"
)

// Config represents a BBin config object
var Config struct {
	HostName         string  `env:"BBIN_HOSTNAME"`
	Port             string  `env:"BBIN_PORT"`
	DatabaseLocation string  `env:"BBIN_DB_LOCATION" envDefault:"/tmp/badger"`
	IDLength         int     `env:"BBIN_ID_LENGTH" envDefault:"8"`
	MaxPayloadSize   float64 `env:"BBIN_MAX_PAYLOAD_SIZE" envDefault:"256"`
}

// Load reads variables from the environment and parses them into Config
func Load() error {
	return env.Parse(&Config)
}
