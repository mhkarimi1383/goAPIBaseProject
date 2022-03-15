# Go Simple API Project

This project made to combine some pretty beautiful thing together

## This project is using
* `net/http` as http server
* [gorilla/mux](https://github.com/gorilla/mux) as router
* Sentry as APM
* [sirupsen/logrus](https://github.com/sirupsen/logrus) as base of logging
* [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) for managing configurations
* [slok/go-http-metrics](https://github.com/slok/go-http-metrics) as middleware for exporting prometheus metrics

## Project files structure
* `logger`: package for logging on top of [sirupsen/logrus](https://github.com/sirupsen/logrus) with option to sent logs to Sentry
* `httpHandlers`: all of the http handlers shoud be here
* `httpServer`: http server components are here ready to use
* `structures`: every structure that we want every where should be here
* `configuration`: on top of [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) this will manage our configurations with a simple function that returns an variable with Configuration struct
* `air`: executable version of [cosmtrek/air](https://github.com/cosmtrek/air) is here to help you with live reloading you code (configuration for that is present here)
