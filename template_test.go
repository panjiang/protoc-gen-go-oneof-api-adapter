package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdapterTemplate(t *testing.T) {
	tmpl := &AdapterTemplate{
		ApiName:                "Api",
		PackageName:            "userpb",
		UnionRequestType:       "Request",
		UnionResponseType:      "Response",
		UnionRequestBodyField:  "Body",
		UnionResponseBodyField: "Body",
	}

	for _, v := range []string{"Register", "Login"} {
		hf := HandlerFunc{
			Method:             v,
			RequestType:        v + "Request",
			ResponseType:       v + "Response",
			RequestOneofType:   "Request_" + v + "Request",
			RequestOneofField:  v + "Request",
			ResponseOneofType:  "Response_" + v + "Response",
			ResponseOneofField: v + "Response",
		}
		tmpl.HandlerFuncList = append(tmpl.HandlerFuncList, hf)
	}

	text, err := tmpl.execute()
	require.NoError(t, err)
	print(text)
}
