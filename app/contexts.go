// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: Application Contexts
//
// Command:
// $ goagen
// --design=github.com/Microkubes/authorization-server/design
// --out=$(GOPATH)/src/github.com/Microkubes/authorization-server
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// ConfirmAuthorizationAuthUIContext provides the authUI confirmAuthorization action context.
type ConfirmAuthorizationAuthUIContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Confirmed *bool
}

// NewConfirmAuthorizationAuthUIContext parses the incoming request URL and body, performs validations and creates the
// context used by the authUI controller confirmAuthorization action.
func NewConfirmAuthorizationAuthUIContext(ctx context.Context, r *http.Request, service *goa.Service) (*ConfirmAuthorizationAuthUIContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ConfirmAuthorizationAuthUIContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramConfirmed := req.Params["confirmed"]
	if len(paramConfirmed) > 0 {
		rawConfirmed := paramConfirmed[0]
		if confirmed, err2 := strconv.ParseBool(rawConfirmed); err2 == nil {
			tmp1 := &confirmed
			rctx.Confirmed = tmp1
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("confirmed", rawConfirmed, "boolean"))
		}
	}
	return &rctx, err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ConfirmAuthorizationAuthUIContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ConfirmAuthorizationAuthUIContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// PromptAuthorizationAuthUIContext provides the authUI promptAuthorization action context.
type PromptAuthorizationAuthUIContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewPromptAuthorizationAuthUIContext parses the incoming request URL and body, performs validations and creates the
// context used by the authUI controller promptAuthorization action.
func NewPromptAuthorizationAuthUIContext(ctx context.Context, r *http.Request, service *goa.Service) (*PromptAuthorizationAuthUIContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := PromptAuthorizationAuthUIContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *PromptAuthorizationAuthUIContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *PromptAuthorizationAuthUIContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ShowLoginLoginContext provides the login showLogin action context.
type ShowLoginLoginContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowLoginLoginContext parses the incoming request URL and body, performs validations and creates the
// context used by the login controller showLogin action.
func NewShowLoginLoginContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowLoginLoginContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowLoginLoginContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowLoginLoginContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ShowLoginLoginContext) Unauthorized(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowLoginLoginContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// AuthorizeOauth2ProviderContext provides the oauth2_provider authorize action context.
type AuthorizeOauth2ProviderContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ClientID     string
	RedirectURI  *string
	ResponseType string
	Scope        *string
	State        *string
}

// NewAuthorizeOauth2ProviderContext parses the incoming request URL and body, performs validations and creates the
// context used by the oauth2_provider controller authorize action.
func NewAuthorizeOauth2ProviderContext(ctx context.Context, r *http.Request, service *goa.Service) (*AuthorizeOauth2ProviderContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AuthorizeOauth2ProviderContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramClientID := req.Params["client_id"]
	if len(paramClientID) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("client_id"))
	} else {
		rawClientID := paramClientID[0]
		rctx.ClientID = rawClientID
	}
	paramRedirectURI := req.Params["redirect_uri"]
	if len(paramRedirectURI) > 0 {
		rawRedirectURI := paramRedirectURI[0]
		rctx.RedirectURI = &rawRedirectURI
	}
	paramResponseType := req.Params["response_type"]
	if len(paramResponseType) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("response_type"))
	} else {
		rawResponseType := paramResponseType[0]
		rctx.ResponseType = rawResponseType
		if !(rctx.ResponseType == "code") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response_type`, rctx.ResponseType, []interface{}{"code"}))
		}
	}
	paramScope := req.Params["scope"]
	if len(paramScope) > 0 {
		rawScope := paramScope[0]
		rctx.Scope = &rawScope
	}
	paramState := req.Params["state"]
	if len(paramState) > 0 {
		rawState := paramState[0]
		rctx.State = &rawState
	}
	return &rctx, err
}

// Found sends a HTTP response with status code 302.
func (ctx *AuthorizeOauth2ProviderContext) Found() error {
	ctx.ResponseData.WriteHeader(302)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *AuthorizeOauth2ProviderContext) BadRequest(r *OAuth2ErrorMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.oauth2.error+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// GetTokenOauth2ProviderContext provides the oauth2_provider get_token action context.
type GetTokenOauth2ProviderContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *TokenPayload
}

// NewGetTokenOauth2ProviderContext parses the incoming request URL and body, performs validations and creates the
// context used by the oauth2_provider controller get_token action.
func NewGetTokenOauth2ProviderContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetTokenOauth2ProviderContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetTokenOauth2ProviderContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetTokenOauth2ProviderContext) OK(r *TokenMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.oauth2.token+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *GetTokenOauth2ProviderContext) BadRequest(r *OAuth2ErrorMedia) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.oauth2.error+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}
