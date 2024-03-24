package gee

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

const abortIndex int = math.MaxInt >> 1

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
	handlers   []HandlerFunc
	index      int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Path:    req.URL.Path,
		Method:  req.Method,
		Request: req,
		Writer:  w,
		index:   -1,
	}
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) AbortWithStatus(code int) {
	c.Status(code)
	c.Writer.WriteHeader(code)
	c.Abort()
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) requestHeader(key string) string {
	return c.Request.Header.Get(key)
}

func (c *Context) GetHeader(key string) string {
	return c.requestHeader(key)
}

func (c *Context) DeleteHeader(key string) {
	c.Writer.Header().Del(key)
}
