// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Simple-SNS": Application Contexts
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
	"unicode/utf8"
)

// RegisterAuthorizationContext provides the Authorization register action context.
type RegisterAuthorizationContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *RegisterAuthorizationPayload
}

// NewRegisterAuthorizationContext parses the incoming request URL and body, performs validations and creates the
// context used by the Authorization controller register action.
func NewRegisterAuthorizationContext(ctx context.Context, r *http.Request, service *goa.Service) (*RegisterAuthorizationContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := RegisterAuthorizationContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// registerAuthorizationPayload is the Authorization register action payload.
type registerAuthorizationPayload struct {
	Password   *string `form:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty" xml:"password,omitempty"`
	ScreenName *string `form:"screen_name,omitempty" json:"screen_name,omitempty" yaml:"screen_name,omitempty" xml:"screen_name,omitempty"`
	Userid     *string `form:"userid,omitempty" json:"userid,omitempty" yaml:"userid,omitempty" xml:"userid,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *registerAuthorizationPayload) Validate() (err error) {
	if payload.Userid == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "userid"))
	}
	if payload.ScreenName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "screen_name"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if payload.ScreenName != nil {
		if utf8.RuneCountInString(*payload.ScreenName) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.screen_name`, *payload.ScreenName, utf8.RuneCountInString(*payload.ScreenName), 20, false))
		}
	}
	if payload.Userid != nil {
		if utf8.RuneCountInString(*payload.Userid) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.userid`, *payload.Userid, utf8.RuneCountInString(*payload.Userid), 15, false))
		}
	}
	return
}

// Publicize creates RegisterAuthorizationPayload from registerAuthorizationPayload
func (payload *registerAuthorizationPayload) Publicize() *RegisterAuthorizationPayload {
	var pub RegisterAuthorizationPayload
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	if payload.ScreenName != nil {
		pub.ScreenName = *payload.ScreenName
	}
	if payload.Userid != nil {
		pub.Userid = *payload.Userid
	}
	return &pub
}

// RegisterAuthorizationPayload is the Authorization register action payload.
type RegisterAuthorizationPayload struct {
	Password   string `form:"password" json:"password" yaml:"password" xml:"password"`
	ScreenName string `form:"screen_name" json:"screen_name" yaml:"screen_name" xml:"screen_name"`
	Userid     string `form:"userid" json:"userid" yaml:"userid" xml:"userid"`
}

// Validate runs the validation rules defined in the design.
func (payload *RegisterAuthorizationPayload) Validate() (err error) {
	if payload.Userid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "userid"))
	}
	if payload.ScreenName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "screen_name"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	if utf8.RuneCountInString(payload.ScreenName) > 20 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.screen_name`, payload.ScreenName, utf8.RuneCountInString(payload.ScreenName), 20, false))
	}
	if utf8.RuneCountInString(payload.Userid) > 15 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.userid`, payload.Userid, utf8.RuneCountInString(payload.Userid), 15, false))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *RegisterAuthorizationContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *RegisterAuthorizationContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RegisterAuthorizationContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// Conflict sends a HTTP response with status code 409.
func (ctx *RegisterAuthorizationContext) Conflict(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 409, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *RegisterAuthorizationContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// CreatePostContext provides the Post create action context.
type CreatePostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreatePostPayload
}

// NewCreatePostContext parses the incoming request URL and body, performs validations and creates the
// context used by the Post controller create action.
func NewCreatePostContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreatePostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreatePostContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createPostPayload is the Post create action payload.
type createPostPayload struct {
	Body *string `form:"body,omitempty" json:"body,omitempty" yaml:"body,omitempty" xml:"body,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createPostPayload) Validate() (err error) {
	if payload.Body == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "body"))
	}
	if payload.Body != nil {
		if utf8.RuneCountInString(*payload.Body) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.body`, *payload.Body, utf8.RuneCountInString(*payload.Body), 500, false))
		}
	}
	return
}

// Publicize creates CreatePostPayload from createPostPayload
func (payload *createPostPayload) Publicize() *CreatePostPayload {
	var pub CreatePostPayload
	if payload.Body != nil {
		pub.Body = *payload.Body
	}
	return &pub
}

// CreatePostPayload is the Post create action payload.
type CreatePostPayload struct {
	Body string `form:"body" json:"body" yaml:"body" xml:"body"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreatePostPayload) Validate() (err error) {
	if payload.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "body"))
	}
	if utf8.RuneCountInString(payload.Body) > 500 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`raw.body`, payload.Body, utf8.RuneCountInString(payload.Body), 500, false))
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *CreatePostContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreatePostContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreatePostContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreatePostContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeletePostContext provides the Post delete action context.
type DeletePostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PostID string
}

// NewDeletePostContext parses the incoming request URL and body, performs validations and creates the
// context used by the Post controller delete action.
func NewDeletePostContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeletePostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeletePostContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPostID := req.Params["post_id"]
	if len(paramPostID) > 0 {
		rawPostID := paramPostID[0]
		rctx.PostID = rawPostID
		if ok := goa.ValidatePattern(`[0-9A-Z]{26}`, rctx.PostID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`post_id`, rctx.PostID, `[0-9A-Z]{26}`))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeletePostContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeletePostContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeletePostContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeletePostContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ReferencePostContext provides the Post reference action context.
type ReferencePostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PostID string
}

// NewReferencePostContext parses the incoming request URL and body, performs validations and creates the
// context used by the Post controller reference action.
func NewReferencePostContext(ctx context.Context, r *http.Request, service *goa.Service) (*ReferencePostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ReferencePostContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPostID := req.Params["post_id"]
	if len(paramPostID) > 0 {
		rawPostID := paramPostID[0]
		rctx.PostID = rawPostID
		if ok := goa.ValidatePattern(`[0-9A-Z]{26}`, rctx.PostID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`post_id`, rctx.PostID, `[0-9A-Z]{26}`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ReferencePostContext) OK(r *Post) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.post+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ReferencePostContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ReferencePostContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ReferencePostContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}