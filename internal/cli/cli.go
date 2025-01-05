package cli

import (
	"github.com/rs/zerolog"
)

type LogFormat string

const (
	// JSONFormat prints messages as single json record
	JSONFormat LogFormat = "json"
	// SimpleFormat prints messages in human readable way
	SimpleFormat LogFormat = "simple"
	// NoFormat returned in situation when format is not recognised
	NoFormat LogFormat = "noformat"
)

// CLI https://github.com/alecthomas/kong
type CLI struct {
	// Namespaces to be isolated for the operator separated by comma.
	AllowedNamespaces []string `optional:"" env:"ALLOWED_NAMESPACES"`

	Manager struct {
		// Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.
		EnableLeaderElection bool `optional:"" env:"ENABLE_LEADER_ELECTION" default:"false"`

		// The address the probe endpoint binds to.
		HealthProbeBindAddress string `optional:"" env:"HEALTH_PROBE_BIND_ADDRESS" default:":8081" validate:"hostname_port"`
	} `embed:""`

	Metrics struct {
		// If set, HTTP/2 will be enabled for the metrics and webhook servers
		EnableHttp2 bool `optional:"" env:"ENABLE_HTTP2" default:"false"`

		// If set, the metrics endpoint is served securely via HTTPS. Use --metrics-secure=false to use HTTP instead.
		SecureMetrics bool `optional:"" env:"SECURE_METRICS" default:"true"`

		// The address the metrics endpoint binds to.Use :8443 for HTTPS or :8080 for HTTP, or leave as 0 to disable the metrics service.
		MetricsBindAddress string `optional:"" env:"METRICS_BIND_ADDRESS" default:"0"`
	} `embed:""`

	Vault struct {
		// Vault URL address
		Url string `env:"VAULT_ADDR" validate:"required,http_url"`

		// Vault token
		Token string `env:"VAULT_TOKEN" validate:"required"`
	} `embed:""`

	Log Log `embed:""`
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
