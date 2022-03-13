// all structures are here
package structures

type Configuration struct {
	MetricAddress string `env:"METRIC_ADDRESS" env-default:":9090" yaml:"metric_address"`
	APIAddress    string `env:"API_ADDRESS" env-default:":8080" yaml:"api_address"`
	SentryDsn     string `env:"SENTRY_DSN" env-default:"" yaml:"sentry_dsn"`
}
