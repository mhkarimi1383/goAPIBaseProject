package httpHandlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/mhkarimi1383/goAPIBaseProject/logger"
	"github.com/mvrilo/go-redoc"
)

const (
	// rapidocTemplate is a template for rapidoc HTML page
	// TODO: Move this to a template file
	rapidocTemplate = `<!doctype html>
<html>
<head>
	<title>{{.Title}}</title>
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
</html>`
)

var (
	// configuration of redoc
	doc *redoc.Redoc
	// global template
	tmpl *template.Template
)

func init() {
	doc = &redoc.Redoc{
		Title:       information.Title,
		Description: information.Description,
		SpecFile:    "./openapi.json",
		SpecPath:    "/redoc/openapi.json",
	}
	tmpl, err := template.New("rapidoc").Parse(rapidocTemplate)
	if err != nil {
		logger.Fatalf(true, "error in parsing template: %v", err)
	}
	tmpl, err = template.ParseFiles("openapi.json.tpl")
	if err != nil {
		logger.Fatalf(true, "error in parsing template: %v", err)
		return
	}
	f, err := os.Create("openapi.json")
	if err != nil {
		logger.Fatalf(true, "error in creating file: %v", err)
		return
	}
	err = tmpl.Execute(f, information)
	if err != nil {
		logger.Fatalf(true, "error in executing template: %v", err)
		return
	}
	tmpl, _ = template.New("rapidoc").Parse(rapidocTemplate)
}

// handler for rapidoc
func rapiDoc() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.WriteHeader(http.StatusOK)
		err := tmpl.Execute(rw, information)
		if err != nil {
			logger.Fatalf(true, "error in executing template: %v", err)
		}
		return
	})
}
