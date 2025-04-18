// Package handler provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CreateUserRequest defines model for CreateUserRequest.
type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

// CreateVideoRequest defines model for CreateVideoRequest.
type CreateVideoRequest struct {
	Description *string `json:"description,omitempty"`
	OwnerId     string  `json:"owner_id"`
	Thumbnail   *string `json:"thumbnail,omitempty"`
	Title       string  `json:"title"`
	Url         string  `json:"url"`
}

// RankingsResponse defines model for RankingsResponse.
type RankingsResponse struct {
	Rankings *[]VideoRank `json:"rankings,omitempty"`
}

// UpdateUserRequest defines model for UpdateUserRequest.
type UpdateUserRequest struct {
	Email    openapi_types.Email `json:"email"`
	Username string              `json:"username"`
}

// UpdateVideoReactionsRequest defines model for UpdateVideoReactionsRequest.
type UpdateVideoReactionsRequest struct {
	Comments *int64 `json:"comments,omitempty"`
	Likes    *int64 `json:"likes,omitempty"`
	Shares   *int64 `json:"shares,omitempty"`
	Views    *int64 `json:"views,omitempty"`

	// WatchTime Time in seconds
	WatchTime *int64 `json:"watch_time,omitempty"`
}

// UpdateVideoRequest defines model for UpdateVideoRequest.
type UpdateVideoRequest struct {
	Description *string `json:"description,omitempty"`
	Thumbnail   *string `json:"thumbnail,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// UserRankingResponse defines model for UserRankingResponse.
type UserRankingResponse struct {
	Points *float32 `json:"points,omitempty"`
	Rank   *int     `json:"rank,omitempty"`
	UserId *string  `json:"user_id,omitempty"`
}

// UserResponse defines model for UserResponse.
type UserResponse struct {
	Email    *string `json:"email,omitempty"`
	Id       *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
}

// VideoRank defines model for VideoRank.
type VideoRank struct {
	Score   *float32 `json:"score,omitempty"`
	VideoId *string  `json:"video_id,omitempty"`
}

// VideoResponse defines model for VideoResponse.
type VideoResponse struct {
	Comments    *int64     `json:"comments,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	Description *string    `json:"description,omitempty"`
	Id          *string    `json:"id,omitempty"`
	Likes       *int64     `json:"likes,omitempty"`
	OwnerId     *string    `json:"owner_id,omitempty"`
	Shares      *int64     `json:"shares,omitempty"`
	Thumbnail   *string    `json:"thumbnail,omitempty"`
	Title       *string    `json:"title,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Url         *string    `json:"url,omitempty"`
	Views       *int64     `json:"views,omitempty"`
	WatchTime   *int64     `json:"watch_time,omitempty"`
}

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = CreateUserRequest

// PutUsersIdJSONRequestBody defines body for PutUsersId for application/json ContentType.
type PutUsersIdJSONRequestBody = UpdateUserRequest

// PostVideosJSONRequestBody defines body for PostVideos for application/json ContentType.
type PostVideosJSONRequestBody = CreateVideoRequest

// PutVideosIdJSONRequestBody defines body for PutVideosId for application/json ContentType.
type PutVideosIdJSONRequestBody = UpdateVideoRequest

// PatchVideosIdReactionsJSONRequestBody defines body for PatchVideosIdReactions for application/json ContentType.
type PatchVideosIdReactionsJSONRequestBody = UpdateVideoReactionsRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get real-time video rankings
	// (GET /rankings/{top})
	GetRankingsTop(w http.ResponseWriter, r *http.Request, top int64)
	// Get real-time rankings for a specific user
	// (GET /rankings/{top}/{userID})
	GetRankingsTopUserID(w http.ResponseWriter, r *http.Request, top int64, userID string)
	// Create a new user
	// (POST /users)
	PostUsers(w http.ResponseWriter, r *http.Request)
	// Delete user by ID
	// (DELETE /users/{id})
	DeleteUsersId(w http.ResponseWriter, r *http.Request, id string)
	// Get user by ID
	// (GET /users/{id})
	GetUsersId(w http.ResponseWriter, r *http.Request, id string)
	// Update user by ID
	// (PUT /users/{id})
	PutUsersId(w http.ResponseWriter, r *http.Request, id string)
	// Create a video
	// (POST /videos)
	PostVideos(w http.ResponseWriter, r *http.Request)
	// Delete a video by ID
	// (DELETE /videos/{id})
	DeleteVideosId(w http.ResponseWriter, r *http.Request, id string)
	// Get a video by ID
	// (GET /videos/{id})
	GetVideosId(w http.ResponseWriter, r *http.Request, id string)
	// Update a video by ID
	// (PUT /videos/{id})
	PutVideosId(w http.ResponseWriter, r *http.Request, id string)
	// Update video reactions (like, comment, share, view)
	// (PATCH /videos/{id}/reactions)
	PatchVideosIdReactions(w http.ResponseWriter, r *http.Request, id string)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

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

// Delete user by ID
// (DELETE /users/{id})
func (_ Unimplemented) DeleteUsersId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get user by ID
// (GET /users/{id})
func (_ Unimplemented) GetUsersId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update user by ID
// (PUT /users/{id})
func (_ Unimplemented) PutUsersId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a video
// (POST /videos)
func (_ Unimplemented) PostVideos(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete a video by ID
// (DELETE /videos/{id})
func (_ Unimplemented) DeleteVideosId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a video by ID
// (GET /videos/{id})
func (_ Unimplemented) GetVideosId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a video by ID
// (PUT /videos/{id})
func (_ Unimplemented) PutVideosId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update video reactions (like, comment, share, view)
// (PATCH /videos/{id}/reactions)
func (_ Unimplemented) PatchVideosIdReactions(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

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

// DeleteUsersId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUsersId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUsersId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUsersId operation middleware
func (siw *ServerInterfaceWrapper) GetUsersId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutUsersId operation middleware
func (siw *ServerInterfaceWrapper) PutUsersId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUsersId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostVideos operation middleware
func (siw *ServerInterfaceWrapper) PostVideos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostVideos(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteVideosId operation middleware
func (siw *ServerInterfaceWrapper) DeleteVideosId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteVideosId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetVideosId operation middleware
func (siw *ServerInterfaceWrapper) GetVideosId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetVideosId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutVideosId operation middleware
func (siw *ServerInterfaceWrapper) PutVideosId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutVideosId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PatchVideosIdReactions operation middleware
func (siw *ServerInterfaceWrapper) PatchVideosIdReactions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchVideosIdReactions(w, r, id)
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
		r.Get(options.BaseURL+"/rankings/{top}", wrapper.GetRankingsTop)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/rankings/{top}/{userID}", wrapper.GetRankingsTopUserID)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users", wrapper.PostUsers)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/users/{id}", wrapper.DeleteUsersId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users/{id}", wrapper.GetUsersId)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/users/{id}", wrapper.PutUsersId)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/videos", wrapper.PostVideos)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/videos/{id}", wrapper.DeleteVideosId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/videos/{id}", wrapper.GetVideosId)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/videos/{id}", wrapper.PutVideosId)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/videos/{id}/reactions", wrapper.PatchVideosIdReactions)
	})

	return r
}
