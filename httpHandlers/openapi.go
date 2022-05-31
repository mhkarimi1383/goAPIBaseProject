// Package httpHandlers handler for openapi things are here
package httpHandlers

import (
	"net/http"

	"github.com/mvrilo/go-redoc"
)

const (
	rapidocTemplate = `<!doctype html>
<html>
<head>
	<title>Base API Documentation</title>
	<meta charset="utf-8"> <!-- Important: rapi-doc uses utf8 characters -->
	<script type="module" src="https://unpkg.com/rapidoc/dist/rapidoc-min.js"></script>
	<link href='https://fonts.googleapis.com/css?family=JetBrains Mono' rel='stylesheet'>
</head>
<body>
	<rapi-doc
    spec-url = "/redoc/openapi.json"
    show-header = "false"
	show-info = "true"
    mono-font = "JetBrains Mono"
    regular-font = "JetBrains Mono"
	theme = "light"
	render-style="focused"
	allow-try="true" 
	>
	</rapi-doc>
</body>
</html>
`
)

var (
	doc = &redoc.Redoc{
		Title:       "Base API Documentation",
		Description: "Base API Documentation",
		SpecFile:    "./openapi.json",
		SpecPath:    "/redoc/openapi.json",
	}
)

func RapiDoc() http.Handler {
	b := []byte(rapidocTemplate)

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.WriteHeader(http.StatusOK)

		_, _ = rw.Write(b)
		return
	})
}
