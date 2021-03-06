# Go Simple API Project

[![Go Reference](https://pkg.go.dev/badge/github.com/mhkarimi1383/goAPIBaseProject.svg)](https://pkg.go.dev/github.com/mhkarimi1383/goAPIBaseProject)
[![Go Report Card](https://goreportcard.com/badge/github.com/mhkarimi1383/goAPIBaseProject)](https://goreportcard.com/report/github.com/mhkarimi1383/goAPIBaseProject)

This project made to combine some pretty beautiful thing together

## This project is using

* `net/http` as http server
* [gorilla/mux](https://github.com/gorilla/mux) as router
* Sentry as APM
* [sirupsen/logrus](https://github.com/sirupsen/logrus) as base of logging
* [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) for managing configurations
* [slok/go-http-metrics](https://github.com/slok/go-http-metrics) as middleware for exporting prometheus metrics
* [rs/cors](https://github.com/rs/cors) as middleware for CORS
* [mvrilo/go-redoc](https://github.com/mvrilo/go-redoc) for redoc view
* [rapi-doc/RapiDoc](https://github.com/rapi-doc/RapiDoc) for RapiDoc view instead of swagger
* [go-playground/validator](github.com/go-playground/validator) as base of validation

## Project files structure

* `logger`: package for logging on top of [sirupsen/logrus](https://github.com/sirupsen/logrus) with option to sent logs to Sentry
* `httpHandlers`: all of the http handlers should be here
* `httpServer`: http server components are here ready to use
* `type`: every type that we want every where should be here
* `configuration`: on top of [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) this will manage our configurations with a simple function that returns a variable with Configuration struct
* `air`: executable version of [cosmtrek/air](https://github.com/cosmtrek/air) is here to help you with live reloading you code (configuration for that is present here)
* `validator`: package for validating our things

## TODO

> Moved to [GitHub Project](https://github.com/users/mhkarimi1383/projects/1/)
