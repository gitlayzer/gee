package gee

func Server() HandlerFunc {
	return func(c *Context) {
		c.SetHeader("Server", "GeeServer")
	}
}
