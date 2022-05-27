// logging module with sentry capture
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
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableQuote: true,
	})
	cfg, _ := configuration.GetConfig()
	sentryDsn = cfg.SentryDsn
	if sentryDsn == "" {
		sentryControl = false
	}
}

func Fatalf(sendToSentry bool, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	logrus.Fatalln(msg)
	if sentryControl && sendToSentry {
		sentry.CaptureMessage(msg)
	}
}

func Warnf(sendToSentry bool, format string, args ...any) {
	msg := fmt.Errorf(format, args...)
	logrus.Warnln(msg)
	if sentryControl && sendToSentry {
		sentry.CaptureException(msg)
	}
}

func Infof(sendToSentry bool, format string, args ...any) {
	msg := fmt.Errorf(format, args...)
	logrus.Infoln(msg)
	if sentryControl && sendToSentry {
		sentry.CaptureException(msg)
	}
}
