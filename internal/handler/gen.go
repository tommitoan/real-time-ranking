// Package handler provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// Defines values for InteractionResponseStatus.
const (
	Failure InteractionResponseStatus = "failure"
	Success InteractionResponseStatus = "success"
)

// CreateUserRequest defines model for CreateUserRequest.
type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

// InteractionResponse defines model for InteractionResponse.
type InteractionResponse struct {
	Message *string                    `json:"message,omitempty"`
	Status  *InteractionResponseStatus `json:"status,omitempty"`
}

// InteractionResponseStatus defines model for InteractionResponse.Status.
type InteractionResponseStatus string

// RankingsResponse defines model for RankingsResponse.
type RankingsResponse struct {
	Rankings *[]VideoRank `json:"rankings,omitempty"`
}

// UserRankingResponse defines model for UserRankingResponse.
type UserRankingResponse struct {
	Points *float32 `json:"points,omitempty"`
	Rank   *int     `json:"rank,omitempty"`
	UserId *string  `json:"user_id,omitempty"`
}

// UserResponse defines model for UserResponse.
type UserResponse struct {
	Id       *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
}

// VideoRank defines model for VideoRank.
type VideoRank struct {
	Score   *float32 `json:"score,omitempty"`
	VideoId *string  `json:"video_id,omitempty"`
}

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = CreateUserRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Handle interactions (e.g., update user ranking or interactions with the video)
	// (POST /interactions)
	PostInteractions(w http.ResponseWriter, r *http.Request)
	// Get real-time video rankings
	// (GET /rankings/{top})
	GetRankingsTop(w http.ResponseWriter, r *http.Request, top int64)
	// Get real-time rankings for a specific user
	// (GET /rankings/{top}/{userID})
	GetRankingsTopUserID(w http.ResponseWriter, r *http.Request, top int64, userID string)
	// Create a new user
	// (POST /users)
	PostUsers(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Handle interactions (e.g., update user ranking or interactions with the video)
// (POST /interactions)
func (_ Unimplemented) PostInteractions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get real-time video rankings
// (GET /rankings/{top})
func (_ Unimplemented) GetRankingsTop(w http.ResponseWriter, r *http.Request, top int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get real-time rankings for a specific user
// (GET /rankings/{top}/{userID})
func (_ Unimplemented) GetRankingsTopUserID(w http.ResponseWriter, r *http.Request, top int64, userID string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new user
// (POST /users)
func (_ Unimplemented) PostUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostInteractions operation middleware
func (siw *ServerInterfaceWrapper) PostInteractions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostInteractions(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetRankingsTop operation middleware
func (siw *ServerInterfaceWrapper) GetRankingsTop(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "top" -------------
	var top int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "top", runtime.ParamLocationPath, chi.URLParam(r, "top"), &top)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "top", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRankingsTop(w, r, top)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetRankingsTopUserID operation middleware
func (siw *ServerInterfaceWrapper) GetRankingsTopUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "top" -------------
	var top int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "top", runtime.ParamLocationPath, chi.URLParam(r, "top"), &top)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "top", Err: err})
		return
	}

	// ------------- Path parameter "userID" -------------
	var userID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "userID", runtime.ParamLocationPath, chi.URLParam(r, "userID"), &userID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRankingsTopUserID(w, r, top, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostUsers operation middleware
func (siw *ServerInterfaceWrapper) PostUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUsers(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/interactions", wrapper.PostInteractions)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/rankings/{top}", wrapper.GetRankingsTop)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/rankings/{top}/{userID}", wrapper.GetRankingsTopUserID)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users", wrapper.PostUsers)
	})

	return r
}
