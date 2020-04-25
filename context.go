package cwf

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	StatusCode int
	Response   http.ResponseWriter
	Request    *http.Request
	Page404    func(*Context)
	Page500    func(*Context)
}

func (c *Context) Status(status int) {
	c.StatusCode = status
	c.Response.WriteHeader(status)
}

func (c *Context) String(status int, s string) (int, error) {
	c.Status(status)
	return c.Response.Write([]byte(s))
}

func (c *Context) JSON(status int, v interface{}) (int, error) {
	if js, err := json.Marshal(v); err != nil {
		c.Status(status)
		n, _ := c.Response.Write([]byte(err.Error()))
		return n, err
	} else {
		return c.Data(status, "application/json", js)
	}
}

func (c *Context) Data(status int, contentType string, data []byte) (int, error) {
	c.Status(status)
	c.Response.Header().Set("ContentType", contentType)
	return c.Response.Write(data)
}
