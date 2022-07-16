package main

import (
	"bytes"
	"text/template"
)

var confTemplate = `
{{ if .HasComment }}{{ .Comment }}{{ end -}}
type I{{.Name}} interface {
	{{- range .Methods }}
	{{ if .HasComment }}{{ .Comment }}{{ end -}}
	Get{{ .Name }}() {{ .RetType }}
	{{- end }}
}
`

type method struct {
	Name       string
	RetType    string
	Comment    string
	HasComment bool
}

type configWrapper struct {
	Name       string
	Comment    string
	HasComment bool
	Methods    []*method
}

func (e *configWrapper) execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("configs").Parse(confTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}
	return buf.String()
}
