// all types are here
package types

import "time"

// Configuration is used to store the configuration by cleanenv package
type Configuration struct {
	MetricAddress          string `env:"METRIC_ADDRESS" env-default:":9090" yaml:"metric_address"`
	APIAddress             string `env:"API_ADDRESS" env-default:":8080" yaml:"api_address"`
	SentryDsn              string `env:"SENTRY_DSN" env-default:"" yaml:"sentry_dsn"`
	LogFormat              string `env:"LOG_FORMAT" env-default:"text" yaml:"log_format"`
	ApplicationTitle       string `env:"APPLICATION_TITLE" env-default:"goAPIBaseProject" yaml:"application_title"`
	ApplicationDescription string `env:"APPLICATION_DESCRIPTION" env-default:"goAPIBaseProject" yaml:"application_description"`
}

type ApplicationInformation struct {
	Title       string
	Description string
}

type HealthzResponse struct {
	Name    string `json:"name"`
	Message string `json:"msg"`
}

type HelloResponse struct {
	Greeting string    `json:"greeting"`
	Time     time.Time `json:"time"`
}

// a type that used to create untyped map (for json)
type UntypedMap map[any]any

// any acceptable response will be here using generic it will accept one of the given types
type Response interface {
	HealthzResponse |
		HelloResponse
}
