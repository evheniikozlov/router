package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	testDefaultHttpMethod(t, http.MethodGet)
}

func TestPost(t *testing.T) {
	testDefaultHttpMethod(t, http.MethodPost)
}

func TestPut(t *testing.T) {
	testDefaultHttpMethod(t, http.MethodPut)
}

func TestDelete(t *testing.T) {
	testDefaultHttpMethod(t, http.MethodDelete)
}

func TestPatch(t *testing.T) {
	testDefaultHttpMethod(t, http.MethodPatch)
}

func TestConnect(t *testing.T) {
	testDefaultHttpMethod(t, http.MethodConnect)
}

func TestHead(t *testing.T) {
	testDefaultHttpMethod(t, http.MethodHead)
}

func TestOptions(t *testing.T) {
	testDefaultHttpMethod(t, http.MethodOptions)
}

func TestTrace(t *testing.T) {
	testDefaultHttpMethod(t, http.MethodTrace)
}

func TestNotFound(t *testing.T) {
	testHttpMethod(t, NewDefaultRouter(), http.MethodGet, "/not_found", "404 Not Found", http.StatusNotFound)
}

func TestCustomNotFound(t *testing.T) {
	const body = "404 Not Found. Go back to home"
	router := NewDefaultRouter()
	router.NotFound(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, body)
	})
	testHttpMethod(t, router, http.MethodGet, "/not_found", body, http.StatusNotFound)
}

func TestNamedParams(t *testing.T) {
	const name = "cars"
	router := NewDefaultRouter()
	router.Get("/posts/:name", func(responseWriter http.ResponseWriter, request *http.Request, params Params) {
		fmt.Fprint(responseWriter, params.GetString("name"))
	})
	testHttpMethod(t, router, http.MethodGet, fmt.Sprintf("/posts/%s", name), name, http.StatusOK)
}

func TestNamedParamsWithCustomRegexp(t *testing.T) {
	const taskId = 123
	router := NewDefaultRouter()
	router.Get("/tasks/:id^\\d*", func(responseWriter http.ResponseWriter, request *http.Request, params Params) {
		id, _ := params.GetInt("id")
		fmt.Fprint(responseWriter, id)
	})
	testHttpMethod(t, router, http.MethodGet, fmt.Sprintf("/tasks/%d", taskId), strconv.Itoa(taskId), http.StatusOK)
	testHttpMethod(t, router, http.MethodGet, "/tasks/error", "404 Not Found", http.StatusNotFound)
}

func testDefaultHttpMethod(t *testing.T, method string) {
	const body = "body"
	url := fmt.Sprintf("/%s", method)
	router := NewDefaultRouter()
	bindMethodToRouter(router, url, method, body)
	testHttpMethod(t, router, method, url, body, http.StatusOK)
}

func testHttpMethod(t *testing.T, router Router, method, url, body string, status int) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)
	if responseStatus := responseRecorder.Code; responseStatus != status {
		t.Errorf("response status is %d but expected %d", responseStatus, status)
	}
	if responseBody := responseRecorder.Body.String(); responseBody != body {
		t.Errorf("response body is %s but expected %s", responseBody, body)
	}
}

func bindMethodToRouter(router Router, url, method, body string) {
	handler := func(responseWriter http.ResponseWriter, request *http.Request, params Params) {
		fmt.Fprint(responseWriter, body)
	}
	switch method {
	case http.MethodGet:
		router.Get(url, handler)
		break
	case http.MethodPost:
		router.Post(url, handler)
		break
	case http.MethodPut:
		router.Put(url, handler)
		break
	case http.MethodDelete:
		router.Delete(url, handler)
		break
	case http.MethodPatch:
		router.Patch(url, handler)
		break
	case http.MethodConnect:
		router.Connect(url, handler)
		break
	case http.MethodHead:
		router.Head(url, handler)
		break
	case http.MethodOptions:
		router.Options(url, handler)
		break
	case http.MethodTrace:
		router.Trace(url, handler)
		break
	}
}
