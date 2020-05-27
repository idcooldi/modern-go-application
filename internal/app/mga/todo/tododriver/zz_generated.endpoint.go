// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package tododriver

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo"
)

// endpointError identifies an error that should be returned as an endpoint error.
type endpointError interface {
	EndpointError() bool
}

// serviceError identifies an error that should be returned as a service error.
type serviceError interface {
	ServiceError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	CreateTodo     endpoint.Endpoint
	ListTodos      endpoint.Endpoint
	MarkAsComplete endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service todo.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{
		CreateTodo:     kitxendpoint.OperationNameMiddleware("todo.CreateTodo")(mw(MakeCreateTodoEndpoint(service))),
		ListTodos:      kitxendpoint.OperationNameMiddleware("todo.ListTodos")(mw(MakeListTodosEndpoint(service))),
		MarkAsComplete: kitxendpoint.OperationNameMiddleware("todo.MarkAsComplete")(mw(MakeMarkAsCompleteEndpoint(service))),
	}
}

// CreateTodoRequest is a request struct for CreateTodo endpoint.
type CreateTodoRequest struct {
	Title string
}

// CreateTodoResponse is a response struct for CreateTodo endpoint.
type CreateTodoResponse struct {
	Todo todo.Todo
	Err  error
}

func (r CreateTodoResponse) Failed() error {
	return r.Err
}

// MakeCreateTodoEndpoint returns an endpoint for the matching method of the underlying service.
func MakeCreateTodoEndpoint(service todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTodoRequest)

		todo, err := service.CreateTodo(ctx, req.Title)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return CreateTodoResponse{
					Err:  err,
					Todo: todo,
				}, nil
			}

			return CreateTodoResponse{
				Err:  err,
				Todo: todo,
			}, err
		}

		return CreateTodoResponse{Todo: todo}, nil
	}
}

// ListTodosRequest is a request struct for ListTodos endpoint.
type ListTodosRequest struct{}

// ListTodosResponse is a response struct for ListTodos endpoint.
type ListTodosResponse struct {
	Todos []todo.Todo
	Err   error
}

func (r ListTodosResponse) Failed() error {
	return r.Err
}

// MakeListTodosEndpoint returns an endpoint for the matching method of the underlying service.
func MakeListTodosEndpoint(service todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		todos, err := service.ListTodos(ctx)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return ListTodosResponse{
					Err:   err,
					Todos: todos,
				}, nil
			}

			return ListTodosResponse{
				Err:   err,
				Todos: todos,
			}, err
		}

		return ListTodosResponse{Todos: todos}, nil
	}
}

// MarkAsCompleteRequest is a request struct for MarkAsComplete endpoint.
type MarkAsCompleteRequest struct {
	Id string
}

// MarkAsCompleteResponse is a response struct for MarkAsComplete endpoint.
type MarkAsCompleteResponse struct {
	Err error
}

func (r MarkAsCompleteResponse) Failed() error {
	return r.Err
}

// MakeMarkAsCompleteEndpoint returns an endpoint for the matching method of the underlying service.
func MakeMarkAsCompleteEndpoint(service todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MarkAsCompleteRequest)

		err := service.MarkAsComplete(ctx, req.Id)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return MarkAsCompleteResponse{Err: err}, nil
			}

			return MarkAsCompleteResponse{Err: err}, err
		}

		return MarkAsCompleteResponse{}, nil
	}
}
