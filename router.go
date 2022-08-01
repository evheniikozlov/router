package router

import (
	"net/http"
)

type Router interface {
	http.Handler
	Get(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	Post(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	Put(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	Delete(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	Patch(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	Connect(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	Head(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	Options(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	Trace(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	NotFound(handler func(responseWriter http.ResponseWriter, request *http.Request))
}
