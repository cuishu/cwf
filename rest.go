package cwf

type ControllerInterface interface {
	GET(*Context)
	POST(*Context)
	PUT(*Context)
	DELETE(*Context)
	TRACE(*Context)
	OPTIONS(*Context)
	HEAD(*Context)
}

type Controller struct{}

func (c *Controller) GET(ctx *Context)     { ctx.Page404(ctx) }
func (c *Controller) POST(ctx *Context)    { ctx.Page404(ctx) }
func (c *Controller) PUT(ctx *Context)     { ctx.Page404(ctx) }
func (c *Controller) DELETE(ctx *Context)  { ctx.Page404(ctx) }
func (c *Controller) TRACE(ctx *Context)   { ctx.Page404(ctx) }
func (c *Controller) OPTIONS(ctx *Context) { ctx.Page404(ctx) }
func (c *Controller) HEAD(ctx *Context)    { ctx.Page404(ctx) }

func (c *CWF) REST(url string, rest ControllerInterface) {
	c.GET(url, rest.GET)
	c.POST(url, rest.POST)
	c.PUT(url, rest.PUT)
	c.DELETE(url, rest.DELETE)
	c.TRACE(url, rest.TRACE)
	c.OPTIONS(url, rest.OPTIONS)
	c.HEAD(url, rest.HEAD)
}
