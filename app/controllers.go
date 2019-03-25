// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Simple-SNS": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/eniehack/simple-sns-go/design
// --out=D:\project\simple-sns-go
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AuthorizationController is the controller interface for the Authorization actions.
type AuthorizationController interface {
	goa.Muxer
	Register(*RegisterAuthorizationContext) error
}

// MountAuthorizationController "mounts" a Authorization resource controller on the given service.
func MountAuthorizationController(service *goa.Service, ctrl AuthorizationController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRegisterAuthorizationContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RegisterAuthorizationPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Register(rctx)
	}
	h = handleSecurity("JWT", h, "api:access")
	service.Mux.Handle("POST", "/api/v1/auth/new", ctrl.MuxHandler("register", h, unmarshalRegisterAuthorizationPayload))
	service.LogInfo("mount", "ctrl", "Authorization", "action", "Register", "route", "POST /api/v1/auth/new", "security", "JWT")
}

// unmarshalRegisterAuthorizationPayload unmarshals the request body into the context request data Payload field.
func unmarshalRegisterAuthorizationPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &registerAuthorizationPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// PostController is the controller interface for the Post actions.
type PostController interface {
	goa.Muxer
	Create(*CreatePostContext) error
	Delete(*DeletePostContext) error
	Reference(*ReferencePostContext) error
}

// MountPostController "mounts" a Post resource controller on the given service.
func MountPostController(service *goa.Service, ctrl PostController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreatePostContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreatePostPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("JWT", h, "api:access")
	service.Mux.Handle("POST", "/api/v1/posts/new", ctrl.MuxHandler("create", h, unmarshalCreatePostPayload))
	service.LogInfo("mount", "ctrl", "Post", "action", "Create", "route", "POST /api/v1/posts/new", "security", "JWT")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeletePostContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("JWT", h, "api:access")
	service.Mux.Handle("DELETE", "/api/v1/posts/:post_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Post", "action", "Delete", "route", "DELETE /api/v1/posts/:post_id", "security", "JWT")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewReferencePostContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Reference(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/posts/:post_id", ctrl.MuxHandler("reference", h, nil))
	service.LogInfo("mount", "ctrl", "Post", "action", "Reference", "route", "GET /api/v1/posts/:post_id")
}

// unmarshalCreatePostPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreatePostPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createPostPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}