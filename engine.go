package ginseng

import "github.com/gin-gonic/gin"

// Ginseng core engine
// This engine is based on the Gin framework
type Engine struct {

	// Gin engine
	gin *gin.Engine

	// Initialize functions will be called before setting up the middleware and routes
	initFuncs []func()

	// Middleware functions will be called after the initialization functions and before setting up routes
	middleware []gin.HandlerFunc

	// Routes will be set up after the middleware functions
	routes []route

	// Pre-run functions will be called after the routes are set up and before the engine is started
	preRunFuncs []func()
}

// route store the route information
type route struct {
	method   string
	path     string
	handlers []gin.HandlerFunc
}

// NewEngine create a new engine
func NewEngine() *Engine {
	return &Engine{
		gin: gin.Default(),
	}
}

// PrependInitFunc prepend the init functions
func PrependInitFunc(e *Engine, f ...func()) {
	e.initFuncs = append(f, e.initFuncs...)
}

// AppendInitFunc append the init functions
func AppendInitFunc(e *Engine, f ...func()) {
	e.initFuncs = append(e.initFuncs, f...)
}

// PrependMiddleware prepend the middleware functions
// The generic type of the general middleware should be struct{}
func PrependMiddleware(e *Engine, middleware ...HandlerFunc[struct{}]) {
	ginHandlers := make([]gin.HandlerFunc, len(middleware))
	for i, m := range middleware {
		ginHandlers[i] = _toGinHandler(e, m)
	}
	e.middleware = append(ginHandlers, e.middleware...)
}

// AppendMiddleware append the middleware functions
// The generic type of the general middleware should be struct{}
func AppendMiddleware(e *Engine, middleware ...HandlerFunc[struct{}]) {
	ginHandlers := make([]gin.HandlerFunc, len(middleware))
	for i, m := range middleware {
		ginHandlers[i] = _toGinHandler(e, m)
	}
	e.middleware = append(e.middleware, ginHandlers...)
}

// PrependPreRunFunc prepend the pre-run functions
func PrependPreRunFunc(e *Engine, f func()) {
	e.preRunFuncs = append([]func(){f}, e.preRunFuncs...)
}

// AppendPreRunFunc append the pre-run functions
func AppendPreRunFunc(e *Engine, f func()) {
	e.preRunFuncs = append(e.preRunFuncs, f)
}
