package cli

import "github.com/rs/zerolog"

type LogFormat string

const (
	// JSONFormat prints messages as single json record
	JSONFormat LogFormat = "json"
	// SimpleFormat prints messages in human readable way
	SimpleFormat LogFormat = "simple"
	// NoFormat returned in situation when format is not recognised
	NoFormat LogFormat = "noformat"
)

type Config struct {
	LocalServerPath string `optional:"" env:"LOCAL_SERVER_PATH"`
	LocaLServerPort string `optional:"" env:"LOCAL_SERVER_PORT" default:"8090"`
}

func (c Config) IsLocalServerAllowed() bool {
	return c.LocalServerPath != "" && c.LocaLServerPort != ""
}

type Log struct {
	// Level [panic, fatal, error,warn,info,debug,trace], defines level of logger, default: info
	Level zerolog.Level `env:"LOG_LEVEL" enum:"debug,info,warn,error,trace" default:"info"`

	// Format [simple,json] specifies how the logger prints values
	Format LogFormat `env:"LOG_FORMAT" enum:"simple,json,noformat" default:"simple"`

	// NoColor disables color output
	NoColor bool `env:"LOG_NO_COLOR" default:"false"`

	// logs messages from reconciliation events like Reconciliation started etc...
	LogInternalEvents bool `env:"LOG_INTERNAL_EVENTS" default:"false"`
}
