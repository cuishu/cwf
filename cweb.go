// cwf project cwf.go
package cwf

import (
	"net/http"
	"time"
)

type CWF struct {
	handler *Handler
}

func Default() *CWF {
	cwf := &CWF{}
	cwf.handler = NewHandler()
	return cwf
}

func (c *CWF) GET(url string, f func(c *Context)) {
	c.handler.HandleFunc(url, GET, f)
}

func (c *CWF) POST(url string, f func(c *Context)) {
	c.handler.HandleFunc(url, POST, f)
}

func (c *CWF) PUT(url string, f func(c *Context)) {
	c.handler.HandleFunc(url, PUT, f)
}

func (c *CWF) DELETE(url string, f func(c *Context)) {
	c.handler.HandleFunc(url, DELETE, f)
}

func (c *CWF) OPTIONS(url string, f func(c *Context)) {
	c.handler.HandleFunc(url, OPTIONS, f)
}

func (c *CWF) HEAD(url string, f func(c *Context)) {
	c.handler.HandleFunc(url, HEAD, f)
}

func (c *CWF) TRACE(url string, f func(c *Context)) {
	c.handler.HandleFunc(url, TRACE, f)
}

func (c *CWF) Listen(addr string) {
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 60, //设置3秒的写超时
		Handler:      c.handler,
	}
	logger.Fatal(server.ListenAndServe())
}
