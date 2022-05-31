// all http handlers are here
package httpHandlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mhkarimi1383/goAPIBaseProject/configuration"
	"github.com/mhkarimi1383/goAPIBaseProject/httpServer"
	"github.com/mhkarimi1383/goAPIBaseProject/logger"
	"github.com/mhkarimi1383/goAPIBaseProject/types"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	middlewarestd "github.com/slok/go-http-metrics/middleware/std"
)

var (
	apiAddress    string
	metricAddress string
)

func init() {
	cfg, err := configuration.GetConfig()
	if err != nil {
		logger.Fatalf(true, "error in initializing configuration: %v", err)
	}
	apiAddress = cfg.APIAddress
	metricAddress = cfg.MetricAddress
}

func greeting(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") == "application/json" {
		// jsonResponse := types.UntypedMap{
		// 	"greeting": "Hello",
		// 	"time": "I am healthy",
		// }
		jsonResponse := types.HelloResponse{
			Greeting: "Hello",
			Time:     time.Now(),
		}
		err := responseWriter(w, &jsonResponse, http.StatusOK)
		if err != nil {
			logger.Warnf(true, "error while sending response %v", err)
		}
		return
	} else {
		strResponse := fmt.Sprintf("Hello, current time is %v", time.Now().Format(time.RFC3339))
		err := responseWriter(w, &strResponse, http.StatusOK)
		if err != nil {
			logger.Warnf(true, "error while sending response %v", err)
		}
		return
	}
}

func greetingHandler() http.Handler {
	return http.HandlerFunc(greeting)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	name := parameters["name"]
	if r.Header.Get("Content-Type") == "application/json" {
		// jsonResponse := types.UntypedMap{
		// 	"name":    name,
		// 	"msg": "I am healthy",
		// }
		jsonResponse := types.HealthzResponse{
			Name:    name,
			Message: "I am healthy",
		}
		err := responseWriter(w, &jsonResponse, http.StatusOK)
		if err != nil {
			logger.Warnf(true, "error while sending response %v", err)
		}
		return
	} else {
		strResponse := fmt.Sprintf("Hello %s", name)
		err := responseWriter(w, &strResponse, http.StatusOK)
		if err != nil {
			logger.Warnf(true, "error while sending response %v", err)
		}
		return
	}
}

func healthzHandler() http.Handler {
	return http.HandlerFunc(healthz)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprint(w, "404; sorry the page that you want is not exist")
	if err != nil {
		logger.Warnf(true, "error while sending response %v", err)
	}
}

func notFoundHandler() http.Handler {
	return http.HandlerFunc(notFound)
}

func RunServer() {
	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.PathPrefix("/redoc").Handler(doc.Handler()) // not logging this section
	router.PathPrefix("/docs").Handler(RapiDoc())      // not logging this section
	router.Handle("/healthz/{name}", httpServer.WithLogging(healthzHandler()))
	router.Handle("/greeting", httpServer.WithLogging(greetingHandler()))
	router.NotFoundHandler = httpServer.WithLogging(notFoundHandler())
	mrouter := middlewarestd.Handler("", mdlw, router)
	go func() {
		logger.Infof(false, "starting metric server on %v", metricAddress)
		logger.Fatalf(true, "error in metric http server: %v", http.ListenAndServe(metricAddress, promhttp.Handler()))
	}()
	handler := cors.Default().Handler(mrouter)
	logger.Infof(false, "starting main server on %v", apiAddress)
	logger.Fatalf(true, "error in main http server: %v", http.ListenAndServe(apiAddress, handler))
}
