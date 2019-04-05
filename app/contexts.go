// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "area": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/hryktrd/goaTest/design
// --out=$(GOPATH)\src\github.com\hryktrd\goaTest
// --regen=true
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// ListPointContext provides the point list action context.
type ListPointContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListPointContext parses the incoming request URL and body, performs validations and creates the
// context used by the point controller list action.
func NewListPointContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListPointContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListPointContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListPointContext) OK(r PointAreaCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.point.area+json; type=collection")
	}
	if r == nil {
		r = PointAreaCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListPointContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ShowPointContext provides the point show action context.
type ShowPointContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PointID int
}

// NewShowPointContext parses the incoming request URL and body, performs validations and creates the
// context used by the point controller show action.
func NewShowPointContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowPointContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowPointContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPointID := req.Params["pointId"]
	if len(paramPointID) > 0 {
		rawPointID := paramPointID[0]
		if pointID, err2 := strconv.Atoi(rawPointID); err2 == nil {
			rctx.PointID = pointID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("pointId", rawPointID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPointContext) OK(r *PointArea) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.point.area+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPointContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
