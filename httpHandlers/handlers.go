// all http handlers are here
package httpHandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mhkarimi1383/goAPIBaseProject/httpServer"
	"github.com/mhkarimi1383/goAPIBaseProject/logger"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	middlewarestd "github.com/slok/go-http-metrics/middleware/std"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "I AM Healthy :) "+parameters["name"])
	if err != nil {
		log.Println("err")
	}
}

func healthzHandler() http.Handler {
	return http.HandlerFunc(healthz)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprint(w, "404; sorry the page that you want is not exist")
	if err != nil {
		log.Println("err")
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
	router.Handle("/healthz", httpServer.WithLogging(healthzHandler()))
	router.NotFoundHandler = httpServer.WithLogging(notFoundHandler())
	mrouter := middlewarestd.Handler("", mdlw, router)
	go func() {
		logger.Fatalf(true, "error in metric http server: %v", http.ListenAndServe(":9090", promhttp.Handler()))
	}()
	logger.Fatalf(true, "error in main http server: %v", http.ListenAndServe(":8080", mrouter))
}
