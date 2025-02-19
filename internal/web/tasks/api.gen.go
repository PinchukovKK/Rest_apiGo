// Package tasks provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	Id     *uint   `json:"id,omitempty"`
	IsDone *bool   `json:"is_done,omitempty"`
	Task   *string `json:"task,omitempty"`
	UserId *uint    `json:"user_id,omitempty"`
}

// GetTasksParams defines parameters for GetTasks.
type GetTasksParams struct {
	UserId uint `form:"user_id" json:"user_id"`
}

// PatchTasksPatchIdJSONRequestBody defines body for PatchTasksPatchId for application/json ContentType.
type PatchTasksPatchIdJSONRequestBody = Task

// PostTasksPostJSONRequestBody defines body for PostTasksPost for application/json ContentType.
type PostTasksPostJSONRequestBody = Task

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx echo.Context, params GetTasksParams) error
	// Update a task
	// (PATCH /tasks/patch/{id})
	PatchTasksPatchId(ctx echo.Context, id int) error
	// Create a new task
	// (POST /tasks/post)
	PostTasksPost(ctx echo.Context) error
	// Delete a task
	// (DELETE /tasks/{id})
	DeleteTasksId(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasks(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTasksParams
	// ------------- Required query parameter "user_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "user_id", ctx.QueryParams(), &params.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasks(ctx, params)
	return err
}

// PatchTasksPatchId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchTasksPatchId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchTasksPatchId(ctx, id)
	return err
}

// PostTasksPost converts echo context to params.
func (w *ServerInterfaceWrapper) PostTasksPost(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostTasksPost(ctx)
	return err
}

// DeleteTasksId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTasksId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteTasksId(ctx, id)
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

	router.GET(baseURL+"/tasks", wrapper.GetTasks)
	router.PATCH(baseURL+"/tasks/patch/:id", wrapper.PatchTasksPatchId)
	router.POST(baseURL+"/tasks/post", wrapper.PostTasksPost)
	router.DELETE(baseURL+"/tasks/:id", wrapper.DeleteTasksId)

}

type GetTasksRequestObject struct {
	Params GetTasksParams
}

type GetTasksResponseObject interface {
	VisitGetTasksResponse(w http.ResponseWriter) error
}

type GetTasks200JSONResponse []Task

func (response GetTasks200JSONResponse) VisitGetTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchTasksPatchIdRequestObject struct {
	Id   int `json:"id"`
	Body *PatchTasksPatchIdJSONRequestBody
}

type PatchTasksPatchIdResponseObject interface {
	VisitPatchTasksPatchIdResponse(w http.ResponseWriter) error
}

type PatchTasksPatchId200JSONResponse Task

func (response PatchTasksPatchId200JSONResponse) VisitPatchTasksPatchIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTasksPostRequestObject struct {
	Body *PostTasksPostJSONRequestBody
}

type PostTasksPostResponseObject interface {
	VisitPostTasksPostResponse(w http.ResponseWriter) error
}

type PostTasksPost201JSONResponse Task

func (response PostTasksPost201JSONResponse) VisitPostTasksPostResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTasksIdRequestObject struct {
	Id int `json:"id"`
}

type DeleteTasksIdResponseObject interface {
	VisitDeleteTasksIdResponse(w http.ResponseWriter) error
}

type DeleteTasksId204Response struct {
}

func (response DeleteTasksId204Response) VisitDeleteTasksIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all tasks
	// (GET /tasks)
	GetTasks(ctx context.Context, request GetTasksRequestObject) (GetTasksResponseObject, error)
	// Update a task
	// (PATCH /tasks/patch/{id})
	PatchTasksPatchId(ctx context.Context, request PatchTasksPatchIdRequestObject) (PatchTasksPatchIdResponseObject, error)
	// Create a new task
	// (POST /tasks/post)
	PostTasksPost(ctx context.Context, request PostTasksPostRequestObject) (PostTasksPostResponseObject, error)
	// Delete a task
	// (DELETE /tasks/{id})
	DeleteTasksId(ctx context.Context, request DeleteTasksIdRequestObject) (DeleteTasksIdResponseObject, error)
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

// GetTasks operation middleware
func (sh *strictHandler) GetTasks(ctx echo.Context, params GetTasksParams) error {
	var request GetTasksRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasks(ctx.Request().Context(), request.(GetTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksResponseObject); ok {
		return validResponse.VisitGetTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchTasksPatchId operation middleware
func (sh *strictHandler) PatchTasksPatchId(ctx echo.Context, id int) error {
	var request PatchTasksPatchIdRequestObject

	request.Id = id

	var body PatchTasksPatchIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchTasksPatchId(ctx.Request().Context(), request.(PatchTasksPatchIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchTasksPatchId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchTasksPatchIdResponseObject); ok {
		return validResponse.VisitPatchTasksPatchIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostTasksPost operation middleware
func (sh *strictHandler) PostTasksPost(ctx echo.Context) error {
	var request PostTasksPostRequestObject

	var body PostTasksPostJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostTasksPost(ctx.Request().Context(), request.(PostTasksPostRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTasksPost")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostTasksPostResponseObject); ok {
		return validResponse.VisitPostTasksPostResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteTasksId operation middleware
func (sh *strictHandler) DeleteTasksId(ctx echo.Context, id int) error {
	var request DeleteTasksIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTasksId(ctx.Request().Context(), request.(DeleteTasksIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTasksId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteTasksIdResponseObject); ok {
		return validResponse.VisitDeleteTasksIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
