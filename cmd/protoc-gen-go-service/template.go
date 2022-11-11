package main

import (
	"bytes"
	"strings"
	"text/template"
)

var serviceTemplate = `
{{$svrType := .ServiceType}}

type {{.ServiceType}}Service struct{
	pb.Unimplemented{{.ServiceType}}Server
}
	
func New() pb.{{.ServiceType}}HTTPServer {
	return &{{.ServiceType}}Service{}
}

{{range .MethodSets}}
func (c *{{$svrType}}Service) {{.Name}}(ctx context.Context, param *pb.{{.Request}}) (*pb.{{.Reply}}, error) {
	return &pb.{{.Reply}}{}, nil
}
{{end}}
`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
}

type methodDesc struct {
	// method
	Name         string
	OriginalName string // The parsed original name
	Num          int
	Request      string
	Reply        string
	// http_rule
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("service").Parse(strings.TrimSpace(serviceTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
