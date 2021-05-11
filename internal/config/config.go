package config

import (
	"github.com/caarlos0/env/v6"
)

// Config represents a Sojourner config object
var Config struct {
	HostName         string  `env:"SOJOURNER_HOSTNAME"`
	Port             string  `env:"SOJOURNER_PORT"`
	DatabaseLocation string  `env:"SOJOURNER_DB_LOCATION" envDefault:"/tmp/badger"`
	IDLength         int     `env:"SOJOURNER_ID_LENGTH" envDefault:"8"`
	MaxPayloadSize   float64 `env:"SOJOURNER_MAX_PAYLOAD_SIZE" envDefault:"256"`
}

// Load reads variables from the environment and parses them into Config
func Load() error {
	return env.Parse(&Config)
}
