package router

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type defaultRouter struct {
	handlers        map[handlerIdentifier]func(responseWriter http.ResponseWriter, request *http.Request, params Params)
	regexps         map[string]map[string]*regexp.Regexp
	notFoundHandler func(writer http.ResponseWriter, request *http.Request)
}

type handlerIdentifier struct {
	method string
	path   string
}

const regexpStart, regexpEnd, parameterDelimiter, defaultParameterRegexp = "^", "$", "^", ".*"

func NewDefaultRouter() *defaultRouter {
	router := defaultRouter{}
	router.handlers = make(map[handlerIdentifier]func(responseWriter http.ResponseWriter, request *http.Request, params Params))
	router.regexps = make(map[string]map[string]*regexp.Regexp)
	router.regexps[http.MethodGet] = make(map[string]*regexp.Regexp)
	router.regexps[http.MethodPost] = make(map[string]*regexp.Regexp)
	router.regexps[http.MethodPut] = make(map[string]*regexp.Regexp)
	router.regexps[http.MethodDelete] = make(map[string]*regexp.Regexp)
	router.regexps[http.MethodPatch] = make(map[string]*regexp.Regexp)
	router.regexps[http.MethodConnect] = make(map[string]*regexp.Regexp)
	router.regexps[http.MethodHead] = make(map[string]*regexp.Regexp)
	router.regexps[http.MethodOptions] = make(map[string]*regexp.Regexp)
	router.regexps[http.MethodTrace] = make(map[string]*regexp.Regexp)
	router.notFoundHandler = func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, "404 Not Found")
	}
	return &router
}

func (router *defaultRouter) Get(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handle(http.MethodGet, path, handler)
}

func (router *defaultRouter) Post(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handle(http.MethodPost, path, handler)
}

func (router *defaultRouter) Put(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handle(http.MethodPut, path, handler)
}

func (router *defaultRouter) Delete(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handle(http.MethodDelete, path, handler)
}

func (router *defaultRouter) Patch(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handle(http.MethodPatch, path, handler)
}

func (router *defaultRouter) Connect(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handle(http.MethodConnect, path, handler)
}

func (router *defaultRouter) Head(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handle(http.MethodHead, path, handler)
}

func (router *defaultRouter) Options(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handle(http.MethodOptions, path, handler)
}

func (router *defaultRouter) Trace(path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handle(http.MethodTrace, path, handler)
}

func (router *defaultRouter) NotFound(handler func(writer http.ResponseWriter, request *http.Request)) {
	router.notFoundHandler = handler
}

func (router defaultRouter) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	regexps := router.regexps[request.Method]
	matchedHandlerIdentifier := handlerIdentifier{method: request.Method}
	for path, pathRegexp := range regexps {
		if pathRegexp.MatchString(request.URL.Path) && path > matchedHandlerIdentifier.path {
			matchedHandlerIdentifier.path = path
		}
	}
	if matchedHandlerIdentifier.path != "" {
		router.handlers[matchedHandlerIdentifier](responseWriter, request, NewParamsByRegexp(request.URL.Path, router.regexps[matchedHandlerIdentifier.method][matchedHandlerIdentifier.path]))
	} else {
		router.notFoundHandler(responseWriter, request)
	}
}

func (router *defaultRouter) handle(method string, path string, handler func(responseWriter http.ResponseWriter, request *http.Request, params Params)) {
	router.handlers[handlerIdentifier{method: method, path: path}] = handler
	router.regexps[method][path] = router.parsePathToRegexp(path)
}

func (router defaultRouter) parsePathToRegexp(path string) *regexp.Regexp {
	pathRegexp := regexpStart
	var characterIndex int
	for characterIndex < len(path) {
		character := path[characterIndex]
		pathRegexpPart := string(character)
		if character == ':' {
			parameter := strings.Split(path[characterIndex+1:], "/")[0]
			characterIndex += len(parameter)
			parameterRegexp := defaultParameterRegexp
			if strings.Contains(parameter, parameterDelimiter) {
				parameterRegexp = strings.Split(parameter, parameterDelimiter)[1]
				parameter = strings.Split(parameter, parameterDelimiter)[0]
			}
			pathRegexpPart = fmt.Sprintf("(?P<%s>%s)", parameter, parameterRegexp)
		}
		pathRegexp += pathRegexpPart
		characterIndex++
	}
	pathRegexp += regexpEnd
	return regexp.MustCompile(pathRegexp)
}
