package cwf

import (
	"net/http"
	"time"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
	HEAD    = "HEAD"
)

type route struct {
	GET     func(ctx *Context)
	POST    func(ctx *Context)
	PUT     func(ctx *Context)
	DELETE  func(ctx *Context)
	TRACE   func(ctx *Context)
	OPTIONS func(ctx *Context)
	HEAD    func(ctx *Context)
}

type Handler struct {
	routerMap map[string]route
	Page404   func(*Context)
	Page500   func(*Context)
}

func page404(ctx *Context) {
	ctx.String(404, "404 Not Found.")
}

func page500(ctx *Context) {
	ctx.String(500, "500 Internal Server Error.")
}

func NewHandler() *Handler {
	handler := &Handler{}
	handler.Page404 = page404
	handler.Page500 = page500
	handler.routerMap = make(map[string]route)
	return handler
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	startTime := time.Now()
	ctx := &Context{Response: res, Request: req, Page404: h.Page404, Page500: h.Page500}
	r := h.routerMap[url]
	var f func(ctx *Context)
	switch req.Method {
	case GET:
		f = r.GET
	case POST:
		f = r.POST
	case PUT:
		f = r.PUT
	case DELETE:
		f = r.DELETE
	case OPTIONS:
		f = r.OPTIONS
	case TRACE:
		f = r.TRACE
	case HEAD:
		f = r.HEAD
	}
	if f != nil {
		f(ctx)
	} else {
		h.Page404(ctx)
	}
	var statusColor string
	if ctx.StatusCode < 400 {
		statusColor = green
	} else if ctx.StatusCode < 500 {
		statusColor = yellow
	} else {
		statusColor = red
	}
	endTime := time.Now()
	logger.Info(cyan+req.Method+reset,
		statusColor, ctx.StatusCode, reset,
		url, endTime.Sub(startTime))
}

func (h *Handler) HandleFunc(url string, method string, f func(ctx *Context)) {
	var r route
	var ok bool
	if r, ok = h.routerMap[url]; !ok {
		r = route{}
	}
	switch method {
	case GET:
		r.GET = f
	case POST:
		r.POST = f
	case PUT:
		r.PUT = f
	case DELETE:
		r.DELETE = f
	case OPTIONS:
		r.OPTIONS = f
	case TRACE:
		r.TRACE = f
	case HEAD:
		r.HEAD = f
	default:
		logger.Error("Bad method: ", method)
	}
	h.routerMap[url] = r
}
