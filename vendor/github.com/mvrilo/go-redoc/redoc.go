package redoc

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	_ "embed"
)

// ErrSpecNotFound error for when spec file not found
var ErrSpecNotFound = errors.New("spec not found")

// Redoc configuration
type Redoc struct {
	DocsPath    string
	SpecPath    string
	SpecFile    string
	Title       string
	Description string
}

// HTML represents the redoc index.html page
//go:embed assets/index.html
var HTML string

// JavaScript represents the redoc standalone javascript
//go:embed assets/redoc.standalone.js
var JavaScript string

// Body returns the final html with the js in the body
func (r Redoc) Body() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	tpl, err := template.New("redoc").Parse(HTML)
	if err != nil {
		return nil, err
	}

	if err = tpl.Execute(buf, map[string]string{
		"body":        JavaScript,
		"title":       r.Title,
		"url":         r.SpecPath,
		"description": r.Description,
	}); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Handler sets some defaults and returns a HandlerFunc
func (r Redoc) Handler() http.HandlerFunc {
	data, err := r.Body()
	if err != nil {
		panic(err)
	}

	specFile := r.SpecFile
	if specFile == "" {
		panic(ErrSpecNotFound)
	}

	specPath := r.SpecPath
	if specPath == "" {
		specPath = "/openapi.json"
	}

	spec, err := ioutil.ReadFile(specFile)
	if err != nil {
		panic(err)
	}

	docsPath := r.DocsPath
	return func(w http.ResponseWriter, req *http.Request) {
		method := strings.ToLower(req.Method)
		if method != "get" && method != "head" {
			return
		}

		if strings.HasSuffix(req.URL.Path, r.SpecPath) {
			w.WriteHeader(200)
			w.Header().Set("content-type", "application/json")
			w.Write(spec)
			return
		}

		if docsPath == "" || docsPath == req.URL.Path {
			w.WriteHeader(200)
			w.Header().Set("content-type", "text/html")
			w.Write(data)
		}
	}
}
