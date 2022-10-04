package ginseng

import (
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Ginseng handler function
type HandlerFunc[T any] func(*Context[T])

// Get setup the GET route
func Get[T any](e *Engine, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	_setupRoute(e, "GET", path, handler, middleware...)
}

// Post setup the POST route
func Post[T any](e *Engine, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	_setupRoute(e, "POST", path, handler, middleware...)
}

// Put setup the PUT route
func Put[T any](e *Engine, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	_setupRoute(e, "PUT", path, handler, middleware...)
}

// Delete setup the DELETE route
func Delete[T any](e *Engine, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	_setupRoute(e, "DELETE", path, handler, middleware...)
}

func _setupRoute[T any](e *Engine, method string, path string, handler HandlerFunc[T], middleware ...gin.HandlerFunc) {
	ginHandlers := append(middleware, _toGinHandler(e, handler))
	e.routes = append(e.routes, route{
		method:   method,
		path:     path,
		handlers: ginHandlers,
	})
}

func _toGinHandler[T any](e *Engine, handler HandlerFunc[T]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := new(T)
		tagKeys := _parseRequestTagKeys(request)

		for _, key := range tagKeys {
			switch key {
			case "uri":
				if err := ctx.ShouldBindUri(request); err != nil {
					log.Println(err)
				}
			case "json":
				if err := ctx.ShouldBindJSON(request); err != nil {
					log.Println(err)
				}
			case "form":
				if err := ctx.ShouldBindQuery(request); err != nil {
					log.Println(err)
				}
			}
		}

		context := &Context[T]{
			ginCtx:   ctx,
			Request:  request,
			Response: nil,
			param:    make(map[string]interface{}),
		}

		handler(context)
	}
}

// get all tag keys from the struct
func _parseRequestTagKeys[T any](request *T) []string {
	var keyMap = make(map[string]bool)

	numOfField := reflect.TypeOf(request).Elem().NumField()
	for i := 0; i < numOfField; i++ {
		tag := reflect.TypeOf(request).Elem().Field(i).Tag

		check := tag.Get("uri")
		if check != "" {
			keyMap["uri"] = true
		}

		check = tag.Get("json")
		if check != "" {
			keyMap["json"] = true
		}

		check = tag.Get("form")
		if check != "" {
			keyMap["form"] = true
		}
	}

	var keys []string
	for key := range keyMap {
		keys = append(keys, key)
	}

	return keys
}
