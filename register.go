package ginseng

// RegisterHandlerFunc is used to register/install a plugins
// Each plugin should expose this function and its option type
// The exposed function should be named as "RegisterHandler"
// Each plugin should expose a const string to identify itself, which should be named as "PluginID"
// RegisterHandlerFunc of packages should not have dependencies on each other
type RegisterHandlerFunc[T any] func(e *Engine, option *T)

// Register is used to register/install a plugins
// Each plugin should expose this function and its option type
// The exposed function should be named as "RegisterHandler"
// Each plugin should expose a const string to identify itself, which should be named as "PluginID"
func Register[T any](e *Engine, pluginID string, f RegisterHandlerFunc[T], option *T) {
	if _, ok := e.registerKeys[pluginID]; ok {
		return // already registered, skip
	}
	e.registerKeys[pluginID] = true
	f(e, option)
}
