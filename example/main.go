package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	v1 "github.com/panjiang/protoc-gen-oneof-api-adapter/example/api/product/app/v1"
	"google.golang.org/protobuf/proto"
)

type ApiHandler struct {
}

// LoginWithAccount implements v1.ApiHandler.
func (a *ApiHandler) LoginWithAccount(context.Context, *v1.LoginWithAccountRequest) (*v1.LoginWithAccountResponse, error) {
	panic("unimplemented")
}

// LoginWithOpenID implements v1.ApiHandler.
func (a *ApiHandler) LoginWithOpenID(context.Context, *v1.LoginWithOpenIDRequest) (*v1.LoginWithOpenIDResponse, error) {
	panic("unimplemented")
}

type GinHandler struct {
	apiAdapter v1.ApiAdapter
}

func (h *GinHandler) parseRequest(c *gin.Context) (*v1.Request, error) {
	if c.ContentType() != binding.MIMEPROTOBUF {
		return nil, fmt.Errorf("unsupported ContentType:%s", c.ContentType())
	}

	if c.Request.ContentLength == 0 {
		return nil, errors.New("empty http body")
	}

	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	request := new(v1.Request)
	err = proto.Unmarshal(b, request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (h *GinHandler) Handle(c *gin.Context) {
	ctx := c.Request.Context()

	request, err := h.parseRequest(c)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	response, err := h.apiAdapter.Dispatch(ctx, request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.ProtoBuf(200, response)
}

func main() {
	apiHandler := &ApiHandler{}
	apiAdapter := v1.NewApiAdapter(apiHandler)
	ginHandler := GinHandler{apiAdapter: apiAdapter}

	r := gin.New()
	r.POST("/", ginHandler.Handle)
	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
