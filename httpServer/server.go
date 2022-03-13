// our custom http server goes here
package httpServer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

// simply for holding some data
type ResponseData struct {
	status int
	size   int
}

// custom response writer with compatibility for original one
type LoggingResponseWriter struct {
	http.ResponseWriter
	ResponseData *ResponseData
}

func (r *LoggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b) // write response using original http.ResponseWriter
	r.ResponseData.size += size            // capture size
	return size, err
}

func (r *LoggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode) // write status code using original http.ResponseWriter
	r.ResponseData.status = statusCode       // capture status code
}

// return http handler with logging capabilities
func WithLogging(h http.Handler) http.Handler {
	loggingFn := func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()

		responseData := &ResponseData{
			status: 0,
			size:   0,
		}
		lrw := LoggingResponseWriter{
			ResponseWriter: rw, // compose original http.ResponseWriter
			ResponseData:   responseData,
		}
		h.ServeHTTP(&lrw, req) // inject our implementation of http.ResponseWriter
		duration := time.Since(start)
		if responseData.status == http.StatusOK {
			logrus.WithFields(logrus.Fields{
				"uri":            req.RequestURI,
				"method":         req.Method,
				"status":         responseData.status,
				"duration":       duration,
				"size":           responseData.size,
				"remote address": req.RemoteAddr,
			}).Info("request successfully done")
		} else {
			logrus.WithFields(logrus.Fields{
				"uri":            req.RequestURI,
				"method":         req.Method,
				"status":         responseData.status,
				"duration":       duration,
				"size":           responseData.size,
				"remote address": req.RemoteAddr,
			}).Warn("request done with problem")
			sentry.CaptureException(fmt.Errorf("request from alertmanager done with problem, uri: %v, method: %v, status: %v, duration: %v, size: %v, remote address: %v",
				req.RequestURI,
				req.Method,
				responseData.status,
				duration,
				responseData.size,
				req.RemoteAddr))
		}
	}
	return http.HandlerFunc(loggingFn)
}
