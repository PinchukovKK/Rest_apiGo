// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// User defines model for User.
type User struct {
	Email    *string `json:"email,omitempty"`
	Id       *uint   `json:"id,omitempty"`
	Password *string `json:"password,omitempty"`
}

// PatchUsersByIdJSONRequestBody defines body for PatchUsersById for application/json ContentType.
type PatchUsersByIdJSONRequestBody = User

// PostUsersPostJSONRequestBody defines body for PostUsersPost for application/json ContentType.
type PostUsersPostJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// Update a user
	// (PATCH /users/by/{id})
	PatchUsersById(ctx echo.Context, id int) error
	// Delete a user
	// (DELETE /users/delete/{id})
	DeleteUsersDeleteId(ctx echo.Context, id int) error
	// Create a new user
	// (POST /users/post)
	PostUsersPost(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PatchUsersById converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUsersById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUsersById(ctx, id)
	return err
}

// DeleteUsersDeleteId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsersDeleteId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUsersDeleteId(ctx, id)
	return err
}

// PostUsersPost converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsersPost(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsersPost(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.PATCH(baseURL+"/users/by/:id", wrapper.PatchUsersById)
	router.DELETE(baseURL+"/users/delete/:id", wrapper.DeleteUsersDeleteId)
	router.POST(baseURL+"/users/post", wrapper.PostUsersPost)

}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []User

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchUsersByIdRequestObject struct {
	Id   int `json:"id"`
	Body *PatchUsersByIdJSONRequestBody
}

type PatchUsersByIdResponseObject interface {
	VisitPatchUsersByIdResponse(w http.ResponseWriter) error
}

type PatchUsersById200JSONResponse User

func (response PatchUsersById200JSONResponse) VisitPatchUsersByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteUsersDeleteIdRequestObject struct {
	Id int `json:"id"`
}

type DeleteUsersDeleteIdResponseObject interface {
	VisitDeleteUsersDeleteIdResponse(w http.ResponseWriter) error
}

type DeleteUsersDeleteId204Response struct {
}

func (response DeleteUsersDeleteId204Response) VisitDeleteUsersDeleteIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type PostUsersPostRequestObject struct {
	Body *PostUsersPostJSONRequestBody
}

type PostUsersPostResponseObject interface {
	VisitPostUsersPostResponse(w http.ResponseWriter) error
}

type PostUsersPost201JSONResponse User

func (response PostUsersPost201JSONResponse) VisitPostUsersPostResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Update a user
	// (PATCH /users/by/{id})
	PatchUsersById(ctx context.Context, request PatchUsersByIdRequestObject) (PatchUsersByIdResponseObject, error)
	// Delete a user
	// (DELETE /users/delete/{id})
	DeleteUsersDeleteId(ctx context.Context, request DeleteUsersDeleteIdRequestObject) (DeleteUsersDeleteIdResponseObject, error)
	// Create a new user
	// (POST /users/post)
	PostUsersPost(ctx context.Context, request PostUsersPostRequestObject) (PostUsersPostResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchUsersById operation middleware
func (sh *strictHandler) PatchUsersById(ctx echo.Context, id int) error {
	var request PatchUsersByIdRequestObject

	request.Id = id

	var body PatchUsersByIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchUsersById(ctx.Request().Context(), request.(PatchUsersByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchUsersById")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchUsersByIdResponseObject); ok {
		return validResponse.VisitPatchUsersByIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteUsersDeleteId operation middleware
func (sh *strictHandler) DeleteUsersDeleteId(ctx echo.Context, id int) error {
	var request DeleteUsersDeleteIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUsersDeleteId(ctx.Request().Context(), request.(DeleteUsersDeleteIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUsersDeleteId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUsersDeleteIdResponseObject); ok {
		return validResponse.VisitDeleteUsersDeleteIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUsersPost operation middleware
func (sh *strictHandler) PostUsersPost(ctx echo.Context) error {
	var request PostUsersPostRequestObject

	var body PostUsersPostJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsersPost(ctx.Request().Context(), request.(PostUsersPostRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsersPost")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUsersPostResponseObject); ok {
		return validResponse.VisitPostUsersPostResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
