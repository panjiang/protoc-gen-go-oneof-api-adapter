package main

import (
	"bytes"
	_ "embed"
	"text/template"
	"unicode"
)

//go:embed template.go.tpl
var tpl string

type AdapterTemplate struct {
	ApiName                string // Api
	PackageName            string // userpb
	UnionRequestType       string // Request
	UnionResponseType      string // Response
	UnionRequestBodyField  string // Body
	UnionResponseBodyField string // Body
	HandlerFuncList        []HandlerFunc
}

func (a *AdapterTemplate) AdapterInterfaceName() string {
	return toUpperFirst(a.ApiName) + "Adapter"
}

func (a *AdapterTemplate) AdapterImplementationName() string {
	return toLowerFirst(a.ApiName) + "Adapter"
}

func (a *AdapterTemplate) HandlerInterfaceName() string {
	return a.ApiName + "Handler"
}

func (a *AdapterTemplate) WithPackageName(name string) string {
	if a.PackageName == "" {
		return name
	}
	return a.PackageName + "." + name
}

func (a *AdapterTemplate) execute() (string, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(tpl)
	if err != nil {
		return "", err
	}
	if err := tmpl.Execute(buf, a); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type HandlerFunc struct {
	Method             string // LoginWithOpenID
	RequestType        string // LoginWithOpenIDRequest
	ResponseType       string // LoginWithOpenIDResponse
	RequestOneofType   string // Request_LoginWithOpenIDRequest
	ResponseOneofType  string // Response_LoginWithOpenIDResponse
	RequestOneofField  string // LoginWithOpenIDRequest
	ResponseOneofField string // LoginWithOpenIDResponse
}

func toLowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func toUpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}
