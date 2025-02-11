
var (
	Err{{ $.ApiName }}UnknownRequest = errors.New("unknown request")
)

type {{ $.HandlerInterfaceName }} interface {
{{- range .HandlerFuncList }}
	{{ .Method }}(context.Context, *{{ $.WithPackageName .RequestType }}) (*{{ $.WithPackageName .ResponseType }}, error)
{{- end }}
}

type {{ $.AdapterInterfaceName }} interface {
	Dispatch(context.Context, *{{ $.WithPackageName $.UnionRequestType }}) (*{{ $.WithPackageName $.UnionResponseType }}, error)
}

func New{{ $.AdapterInterfaceName }}(handler {{ $.HandlerInterfaceName }}) {{ $.AdapterInterfaceName }} {
	return &{{ $.AdapterImplementationName }}{
		handler: handler,
	}
}

type {{ $.AdapterImplementationName }} struct {
	handler {{ $.HandlerInterfaceName }}
}

func (a *{{ $.AdapterImplementationName }}) Dispatch(ctx context.Context, request *{{ $.WithPackageName $.UnionRequestType }}) (*{{ $.WithPackageName $.UnionResponseType }}, error) {
    response := new({{ $.WithPackageName $.UnionResponseType }})
	switch req := request.{{ $.UnionRequestBodyField }}.(type) {
    {{- range .HandlerFuncList }}
    case *{{ $.WithPackageName .RequestOneofType}}:
		resp, err := a.handler.{{.Method}}(ctx, req.{{ .RequestOneofField }})
		if err != nil {
			return nil, err
		}
		response.{{ $.UnionResponseBodyField }} = &{{ $.WithPackageName .ResponseOneofType }}{
			{{ .ResponseOneofField }}: resp,
		}
    {{- end }}
    default:
		return nil, Err{{ $.ApiName }}UnknownRequest
    }
	return response, nil
}