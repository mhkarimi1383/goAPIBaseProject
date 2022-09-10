// Package logger logging module with sentry capture on top of logrus
// TODO: also send logs to logstash
package logger

import (
	"fmt"
	"github.com/mhkarimi1383/goAPIBaseProject/configuration"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

var (
	// global variable to hold the sentry dsn
	sentryDsn string
	// global variable to hold whether to send the logs to sentry or not
	sentryControl = true
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
			ForceQuote: true,
		})
	}
}

// Fatalf log a message with Fatal level
func Fatalf(sendToSentry bool, format string, args ...any) {
	format = escaper(format)
	msg := fmt.Errorf(format, args...)
	if sentryControl && sendToSentry {
		sentry.CaptureException(msg)
	}
	logrus.Fatalln(msg)
}

// Warnf log a message with Warning level
func Warnf(sendToSentry bool, format string, args ...any) {
	format = escaper(format)
	msg := fmt.Errorf(format, args...)
	if sentryControl && sendToSentry {
		sentry.CaptureException(msg)
	}
	logrus.Warnln(msg)
}

// Infof log a message with Info level
func Infof(sendToSentry bool, format string, args ...any) {
	format = escaper(format)
	msg := fmt.Errorf(format, args...)
	if sentryControl && sendToSentry {
		sentry.CaptureMessage(msg.Error())
	}
	logrus.Infoln(msg)
}

// Debugf log a message with Debug level
func Debugf(sendToSentry bool, format string, args ...any) {
	format = escaper(format)
	msg := fmt.Errorf(format, args...)
	if sentryControl && sendToSentry {
		sentry.CaptureMessage(msg.Error())
	}
	logrus.Debugln(msg)
}
