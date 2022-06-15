// logging module with sentry capture
// TODO: also send logs to logstash
package logger

import (
	"fmt"

	"github.com/mhkarimi1383/goAPIBaseProject/configuration"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

var (
	sentryDsn     string
	sentryControl bool = true
)

func init() {
	cfg, _ := configuration.GetConfig()
	sentryDsn = cfg.SentryDsn
	if sentryDsn == "" {
		sentryControl = false
	}
	if cfg.LogFormat == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableQuote: true,
		})
	}
}

func Fatalf(sendToSentry bool, format string, args ...any) {
	msg := fmt.Errorf(format, args...)
	if sentryControl && sendToSentry {
		sentry.CaptureException(msg)
	}
	logrus.Fatalln(msg)
}

func Warnf(sendToSentry bool, format string, args ...any) {
	msg := fmt.Errorf(format, args...)
	if sentryControl && sendToSentry {
		sentry.CaptureException(msg)
	}
	logrus.Warnln(msg)
}

func Infof(sendToSentry bool, format string, args ...any) {
	msg := fmt.Errorf(format, args...)
	if sentryControl && sendToSentry {
		sentry.CaptureMessage(msg.Error())
	}
	logrus.Infoln(msg)
}

func Debugf(sendToSentry bool, format string, args ...any) {
	msg := fmt.Errorf(format, args...)
	if sentryControl && sendToSentry {
		sentry.CaptureMessage(msg.Error())
	}
	logrus.Debugln(msg)
}
