// all types are here
package types

import "time"

type Configuration struct {
	MetricAddress string `env:"METRIC_ADDRESS" env-default:":9090" yaml:"metric_address"`
	APIAddress    string `env:"API_ADDRESS" env-default:":8080" yaml:"api_address"`
	SentryDsn     string `env:"SENTRY_DSN" env-default:"" yaml:"sentry_dsn"`
	LogFormat     string `env:"LOG_FORMAT" env-default:"text" yaml:"log_format"`
}

type HealthzResponse struct {
	Name    string `json:"name"`
	Message string `json:"msg"`
}

type HelloResponse struct {
	Greeting string    `json:"greeting"`
	Time     time.Time `json:"time"`
}

type UntypedMap map[any]any

type Response interface {
	HealthzResponse |
		HelloResponse
}
