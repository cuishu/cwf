package cwf

import (
	"net/http"
)

type Context struct {
	StatusCode int
	Response   http.ResponseWriter
	Request    *http.Request
}

func (c *Context) String(status int, s string) (int, error) {
	c.StatusCode = status
	c.Response.WriteHeader(status)
	return c.Response.Write([]byte(s))
}
