package ginseng

import "github.com/gin-gonic/gin"

// Ginseng context
// This context is based on the Gin context
type Context[T any] struct {
	ginCtx   *gin.Context
	Request  *T
	Response interface{}
	param    map[string]interface{} // additional parameters
}

// Get request method
func (ctx *Context[T]) Method() string {
	return ctx.ginCtx.Request.Method
}

// Get request path
func (ctx *Context[T]) Path() string {
	path := ctx.ginCtx.FullPath()
	if path == "" {
		path = ctx.ginCtx.Request.URL.Path
	}
	return path
}

// Get client IP
func (ctx *Context[T]) ClientIP() string {
	return ctx.ginCtx.ClientIP()
}

// Get request header
func (ctx *Context[T]) Header(key string) string {
	return ctx.ginCtx.GetHeader(key)
}

// Get Param
// Param is the additional parameter set by the plugin
func (ctx *Context[T]) Param(key string) interface{} {
	return ctx.param[key]
}

// Set Param
// Param is the additional parameter set by the plugin
func (ctx *Context[T]) SetParam(key string, value interface{}) {
	ctx.param[key] = value
}
