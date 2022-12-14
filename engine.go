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

	// Register keys map to avoid duplicate registration
	registerKeys map[string]bool
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
		gin:          gin.Default(),
		initFuncs:    make([]func(), 0),
		middleware:   make([]gin.HandlerFunc, 0),
		routes:       make([]route, 0),
		preRunFuncs:  make([]func(), 0),
		registerKeys: make(map[string]bool),
	}
}

// Get gin engine
func (e *Engine) Gin() *gin.Engine {
	return e.gin
}

// PrependInitFunc prepend the init functions
func (e *Engine) PrependInitFunc(f ...func()) {
	e.initFuncs = append(f, e.initFuncs...)
}

// AppendInitFunc append the init functions
func (e *Engine) AppendInitFunc(f ...func()) {
	e.initFuncs = append(e.initFuncs, f...)
}

// PrependMiddleware prepend the middleware functions
// The generic type of the general middleware should be struct{}
func (e *Engine) PrependMiddleware(middleware ...gin.HandlerFunc) {
	e.middleware = append(middleware, e.middleware...)
}

// AppendMiddleware append the middleware functions
// The generic type of the general middleware should be struct{}
func (e *Engine) AppendMiddleware(middleware ...gin.HandlerFunc) {
	e.middleware = append(e.middleware, middleware...)
}

// PrependPreRunFunc prepend the pre-run functions
func (e *Engine) PrependPreRunFunc(f func()) {
	e.preRunFuncs = append([]func(){f}, e.preRunFuncs...)
}

// AppendPreRunFunc append the pre-run functions
func (e *Engine) AppendPreRunFunc(f func()) {
	e.preRunFuncs = append(e.preRunFuncs, f)
}

// Run start the engine
func (e *Engine) Run(addr string) {
	// init func
	for _, f := range e.initFuncs {
		f()
	}

	// middleware
	for _, m := range e.middleware {
		e.Gin().Use(m)
	}

	// route
	for _, r := range e.routes {
		e.Gin().Handle(r.method, r.path, r.handlers...)
	}

	// pre func
	for _, f := range e.preRunFuncs {
		f()
	}

	// run
	e.Gin().Run(addr)
}

// CheckDependencies check the dependencies
// If the dependencies are not registered, it will panic
func (e *Engine) CheckDependencies(pluginIDs ...string) {
	for _, pluginID := range pluginIDs {
		if _, ok := e.registerKeys[pluginID]; !ok {
			panic("The plugin " + pluginID + " is not registered")
		}
	}
}
