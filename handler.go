package cwf

import (
	"net/http"
	"time"
)

type Handler struct {
	getMap     map[string]func(ctx *Context)
	postMap    map[string]func(ctx *Context)
	putMap     map[string]func(ctx *Context)
	deleteMap  map[string]func(ctx *Context)
	traceMap   map[string]func(ctx *Context)
	optionsMap map[string]func(ctx *Context)
	headMap    map[string]func(ctx *Context)
}

func NewHandler() *Handler {
	handler := &Handler{}
	handler.getMap = make(map[string]func(ctx *Context))
	handler.postMap = make(map[string]func(ctx *Context))
	handler.putMap = make(map[string]func(ctx *Context))
	handler.deleteMap = make(map[string]func(ctx *Context))
	handler.traceMap = make(map[string]func(ctx *Context))
	handler.optionsMap = make(map[string]func(ctx *Context))
	handler.headMap = make(map[string]func(ctx *Context))
	return handler
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	startTime := time.Now()
	ctx := &Context{Response: res, Request: req}
	var f func(ctx *Context)
	switch req.Method {
	case GET:
		f = h.getMap[url]
	case POST:
		f = h.postMap[url]
	case PUT:
		f = h.putMap[url]
	case DELETE:
		f = h.deleteMap[url]
	case OPTIONS:
		f = h.optionsMap[url]
	case TRACE:
		f = h.traceMap[url]
	case HEAD:
		f = h.headMap[url]
	}
	if f != nil {
		f(ctx)
	} else {
		ctx.String(404, "not found.")
	}
	endTime := time.Now()
	logger.Info(cyan, req.Method, ctx.StatusCode, reset, url, endTime.Sub(startTime))
}

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
	HEAD    = "HEAD"
)

func (h *Handler) HandleFunc(url string, method string, f func(ctx *Context)) {
	switch method {
	case GET:
		h.getMap[url] = f
	case POST:
		h.postMap[url] = f
	case PUT:
		h.putMap[url] = f
	case DELETE:
		h.deleteMap[url] = f
	case OPTIONS:
		h.optionsMap[url] = f
	case TRACE:
		h.traceMap[url] = f
	case HEAD:
		h.headMap[url] = f
	default:
		logger.Error("Bad method: ", method)
	}
}
