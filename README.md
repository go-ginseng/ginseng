# Ginseng

Ginseng is a core go backend engine based on [gin](https://github.com/gin-gonic/gin).

## Installation

```bash
go get -u github.com/go-ginseng/ginseng
```

## Define your package

```go
package plugin

import "github.com/go-ginseng/ginseng"

const PluginID = "2141b9a1-f505-44c2-85a7-9a8c972e2d7d"

type HelloWorldRequest struct {
    Name string `json:"name"`
}

type HelloWorldOption struct {
    Language string `json:"language"`
}

var language = "en"

func HelloWorldHandler(ctx *ginseng.Context[HelloWorldRequest]) {
    switch language {
    case "en":
        ctx.GinCtx().JSON(200, gin.H{"message": "Hello " + ctx.Request.Name})
    case "zh":
        ctx.GinCtx().JSON(200, gin.H{"message": "你好 " + ctx.Request.Name})
    }
}

func RegisterHandler(e *ginseng.Engine, option *HelloWorldOption) {
    language = option.Language
    ginseng.Get(e, "/hello", HelloWorldHandler)
}
```

## Setup your engine

```go
package main

import (
    "github.com/go-ginseng/ginseng"
    "plugin"
)

func main() {
    e := ginseng.NewEngine()
    e.Register(e, plugin.PluginID, plugin.RegisterHandler, &plugin.HelloWorldOption{
        Language: "en",
    })
    e.Run("127.0.0.1:5000")
}
```
